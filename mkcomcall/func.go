package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"io"
	"strings"
)

// writeFunc writes a syscall wrapper for f to w.
// It modifies f's data, so be warned!
func (m *module) writeFunc(w io.Writer, f *ast.FuncDecl) error {
	f.Recv = nil // Remove the bogus receiver (really the DLL name).
	err := m.printConfig.Fprint(w, m.fileSet, f)
	if err != nil {
		return err
	}
	fmt.Fprintln(w, " {")

	params, setupCode, resultCode, err := m.analyzeParameterList(f.Type)
	if err != nil {
		return err
	}

	prefix := ""
	if m.packageName != "com" {
		prefix = "com."
	}

	if setupCode != "" {
		fmt.Fprintln(w, setupCode)
	}
	fmt.Fprintf(w, "\t_res, _, _ := %sSyscall(proc%s.Addr(),\n\t\t%s)\n", prefix, f.Name.Name, strings.Join(params, ",\n\t\t"))

	fmt.Fprint(w, resultCode)
	fmt.Fprintln(w, "\treturn")

	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)
	return nil
}

// analyzeParameterList analyzes a function or method's parameter list and
// creates the actual parameter list that will be used to call it.
// setupCode is the code that needs to be called before the actuall call to the
// DLL function. resultCode is the code, to be run after the function call,
// that sets up the return value. It assumes that the first return value from
// the DLL function is assigned to a variable named _res.
func (m *module) analyzeParameterList(ft *ast.FuncType) (params []string, setupCode, resultCode string, err error) {
	prefix := ""
	if m.packageName != "com" {
		prefix = "com."
	}

	// Find errReturn or intReturn.
	if results := ft.Results; results != nil && len(results.List) > 0 {
		lastResult := results.List[len(results.List)-1]
		name := lastResult.Names[0].Name
		switch lrt := lastResult.Type.(type) {
		case *ast.Ident:
			switch lrt.Name {
			case "error":
				resultCode = "\tif _res != 0 {\n\t\t" + name + " = " + prefix + "HResult(_res)\n\t}\n"
				results.List = results.List[:len(results.List)-1]
			case "int":
				resultCode = "\t" + name + " = int(_res)\n"
				results.List = results.List[:len(results.List)-1]
			}

		case *ast.StarExpr:
			b := new(bytes.Buffer)
			err = m.printConfig.Fprint(b, m.fileSet, lrt)
			if err != nil {
				return
			}
			resultCode = "\t" + name + " = (" + b.String() + ")(unsafe.Pointer(_res))\n"
			results.List = results.List[:len(results.List)-1]
		}
	}
	if resultCode == "" {
		resultCode = "\t_ = _res\n"
	}

	for _, p := range ft.Params.List {
		if p.Names == nil {
			err = errors.New("anonymous parameters are not supported")
			return
		}
		for _, ident := range p.Names {
			switch t := p.Type.(type) {
			case *ast.Ident:
				switch t.Name {
				case "string":
					params = append(params, prefix+"BStrFromString("+ident.Name+").P")

				case "bool":
					params = append(params, prefix+"VariantBool("+ident.Name+")")

				default:
					params = append(params, ident.Name)
				}

			case *ast.StarExpr:
				params = append(params, ident.Name)

			case *ast.ArrayType:
				if t.Len == nil {
					// It's a slice.
					params = append(params, "&"+ident.Name+"[0]",
						"len("+ident.Name+")")
				} else {
					// It's an array.
					params = append(params, "&"+ident.Name)
				}

			case *ast.InterfaceType:
				params = append(params, prefix+"ToVariant("+ident.Name+")")

			default:
				params = append(params, ident.Name)
			}
		}
	}

	if ft.Results != nil {
		for _, p := range ft.Results.List {
			if p.Names == nil {
				err = errors.New("anonymous parameters are not supported")
				return
			}

			switch t := p.Type.(type) {
			case *ast.Ident:
				if t.Name == "string" {
					for _, ident := range p.Names {
						tmpName := "_tmp_" + ident.Name
						setupCode += fmt.Sprintf("\tvar %s %sBStr\n", tmpName, prefix)
						resultCode += fmt.Sprintf("\t%s = %s.String()\n\t%sSysFreeString(%s)\n",
							ident.Name, tmpName, prefix, tmpName)

						params = append(params, "&"+tmpName)
					}
					continue
				}

			case *ast.InterfaceType:
				for _, ident := range p.Names {
					tmpName := "_tmp_" + ident.Name
					setupCode += fmt.Sprintf("\tvar %s %sVariant\n", tmpName, prefix)
					resultCode += fmt.Sprintf("\t%s = %s.ToInterface()\n", ident.Name, tmpName)
					params = append(params, "&"+tmpName)
				}
				continue
			}

			for _, ident := range p.Names {
				params = append(params, "&"+ident.Name)
			}
		}
	}

	return
}
