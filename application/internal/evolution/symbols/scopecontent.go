package symbols

import (
	"bytes"
	"fmt"
	"go/types"
)

var basicKinds = map[types.BasicKind]string{
	types.Invalid: "Invalid",

	// predeclared types
	types.Bool:          "Bool",
	types.Int:           "Int",
	types.Int8:          "Int8",
	types.Int16:         "Int16",
	types.Int32:         "Int32",
	types.Int64:         "Int64",
	types.Uint:          "Uint",
	types.Uint8:         "Uint8",
	types.Uint16:        "Uint16",
	types.Uint32:        "Uint32",
	types.Uint64:        "Uint64",
	types.Uintptr:       "Uintptr",
	types.Float32:       "Float32",
	types.Float64:       "Float64",
	types.Complex64:     "Complex64",
	types.Complex128:    "Complex128",
	types.String:        "String",
	types.UnsafePointer: "UnsafePointer",

	// types for untyped values
	types.UntypedBool:    "UntypedBool",
	types.UntypedInt:     "UntypedInt",
	types.UntypedRune:    "UntypedRune",
	types.UntypedFloat:   "UntypedFloat",
	types.UntypedComplex: "UntypedComplex",
	types.UntypedString:  "UntypedString",
	types.UntypedNil:     "UntypedNil",

	// aliases
	// types.Byte: "Byte",
	// types.Rune: "Rune",
}

var basicInfos = map[types.BasicInfo]string{
	types.IsBoolean:  "IsBoolean",
	types.IsInteger:  "IsInteger",
	types.IsUnsigned: "IsUnsigned",
	types.IsFloat:    "IsFloat",
	types.IsComplex:  "IsComplex",
	types.IsString:   "IsString",
	types.IsUntyped:  "IsUntyped",

	types.IsOrdered:   "IsOrdered",
	types.IsNumeric:   "IsNumeric",
	types.IsConstType: "IsConstType",
}

type ScopeContent struct {
	TypeNamesBasic     []*types.TypeName
	TypeNamesNamed     []*types.TypeName
	TypeNamesInterface []*types.TypeName
	Consts             []*types.Const
	Funcs              []*types.Func
	PkgNames           []*types.PkgName
	Vars               []*types.Var
	Labels             []*types.Label
	Builtins           []*types.Builtin
	Nils               []*types.Nil
}

func NewScopeContent(scope *types.Scope) *ScopeContent {
	sc := &ScopeContent{}

	for _, name := range scope.Names() {
		obj := scope.Lookup(name)

		switch obj := obj.(type) {

		case *types.TypeName: // named types

			switch obj.Type().(type) {
			case *types.Basic:
				sc.TypeNamesBasic = append(sc.TypeNamesBasic, obj)

			case *types.Named:
				sc.TypeNamesNamed = append(sc.TypeNamesNamed, obj)

			case *types.Interface:
				sc.TypeNamesInterface = append(sc.TypeNamesInterface, obj)

			default:
				panic("Unhandled case for Object(TypeName) type")
			}

		case *types.Const:
			sc.Consts = append(sc.Consts, obj)

		case *types.Func:
			sc.Funcs = append(sc.Funcs, obj)

		case *types.PkgName:
			sc.PkgNames = append(sc.PkgNames, obj)

		case *types.Var:
			sc.Vars = append(sc.Vars, obj)

		case *types.Label:
			sc.Labels = append(sc.Labels, obj)

		case *types.Builtin:
			sc.Builtins = append(sc.Builtins, obj)

		case *types.Nil:
			sc.Nils = append(sc.Nils, obj)

		default:
			panic("Unhandled case for Object type")
		}
	}

	return sc
}

func (cs ScopeContent) Brief() string {
	ret := ""
	if len(cs.TypeNamesBasic) > 0 {
		ret = fmt.Sprintf("%s TypeNamesBasic: [", ret)
		for i, obj := range cs.TypeNamesBasic {
			if i != 0 {
				ret = fmt.Sprintf("%s, ", ret)
			}
			ret = fmt.Sprintf("%s%s", ret, obj.Name())
		}
		ret = fmt.Sprintf("%s] ", ret)
	}
	if len(cs.TypeNamesNamed) > 0 {
		ret = fmt.Sprintf("%s TypeNamesNamed: [", ret)
		for i, obj := range cs.TypeNamesNamed {
			if i != 0 {
				ret = fmt.Sprintf("%s, ", ret)
			}
			ret = fmt.Sprintf("%s%s", ret, obj.Name())
		}
		ret = fmt.Sprintf("%s] ", ret)
	}
	if len(cs.TypeNamesInterface) > 0 {
		ret = fmt.Sprintf("%s TypeNamesInterface: [", ret)
		for i, obj := range cs.TypeNamesInterface {
			if i != 0 {
				ret = fmt.Sprintf("%s, ", ret)
			}
			ret = fmt.Sprintf("%s%s", ret, obj.Name())
		}
		ret = fmt.Sprintf("%s] ", ret)
	}
	if len(cs.Consts) > 0 {
		ret = fmt.Sprintf("%s Consts: [", ret)
		for i, obj := range cs.Consts {
			if i != 0 {
				ret = fmt.Sprintf("%s, ", ret)
			}
			ret = fmt.Sprintf("%s%s", ret, obj.Name())
		}
		ret = fmt.Sprintf("%s] ", ret)
	}
	if len(cs.Funcs) > 0 {
		ret = fmt.Sprintf("%s Funcs: [", ret)
		for i, obj := range cs.Funcs {
			if i != 0 {
				ret = fmt.Sprintf("%s, ", ret)
			}
			ret = fmt.Sprintf("%s%s", ret, obj.Name())
		}
		ret = fmt.Sprintf("%s] ", ret)
	}
	if len(cs.PkgNames) > 0 {
		ret = fmt.Sprintf("%s PkgNames: [", ret)
		for i, obj := range cs.PkgNames {
			if i != 0 {
				ret = fmt.Sprintf("%s, ", ret)
			}
			ret = fmt.Sprintf("%s%s", ret, obj.Name())
		}
		ret = fmt.Sprintf("%s] ", ret)
	}
	if len(cs.Vars) > 0 {
		ret = fmt.Sprintf("%s Vars: [", ret)
		for i, obj := range cs.Vars {
			if i != 0 {
				ret = fmt.Sprintf("%s, ", ret)
			}
			ret = fmt.Sprintf("%s%s", ret, obj.Name())
		}
		ret = fmt.Sprintf("%s] ", ret)
	}
	if len(cs.Labels) > 0 {
		ret = fmt.Sprintf("%s Labels: [", ret)
		for i, obj := range cs.Labels {
			if i != 0 {
				ret = fmt.Sprintf("%s, ", ret)
			}
			ret = fmt.Sprintf("%s%s", ret, obj.Name())
		}
		ret = fmt.Sprintf("%s] ", ret)
	}
	if len(cs.Builtins) > 0 {
		ret = fmt.Sprintf("%s Builtins: [", ret)
		for i, obj := range cs.Builtins {
			if i != 0 {
				ret = fmt.Sprintf("%s, ", ret)
			}
			ret = fmt.Sprintf("%s%s", ret, obj.Name())
		}
		ret = fmt.Sprintf("%s] ", ret)
	}
	if len(cs.Nils) > 0 {
		ret = fmt.Sprintf("%s Nils: [", ret)
		for i, obj := range cs.Nils {
			if i != 0 {
				ret = fmt.Sprintf("%s, ", ret)
			}
			ret = fmt.Sprintf("%s%s", ret, obj.Name())
		}
		ret = fmt.Sprintf("%s] ", ret)
	}
	return ret
}

func (cs ScopeContent) Markdown() string {
	buf := bytes.NewBuffer([]byte{})

	if len(cs.TypeNamesBasic) > 0 {
		fmt.Fprintln(buf, "**TypeName / Basic**")
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "| String | Name | Type | Kind | Info | Name |")
		fmt.Fprintln(buf, "|---|")
		for _, obj := range cs.TypeNamesBasic {
			fmt.Fprintf(buf, "| `%s` | %s | %s | %s | %s | %s |\n",
				obj.String(), obj.Name(), obj.Type(),
				basicKinds[obj.Type().(*types.Basic).Kind()],
				basicInfos[obj.Type().(*types.Basic).Info()],
				obj.Type().(*types.Basic).Name(),
			)
		}
		fmt.Fprintln(buf)
	}

	if len(cs.TypeNamesNamed) > 0 {
		fmt.Fprintln(buf, "**TypeName / Named**")
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "| String | Name | Type |")
		fmt.Fprintln(buf, "|---|")
		for _, obj := range cs.TypeNamesNamed {
			fmt.Fprintf(buf, "| `%s` | %s | %s |\n", obj.String(), obj.Name(), obj.Type())
		}
		fmt.Fprintln(buf)
	}

	if len(cs.TypeNamesInterface) > 0 {
		fmt.Fprintln(buf, "**TypeName / Interface**")
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "| String | Name | Type |")
		fmt.Fprintln(buf, "|---|")
		for _, obj := range cs.TypeNamesInterface {
			fmt.Fprintf(buf, "| `%s` | %s | %s |\n", obj.String(), obj.Name(), obj.Type())
		}
		fmt.Fprintln(buf)
	}

	if len(cs.Consts) > 0 {
		fmt.Fprintln(buf, "**Const**")
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "| String | Name | Type | Val |")
		fmt.Fprintln(buf, "|---|")
		for _, obj := range cs.Consts {
			fmt.Fprintf(buf, "| `%s` | %s | %s | %s |\n", obj.String(), obj.Name(), obj.Type(), obj.Val())
		}
		fmt.Fprintln(buf)
	}

	if len(cs.Funcs) > 0 {
		fmt.Fprintln(buf, "**Func**")
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "| String | Name | FullName |")
		fmt.Fprintln(buf, "|---|")
		for _, obj := range cs.Funcs {
			fmt.Fprintf(buf, "| `%s` | %s | %s |\n", obj.String(), obj.Name(), obj.FullName())
		}
		fmt.Fprintln(buf)
	}

	if len(cs.PkgNames) > 0 {
		fmt.Fprintln(buf, "**PkgName**")
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "| String | Name | Type |")
		fmt.Fprintln(buf, "|---|")
		for _, obj := range cs.PkgNames {
			fmt.Fprintf(buf, "| `%s` | %s | %s |\n", obj.String(), obj.Name(), obj.Type())
		}
		fmt.Fprintln(buf)
	}

	if len(cs.Vars) > 0 {
		fmt.Fprintln(buf, "**Var**")
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "| String | Name | Type | Anonymous | Embedded | IsField |")
		fmt.Fprintln(buf, "|---|")
		for _, obj := range cs.Vars {
			fmt.Fprintf(buf, "| `%s` | %s | %s | %t | %t | %t |\n", obj.String(), obj.Name(), obj.Type(), obj.Anonymous(), obj.Embedded(), obj.IsField())
		}
		fmt.Fprintln(buf)
	}

	if len(cs.Labels) > 0 {
		fmt.Fprintln(buf, "**Label**")
		fmt.Fprintln(buf, "| String | Name | Type |")
		fmt.Fprintln(buf, "|---|")
		for _, obj := range cs.Labels {
			fmt.Fprintf(buf, "| `%s` | %s | %s |\n", obj.String(), obj.Name(), obj.Type())
		}
		fmt.Fprintln(buf)
	}

	if len(cs.Builtins) > 0 {
		fmt.Fprintln(buf, "**Builtin**")
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "| String | Name | Type |")
		fmt.Fprintln(buf, "|---|")
		for _, obj := range cs.Builtins {
			fmt.Fprintf(buf, "| `%s` | %s | %s |\n", obj.String(), obj.Name(), obj.Type())
		}
		fmt.Fprintln(buf)
	}

	if len(cs.Nils) > 0 {
		fmt.Fprintln(buf, "**Nil**")
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "| String | Name | Type |")
		fmt.Fprintln(buf, "|---|")
		for _, obj := range cs.Nils {
			fmt.Fprintf(buf, "| `%s` | %s | %s |\n", obj.String(), obj.Name(), obj.Type())
		}
	}

	return buf.String()
}

func (cs ScopeContent) String() string {
	buf := bytes.NewBuffer([]byte{})

	fmt.Fprintln(buf, "Found TypeName:")
	{
		fmt.Fprintln(buf, "    Found Basic:")
		for _, obj := range cs.TypeNamesBasic {
			fmt.Fprintf(buf, "        Name=%-15s | String=%-15s | Type=%-15s | Kind=%-15s | Info=%-15s | Name=%-15s\n",
				obj.Name(), obj.String(), obj.Type(),
				basicKinds[obj.Type().(*types.Basic).Kind()],
				basicInfos[obj.Type().(*types.Basic).Info()],
				obj.Type().(*types.Basic).Name(),
			)
		}

		fmt.Fprintln(buf, "    Found Named:")
		for _, obj := range cs.TypeNamesNamed {
			fmt.Fprintf(buf, "        Name=%-15s | String=%-50s | Type=%-15s\n", obj.Name(), obj.String(), obj.Type())
		}

		fmt.Fprintln(buf, "    Found Interface:")
		for _, obj := range cs.TypeNamesInterface {
			fmt.Fprintf(buf, "        Name=%-15s | String=%-50s | Type=%-15s\n", obj.Name(), obj.String(), obj.Type())
		}
	}

	fmt.Fprintln(buf, "Found Const:")
	for _, obj := range cs.Consts {
		fmt.Fprintf(buf, "    Name=%-15s | Val=%-15s    | String=%-15s\n", obj.Name(), obj.Val(), obj.String())
	}

	fmt.Fprintln(buf, "Found Func:")
	for _, obj := range cs.Funcs {
		fmt.Fprintf(buf, "    Name=%-15s | FullName=%-15s\n", obj.Name(), obj.FullName())
	}

	fmt.Fprintln(buf, "Found PkgName:")
	for _, obj := range cs.PkgNames {
		fmt.Fprintf(buf, "    Name=%-15s | String=%-15s | Type=%-15s\n", obj.Name(), obj.String(), obj.Type())
	}

	fmt.Fprintln(buf, "Found Var:")
	for _, obj := range cs.Vars {
		fmt.Fprintf(buf, "    Anonymous=%t | Embedded=%t | IsField=%t\n", obj.Anonymous(), obj.Embedded(), obj.IsField())
	}

	fmt.Fprintln(buf, "Found Label:")
	for _, obj := range cs.Labels {
		fmt.Fprintf(buf, "    Name=%-15s\n", obj.Name())
	}

	fmt.Fprintln(buf, "Found Builtin:")
	for _, obj := range cs.Builtins {
		fmt.Fprintf(buf, "    Name=%-15s | String=%-15s | Type=%-15s\n", obj.Name(), obj.String(), obj.Type().String())
	}

	fmt.Fprintln(buf, "Found Nil:")
	for _, obj := range cs.Nils {
		fmt.Fprintf(buf, "    Name=%-15s | String=%-15s\n", obj.Name(), obj.String())
	}

	return buf.String()
}

func markdownPackageRecHelper(scope *types.Scope, name string, d int, limit int) *bytes.Buffer {
	h := strings.Repeat("#", d+2)
	f := bytes.NewBuffer([]byte{})
	fmt.Fprintf(f, "%s %s\n\n", h, name)
	fmt.Fprintln(f, NewScopeContent(scope).Markdown())
	if limit != d {
		for i := 0; i < scope.NumChildren(); i++ {
			firstline := strings.Split(scope.String(), "\n")[0]
			title := strings.Join(strings.Split(firstline, " ")[0:2], " ")
			io.Copy(f, markdownPackageRecHelper(scope.Child(i), title, d+1, limit))
		}
	}
	return f
}

// use limit to restrain what depth/level of nested-scopes will be visited
func PrintPackageAsMarkdown(pkg *types.Package, limit int) *bytes.Buffer {
	f := bytes.NewBuffer([]byte{})
	fmt.Fprintf(f, "\n# Universe\n\n")
	fmt.Fprintln(f, NewScopeContent(types.Universe).Markdown())
	io.Copy(f, markdownPackageRecHelper(pkg.Scope(), "The Package", 0, limit))
	return f
}
