package main

import (
	"errors"
	"fmt"
	"go/ast"
	"io"
	"strings"
)

// An iface represents a COM interface that we will be generating code for.
type iface struct {
	declaration  *ast.InterfaceType
	name         string
	iID, classID string

	// extends is the name of an interface that this one extends.
	extends string

	// vtStart is the vTable index of its first method.
	vtStart int

	// methods is the list of methods this interface has.
	methods []*ast.Field
}

// newIface returns an iface for the interface declared in ts.
func newIface(ts *ast.TypeSpec) (*iface, error) {
	i := new(iface)
	i.name = ts.Name.Name
	if ts.Doc != nil {
		for _, comment := range ts.Doc.List {
			text := comment.Text
			text = strings.Trim(text, "/*")
			text = strings.TrimSpace(text)
			switch {
			case strings.HasPrefix(text, "IID"):
				i.iID = getGUID(text)
			case strings.HasPrefix(text, "CLSID"):
				i.classID = getGUID(text)
			}
		}
	}

	ifType, ok := ts.Type.(*ast.InterfaceType)
	if !ok {
		return nil, fmt.Errorf("%s is not defined as an interface type.", i.name)
	}

	for _, meth := range ifType.Methods.List {
		if meth.Names == nil {
			// It is an anonymous member, so it is the type that this interface extends.
			if i.extends != "" {
				return nil, fmt.Errorf("%s seems to be trying to have multiple inheritance.", i.name)
			}
			switch mt := (meth.Type).(type) {
			case *ast.Ident:
				i.extends = mt.Name
			case *ast.SelectorExpr:
				pkg, ok := mt.X.(*ast.Ident)
				if !ok {
					return nil, fmt.Errorf("%s extends a type with too complicated a name.")
				}
				if pkg.Name != "com" {
					return nil, fmt.Errorf("%s extends a type that is not in its own package or in the com package", i.name)
				}
				i.extends = "com." + mt.Sel.Name
			default:
				return nil, fmt.Errorf("%s extends a type that is not supported (%T)", i.name, mt)
			}
			continue
		}

		i.methods = append(i.methods, meth)
	}

	return i, nil
}

// getGUID extracts a GUID (enclosed in braces) from s. If there is no GUID
// present, it returns the empty string.
func getGUID(s string) string {
	opening := strings.Index(s, "{")
	closing := strings.Index(s, "}")
	if opening == -1 || closing < opening {
		return ""
	}
	return s[opening : closing+1]
}

// calcVTStart calculates the VTable offset for the interface named ifName.
func (m *module) calcVTStart(ifName string, depth int) error {
	if depth >= 100 {
		return errors.New("interface inheritance hierarchy is too deep")
	}

	i := m.interfaces[ifName]
	if i.vtStart > 0 {
		// It's already done.
		return nil
	}

	switch i.extends {
	case "":
		i.vtStart = 0
	case "com.IUnknown":
		i.vtStart = 3
	case "com.IDispatch":
		i.vtStart = 7

	default:
		base, ok := m.interfaces[i.extends]
		if !ok {
			return fmt.Errorf("interface %s not found", i.extends)
		}
		if base.vtStart == 0 {
			err := m.calcVTStart(i.extends, depth+1)
			if err != nil {
				return err
			}
		}
		i.vtStart = base.vtStart + len(base.methods)
	}

	return nil
}

// writeInterface writes a wrapper for i to w.
// It modifies i's data, so be warned!
func (m *module) writeInterface(w io.Writer, i *iface) error {
	prefix := "com."
	if m.packageName == "com" {
		prefix = ""
	}

	if i.iID != "" {
		fmt.Fprintf(w, "var IID_%s = %sNewGUID(%q)\n", i.name, prefix, i.iID)
	}
	if i.classID != "" {
		fmt.Fprintf(w, "var CLSID_%s = %sNewGUID(%q)\n", i.name, prefix, i.classID)
	}
	fmt.Fprintf(w, "type %s struct {\n", i.name)
	if i.extends == "" {
		fmt.Fprintf(w, "\t*%sVTable\n", prefix)
	} else {
		fmt.Fprintf(w, "\t%s\n", i.extends)
	}
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	for n, meth := range i.methods {
		fd := &ast.FuncDecl{
			Recv: &ast.FieldList{
				List: []*ast.Field{
					&ast.Field{
						Names: []*ast.Ident{
							&ast.Ident{
								Name: "this",
							},
						},
						Type: &ast.StarExpr{
							X: &ast.Ident{
								Name: i.name,
							},
						},
					},
				},
			},
			Name: &ast.Ident{
				Name: meth.Names[0].Name,
			},
			Type: meth.Type.(*ast.FuncType),
		}
		err := m.printConfig.Fprint(w, m.fileSet, fd)
		if err != nil {
			return err
		}

		fmt.Fprintln(w, " {")

		params, setupCode, errReturn, intReturn, err := m.analyzeParameterList(meth.Type.(*ast.FuncType))
		if err != nil {
			return err
		}
		params = append([]string{"uintptr(unsafe.Pointer(this))"}, params...)
		nParams := len(params)

		if setupCode != "" {
			fmt.Fprintln(w, setupCode)
		}

		funcName := "Syscall"
		paramsNeeded := 3
		switch {
		case len(params) > 18:
			return fmt.Errorf("method %s.%s has too many parameters.", i.name, meth.Names[0].Name)
		case len(params) > 15:
			funcName = "Syscall18"
			paramsNeeded = 18
		case len(params) > 12:
			funcName = "Syscall15"
			paramsNeeded = 15
		case len(params) > 9:
			funcName = "Syscall12"
			paramsNeeded = 12
		case len(params) > 6:
			funcName = "Syscall9"
			paramsNeeded = 9
		case len(params) > 3:
			funcName = "Syscall6"
			paramsNeeded = 6
		}
		for len(params) < paramsNeeded {
			params = append(params, "0")
		}

		fmt.Fprintf(w, "\t_res, _, _ := syscall.%s(this.VTable[%d], %d,\n\t\t%s)\n",
			funcName,
			i.vtStart+n,
			nParams,
			strings.Join(params, ",\n\t\t"))

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
	}

	return nil
}
