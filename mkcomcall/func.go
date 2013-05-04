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
	fmt.Fprintln(w, "{")

	params, setupCode, resultCode, err := m.analyzeParameterList(f.Type)
	if err != nil {
		return err
	}

	if setupCode != "" {
		fmt.Fprintln(w, setupCode)
	}
	fmt.Fprintf(w, "\t_res, _, _ := proc%s.Call(%s)\n", f.Name.Name, strings.Join(params, ",\n\t\t"))

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
	// Find errReturn or intReturn.
	if results := ft.Results; results != nil && len(results.List) > 0 {
		lastResult := results.List[len(results.List)-1]
		name := lastResult.Names[0].Name
		switch lrt := lastResult.Type.(type) {
		case *ast.Ident:
			switch lrt.Name {
			case "error":
				resultCode = "\tif _res != 0 {\n\t\t" + name + " = "
				if m.packageName != "com" {
					resultCode += "com."
				}
				resultCode += "HResult(_res)\n\t}\n"
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
				case "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "byte", "rune":
					params = append(params, "uintptr("+ident.Name+")")
					continue

				case "string":
					s := "uintptr(unsafe.Pointer("
					if m.packageName != "com" {
						s += "com."
					}
					s += "BStrFromString(" + ident.Name + ").P))"
					params = append(params, s)
					continue

				case "BStr":
					params = append(params, "uintptr(unsafe.Pointer("+ident.Name+".P))")
					continue
				}

			case *ast.StarExpr:
				params = append(params, "uintptr(unsafe.Pointer("+ident.Name+"))")
				continue

			case *ast.ArrayType:
				if t.Len == nil {
					// It's a slice.
					params = append(params, "uintptr(unsafe.Pointer(&"+ident.Name+"[0]))",
						"uintptr(len("+ident.Name+"))")
				} else {
					// It's an array.
					params = append(params, "uintptr(unsafe.Pointer(&"+ident.Name+"))")
				}
				continue

			case *ast.SelectorExpr:
				if ident, ok := t.X.(*ast.Ident); ok && ident.Name == "com" && t.Sel.Name == "BStr" {
					params = append(params, "uintptr(unsafe.Pointer("+ident.Name+".P))")
					continue
				}
			}

			buf := new(bytes.Buffer)
			m.printConfig.Fprint(buf, m.fileSet, p.Type)
			err = fmt.Errorf("unsupported parameter type: %s", buf)
			return
		}
	}

	for _, p := range ft.Results.List {
		if p.Names == nil {
			err = errors.New("anonymous parameters are not supported")
			return
		}
		for _, ident := range p.Names {
			params = append(params, "uintptr(unsafe.Pointer(&"+ident.Name+"))")
		}
	}

	return
}
