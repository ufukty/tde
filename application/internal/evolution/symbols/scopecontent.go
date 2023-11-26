package symbols

import (
	"fmt"
	"go/ast"
	"go/types"
	"reflect"
	"tde/internal/astw/traced"
)

type ScopeContent struct {
	TypeNames []*types.TypeName
	Consts    []*types.Const
	Funcs     []*types.Func
	PkgNames  []*types.PkgName
	Vars      []*types.Var
	Labels    []*types.Label
	Builtins  []*types.Builtin
	Nils      []*types.Nil
}

func NewScopeContent(scope *types.Scope) *ScopeContent {
	sc := &ScopeContent{}

	for _, name := range scope.Names() {
		obj := scope.Lookup(name)

		switch obj := obj.(type) {

		case *types.TypeName: // named types
			sc.TypeNames = append(sc.TypeNames, obj)

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

func LocalScope(fd *ast.FuncDecl, spot ast.Node) (*ScopeContent, error) {
	parents := traced.Parents(fd, spot)
	if len(parents) == 0 {
		return nil, fmt.Errorf("no trace found")
	}

	var sc *ScopeContent

	for _, n := range parents {
		// find the scope related with n and its custom children (eg. FuncDecl->FuncType)
		fmt.Println(reflect.TypeOf(n))
		switch n := n.(type) {
		case *ast.Package:
			fmt.Println(n)

		case *ast.File:
			fmt.Println(n)

		case *ast.FuncDecl:
			fmt.Println(n)
			// TODO: n.Type -> scope

		}
	}

	return sc, nil
}
