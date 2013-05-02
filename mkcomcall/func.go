package main

import (
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

	params, setupCode, errReturn, intReturn, err := m.analyzeParameterList(f.Type)
	if err != nil {
		return err
	}

	if setupCode != "" {
		fmt.Fprintln(w, setupCode)
	}
	fmt.Fprintf(w, "\t_res, _, _ := proc%s.Call(%s)\n", f.Name.Name, strings.Join(params, ",\n\t\t"))

	switch {
	case errReturn != "":
		fmt.Fprint(w, "\tif _res != 0 {\n\t\t", errReturn, " = ")
		if m.packageName != "com" {
			fmt.Fprint(w, "com.")
		}
		fmt.Fprint(w, "HResult(_res)\n\t}\n")

	case intReturn != "":
		fmt.Fprint(w, "\t", intReturn, " = int(_res)\n")

	default:
		fmt.Fprintln(w, "\t_ = _res")
	}
	fmt.Fprintln(w, "\treturn")

	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)
	return nil
}

// analyzeParameterList analyzes a function or method's parameter list and
// creates the actual parameter list that will be used to call it.
// setupCode is the code that needs to be called before the actuall call to the
// DLL function. If the last return value is an error or an int, its name will
// be put in errReturn or intReturn respectively.
func (m *module) analyzeParameterList(ft *ast.FuncType) (params []string, setupCode, errReturn, intReturn string, err error) {
	// Find errReturn or intReturn.
	if results := ft.Results; results != nil && len(results.List) > 0 {
		lastResult := results.List[len(results.List)-1]
		if typeIdent, ok := lastResult.Type.(*ast.Ident); ok && len(lastResult.Names) == 1 {
			name := lastResult.Names[0].Name
			switch typeIdent.Name {
			case "error":
				errReturn = name
				results.List = results.List[:len(results.List)-1]
			case "int":
				intReturn = name
				results.List = results.List[:len(results.List)-1]
			}
		}
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
				case "int":
					params = append(params, "uintptr("+ident.Name+")")

				case "string":
					tempName := fmt.Sprintf("_p%d", len(params))
					setupCode += "\tvar " + tempName + " *uint16\n\t" +
						tempName + ", err = syscall.UTF16PtrFromString(" + ident.Name + ")\n" +
						"\tif err != nil {\n" +
						"\t\treturn\n" +
						"\t}\n"
					params = append(params, "uintptr(unsafe.Pointer("+tempName+"))")

				default:
					err = fmt.Errorf("unsupported parameter type: %s", t.Name)
					return
				}

			default:
				err = fmt.Errorf("unsupported parameter type: %s", t)
			}
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
