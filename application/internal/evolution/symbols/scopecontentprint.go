package symbols

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/types"
	"io"
	"reflect"
	"strings"
)

type tablePrinter interface {
	Title(string)
	Head(...string)
	Row(...any)
}

// this implements TablePrinter and accepted by ScopeContent.writeTo()
type markdownTablePrinter struct {
	f io.Writer
}

func (p *markdownTablePrinter) Title(title string) {
	fmt.Fprintf(p.f, "**%s**\n\n", title)
}

func (p *markdownTablePrinter) Head(cols ...string) {
	fmt.Fprintf(p.f, "|")
	for _, col := range cols {
		fmt.Fprintf(p.f, " %s |", col)
	}
	fmt.Fprintf(p.f, "\n|%s\n", strings.Repeat("---|", len(cols)))
}

func (p *markdownTablePrinter) Row(cols ...any) {
	tpl := "| " + strings.Repeat("`%s` |", len(cols)) + "\n"
	fmt.Fprintf(p.f, tpl, cols...)
}

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

func (cs ScopeContent) Brief() string {
	ret := ""
	if len(cs.TypeNames) > 0 {
		ret = fmt.Sprintf("%s TypeNames: [", ret)
		for i, obj := range cs.TypeNames {
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

type Row struct {
	objrel  []any
	typerel []any
}

type table struct {
	title string
	head  []string
	rows  []Row
}

func (t table) WriteTo(tp tablePrinter) {
	tp.Title(t.title)
	tp.Head(t.head...)
	for _, row := range t.rows {
		tp.Row(append(row.objrel, row.typerel...)...)
	}
}

func typerelatedfields(obj types.Object) []string {
	switch t := obj.Type().(type) {
	case *types.Array:
		return []string{"Elem", "Len"}
	case *types.Basic:
		return []string{"Name", "Kind", "Info"}
	case *types.Chan:
		return []string{"Dir", "Elem"}
	case *types.Interface:
		return []string{"IsComparable", "NumEmbeddeds", "NumMethods"}
	case *types.Map:
		return []string{"Key", "Elem"}
	case *types.Named:
		return []string{"NumMethods", "TypeArgs", "TypeParams"}
	case *types.Pointer:
		return []string{"Elem"}
	case *types.Signature:
		return []string{"RecvTypeParams", "Recv", "Params", "Results"}
	case *types.Slice:
		return []string{"Elem"}
	case *types.Struct:
		return []string{"NumFields"}
	case *types.Tuple:
		return []string{"Len"}
	case *types.TypeParam:
		return []string{"Index"}
	case *types.Union:
		return []string{"Len"}
	default:
		panic(fmt.Sprintf("unexpected type: %q", reflect.TypeOf(t)))
	}
}

func typerelatedvalues(obj types.Object) []any {
	switch t := obj.Type().(type) {
	case *types.Array:
		return []any{t.Elem(), t.Len()}
	case *types.Basic:
		return []any{t.Name(), basicKinds[t.Kind()], basicInfos[t.Info()]}
	case *types.Chan:
		return []any{t.Dir(), t.Elem()}
	case *types.Interface:
		return []any{t.IsComparable(), t.NumEmbeddeds(), t.NumMethods()}
	case *types.Map:
		return []any{t.Key(), t.Elem()}
	case *types.Named:
		return []any{t.NumMethods(), t.TypeArgs(), t.TypeParams()}
	case *types.Pointer:
		return []any{t.Elem()}
	case *types.Signature:
		return []any{t.RecvTypeParams(), t.Recv(), t.Params(), t.Results()}
	case *types.Slice:
		return []any{t.Elem()}
	case *types.Struct:
		return []any{t.NumFields()}
	case *types.Tuple:
		return []any{t.Len()}
	case *types.TypeParam:
		return []any{t.Index()}
	case *types.Union:
		return []any{t.Len()}
	default:
		panic(fmt.Sprintf("unexpected type: %q", reflect.TypeOf(t)))
	}
}

func objectrelatedfields(obj types.Object) []string {
	switch o := obj.(type) {
	case *types.TypeName:
		return []string{"String", "Name", "Type", "IsAlias"}

	case *types.Const:
		return []string{"String", "Name", "Type", "Val"}

	case *types.Func:
		return []string{"String", "Name", "Type", "FullName", "Origin", "Pkg"}

	case *types.PkgName:
		return []string{"String", "Name", "Type"}

	case *types.Var:
		return []string{"String", "Name", "Type", "Anonymous", "Embedded", "IsField", "Origin"}

	case *types.Label:
		return []string{"String", "Name", "Type"}

	case *types.Builtin:
		return []string{"String", "Name", "Type"}

	case *types.Nil:
		return []string{"String", "Name", "Type"}

	default:
		panic(fmt.Sprintf("unexpected object type: %q", reflect.TypeOf(o)))
	}
}

func objectrelatedvalues(obj types.Object) []any {
	switch o := obj.(type) {
	case *types.TypeName:
		return []any{obj.String(), obj.Name(), obj.Type(), o.IsAlias()}

	case *types.Const:
		return []any{obj.String(), obj.Name(), obj.Type(), o.Val()}

	case *types.Func:
		return []any{obj.String(), obj.Name(), obj.Type(), o.FullName(), o.Origin(), o.Pkg()}

	case *types.PkgName:
		return []any{obj.String(), obj.Name(), obj.Type()}

	case *types.Var:
		return []any{obj.String(), obj.Name(), obj.Type(), o.Anonymous(), o.Embedded(), o.IsField(), o.Origin()}

	case *types.Label:
		return []any{obj.String(), obj.Name(), obj.Type()}

	case *types.Builtin:
		return []any{obj.String(), obj.Name(), obj.Type()}

	case *types.Nil:
		return []any{obj.String(), obj.Name(), obj.Type()}
	default:
		panic(fmt.Sprintf("unexpected object type: %q", reflect.TypeOf(o)))
	}
}

type tables struct {
	ObjectImpTyp string

	// field names are for the "types" of types
	Array     []types.Object
	Basic     []types.Object
	Chan      []types.Object
	Interface []types.Object
	Map       []types.Object
	Named     []types.Object
	Pointer   []types.Object
	Signature []types.Object
	Slice     []types.Object
	Struct    []types.Object
	Tuple     []types.Object
	TypeParam []types.Object
	Union     []types.Object
}

func (t tables) GetTables() (tbls []table) {
	ts := map[string][]types.Object{
		"Array":     t.Array,
		"Basic":     t.Basic,
		"Chan":      t.Chan,
		"Interface": t.Interface,
		"Map":       t.Map,
		"Named":     t.Named,
		"Pointer":   t.Pointer,
		"Signature": t.Signature,
		"Slice":     t.Slice,
		"Struct":    t.Struct,
		"Tuple":     t.Tuple,
		"TypeParam": t.TypeParam,
		"Union":     t.Union,
	}
	for k, v := range ts {
		if len(v) > 0 {
			table := table{
				title: t.ObjectImpTyp + "/" + k,
				head:  append(objectrelatedfields(v[0]), typerelatedfields(v[0])...),
				rows:  []Row{},
			}
			for _, obj := range v {
				table.rows = append(table.rows, Row{
					typerel: typerelatedvalues(obj),
					objrel:  objectrelatedvalues(obj),
				})
			}
			tbls = append(tbls, table)
		}
	}
	return tbls
}

func (t tables) WriteTo(tp tablePrinter) {
	for _, table := range t.GetTables() {
		table.WriteTo(tp)
	}
}

func newTables[C []O, O types.Object](c C, title string) tables {
	tables := tables{ObjectImpTyp: title}
	for _, obj := range c {
		switch obj.Type().(type) {
		case *types.Array:
			tables.Array = append(tables.Array, obj)
		case *types.Basic:
			tables.Basic = append(tables.Basic, obj)
		case *types.Chan:
			tables.Chan = append(tables.Chan, obj)
		case *types.Interface:
			tables.Interface = append(tables.Interface, obj)
		case *types.Map:
			tables.Map = append(tables.Map, obj)
		case *types.Named:
			tables.Named = append(tables.Named, obj)
		case *types.Pointer:
			tables.Pointer = append(tables.Pointer, obj)
		case *types.Signature:
			tables.Signature = append(tables.Signature, obj)
		case *types.Slice:
			tables.Slice = append(tables.Slice, obj)
		case *types.Struct:
			tables.Struct = append(tables.Struct, obj)
		case *types.Tuple:
			tables.Tuple = append(tables.Tuple, obj)
		case *types.TypeParam:
			tables.TypeParam = append(tables.TypeParam, obj)
		case *types.Union:
			tables.Union = append(tables.Union, obj)
		}
	}
	return tables
}

func (cs ScopeContent) writeTo(tp tablePrinter) {
	if len(cs.TypeNames) > 0 {
		newTables(cs.TypeNames, "TypeNames").WriteTo(tp)
	}
	if len(cs.Consts) > 0 {
		newTables(cs.Consts, "Consts").WriteTo(tp)
	}
	if len(cs.Funcs) > 0 {
		newTables(cs.Funcs, "Funcs").WriteTo(tp)
	}
	if len(cs.PkgNames) > 0 {
		newTables(cs.PkgNames, "PkgNames").WriteTo(tp)
	}
	if len(cs.Vars) > 0 {
		newTables(cs.Vars, "Vars").WriteTo(tp)
	}
	if len(cs.Labels) > 0 {
		newTables(cs.Labels, "Labels").WriteTo(tp)
	}
	if len(cs.Builtins) > 0 {
		newTables(cs.Builtins, "Builtins").WriteTo(tp)
	}
	if len(cs.Nils) > 0 {
		newTables(cs.Nils, "Nils").WriteTo(tp)
	}
}

// debug and practice purposes
func (cs ScopeContent) markdown() string {
	buf := bytes.NewBuffer([]byte{})
	tp := &markdownTablePrinter{f: buf}
	cs.writeTo(tp)
	return buf.String()
}

func markdownPackageRecHelper(f io.Writer, info *types.Info, r ast.Node, scope *types.Scope, d int, limit int) {
	fmt.Fprintf(f, "%s %s\n\n", strings.Repeat("#", d+1), findMeaningfulPathToScope(info, r, scope))
	fmt.Fprintln(f, NewScopeContent(scope).markdown())
	if limit != d {
		for i := 0; i < scope.NumChildren(); i++ {
			markdownPackageRecHelper(f, info, r, scope.Child(i), d+1, limit)
		}
	}
}

// use limit to restrain what depth/level of nested-scopes will be visited
func PrintPackageAsMarkdown(f io.Writer, info *types.Info, r ast.Node, scope *types.Scope, limit int) {
	markdownPackageRecHelper(f, info, r, scope, 0, limit)
}
