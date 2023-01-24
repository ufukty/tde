package field_inspector

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"tde/internal/cfg/astcfg/context"
	"tde/internal/cfg/astcfg/node_constructor"
)

func GetConstructor(typeOf reflect.Type) func(context.Context, int) ast.Node {

	if typeOf.Kind() == reflect.Slice {
		typeOf = typeOf.Elem()
	}

	switch typeOf.Kind() {

	case reflect.Pointer:
		if constructor, ok := node_constructor.ConstructorsByReflectType[typeOf]; ok {
			return constructor
		}

	case reflect.Interface:
		if constructor, ok := node_constructor.InterfaceConstructorsByReflectType[typeOf]; ok {
			return constructor
		}
	}

	return nil
}

// Returns the indices of available spots which can be a nil (unassigned) field
// or a slice type field that bot accepts ast.Node type values
func GetListOfAvailableSpots(node ast.Node) (indices []int) {
	typeOf := reflect.TypeOf(node).Elem()
	valueOf := reflect.ValueOf(node).Elem()
	numberOfFields := typeOf.NumField()

	var fieldType reflect.Type
	var fieldValue reflect.Value

	for i := 0; i < numberOfFields; i++ {
		fieldType = typeOf.FieldByIndex([]int{i}).Type
		fieldValue = valueOf.FieldByIndex([]int{i})

		fmt.Println(fieldType, fieldType.Name(), fieldType.Kind())

		switch fieldType.Kind() {

		case reflect.Pointer:
			fmt.Println(fieldType, fieldValue.IsNil(), fieldValue.CanSet())
			if _, ok := node_constructor.ConstructorsByReflectType[fieldType]; ok && fieldValue.IsNil() && fieldValue.CanSet() {
				indices = append(indices, i)
				fmt.Println("added")
			}

		case reflect.Slice:
			sliceElemType := fieldType.Elem()
			if _, ok := node_constructor.ConstructorsByReflectType[sliceElemType]; ok && fieldValue.IsNil() && fieldValue.CanSet() {
				indices = append(indices, i)
				fmt.Println("added")
			}
		}

		fmt.Println()
	}
	return
}

// func PickRandomAvailableField(typeOf reflect.Type, valueOf reflect.Value) (index int) {
// 	indicesOfAvailableFields :=
// 	return 0
// }

func inspect(n ast.Node) {
	// fmt.Printf("%#+v\n", n)

	typeOf := reflect.TypeOf(n).Elem()
	valueOf := reflect.ValueOf(n).Elem()
	numberOfFields := typeOf.NumField()

	// availableFields := GetListOfAvailableFields()

	for i := 0; i < numberOfFields; i++ {
		fieldType := typeOf.FieldByIndex([]int{i})
		fieldValue := valueOf.FieldByIndex([]int{i})

		fmt.Println(fieldType, fieldValue)
		switch kind := fieldType.Type.Kind(); kind {

		case reflect.Pointer:
			nodeType := fieldType.Type
			fmt.Println(kind, "detected, underlying type's kind is", nodeType)
			if constructor, ok := node_constructor.ConstructorsByReflectType[nodeType]; ok {
				node := constructor(context.NewContext(), 1)
				if fieldValue.CanSet() {
					fmt.Println("settable")
					fieldValue.Set(reflect.ValueOf(node))

					fmt.Println("set the node type:", reflect.TypeOf(node))
				} else {
					fmt.Println("not settable")
				}
				fmt.Println("constructor of", nodeType, "is found")

			}

		case reflect.Slice:
			fmt.Println(kind, "detected, underlying type's kind is", fieldType.Type.Elem().Name())

		case reflect.Interface:
			fmt.Println(kind, "detected, underlying type's kind is", fieldType.Type.Name())
			// fieldValue.set

		default:
			fmt.Println(kind, "passed")
		}

		// if fieldValue.IsNil() {
		// 	fmt.Println("says nil")
		// }

		fmt.Println("")
	}

	// switch n.(type) {
	// case *ast.BlockStmt:

	// case ast.Stmt:
	// }
}

func main() {

	f := &ast.File{
		Name: ast.NewIdent("foo"),
		Decls: []ast.Decl{
			&ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: ast.NewIdent("Bar"),
						Type: ast.NewIdent("uint32"),
					},
				},
			},
		},
	}

	ast.Print(token.NewFileSet(), f.Decls[0])
	inspect(f.Decls[0].(*ast.GenDecl).Specs[0])
	ast.Print(token.NewFileSet(), f.Decls[0])

	// fmt.Println("===")
	// inspect(f.Decls[0])

}
