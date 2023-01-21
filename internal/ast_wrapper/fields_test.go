package ast_wrapper

import (
	"go/ast"
	"reflect"
	"testing"
)

func TestFields(t *testing.T) {

	var forTestingInterfaceFields = struct {
		Stmt ast.Stmt
		Decl ast.Decl
		Spec ast.Spec
		Expr ast.Expr

		ExprSlice  []ast.Expr
		StmtSlice  []ast.Stmt
		DeclSlice  []ast.Decl
		SpecSlice  []ast.Spec
		IdentSlice []*ast.Ident
	}{}

	var constructors = map[NodeType]func() ast.Node{
		ArrayType:      func() ast.Node { return &ast.ArrayType{} },
		AssignStmt:     func() ast.Node { return &ast.AssignStmt{} },
		BadDecl:        func() ast.Node { return &ast.BadDecl{} },
		BadExpr:        func() ast.Node { return &ast.BadExpr{} },
		BadStmt:        func() ast.Node { return &ast.BadStmt{} },
		BasicLit:       func() ast.Node { return &ast.BasicLit{} },
		BinaryExpr:     func() ast.Node { return &ast.BinaryExpr{} },
		BlockStmt:      func() ast.Node { return &ast.BlockStmt{} },
		BranchStmt:     func() ast.Node { return &ast.BranchStmt{} },
		CallExpr:       func() ast.Node { return &ast.CallExpr{} },
		CaseClause:     func() ast.Node { return &ast.CaseClause{} },
		ChanType:       func() ast.Node { return &ast.ChanType{} },
		CommClause:     func() ast.Node { return &ast.CommClause{} },
		Comment:        func() ast.Node { return &ast.Comment{} },
		CommentGroup:   func() ast.Node { return &ast.CommentGroup{} },
		CompositeLit:   func() ast.Node { return &ast.CompositeLit{} },
		DeclStmt:       func() ast.Node { return &ast.DeclStmt{} },
		DeferStmt:      func() ast.Node { return &ast.DeferStmt{} },
		Ellipsis:       func() ast.Node { return &ast.Ellipsis{} },
		EmptyStmt:      func() ast.Node { return &ast.EmptyStmt{} },
		ExprStmt:       func() ast.Node { return &ast.ExprStmt{} },
		Field:          func() ast.Node { return &ast.Field{} },
		FieldList:      func() ast.Node { return &ast.FieldList{} },
		File:           func() ast.Node { return &ast.File{} },
		ForStmt:        func() ast.Node { return &ast.ForStmt{} },
		FuncDecl:       func() ast.Node { return &ast.FuncDecl{} },
		FuncLit:        func() ast.Node { return &ast.FuncLit{} },
		FuncType:       func() ast.Node { return &ast.FuncType{} },
		GenDecl:        func() ast.Node { return &ast.GenDecl{} },
		GoStmt:         func() ast.Node { return &ast.GoStmt{} },
		Ident:          func() ast.Node { return &ast.Ident{} },
		IfStmt:         func() ast.Node { return &ast.IfStmt{} },
		ImportSpec:     func() ast.Node { return &ast.ImportSpec{} },
		IncDecStmt:     func() ast.Node { return &ast.IncDecStmt{} },
		IndexExpr:      func() ast.Node { return &ast.IndexExpr{} },
		IndexListExpr:  func() ast.Node { return &ast.IndexListExpr{} },
		InterfaceType:  func() ast.Node { return &ast.InterfaceType{} },
		KeyValueExpr:   func() ast.Node { return &ast.KeyValueExpr{} },
		LabeledStmt:    func() ast.Node { return &ast.LabeledStmt{} },
		MapType:        func() ast.Node { return &ast.MapType{} },
		Package:        func() ast.Node { return &ast.Package{} },
		ParenExpr:      func() ast.Node { return &ast.ParenExpr{} },
		RangeStmt:      func() ast.Node { return &ast.RangeStmt{} },
		ReturnStmt:     func() ast.Node { return &ast.ReturnStmt{} },
		SelectorExpr:   func() ast.Node { return &ast.SelectorExpr{} },
		SelectStmt:     func() ast.Node { return &ast.SelectStmt{} },
		SendStmt:       func() ast.Node { return &ast.SendStmt{} },
		SliceExpr:      func() ast.Node { return &ast.SliceExpr{} },
		StarExpr:       func() ast.Node { return &ast.StarExpr{} },
		StructType:     func() ast.Node { return &ast.StructType{} },
		SwitchStmt:     func() ast.Node { return &ast.SwitchStmt{} },
		TypeAssertExpr: func() ast.Node { return &ast.TypeAssertExpr{} },
		TypeSpec:       func() ast.Node { return &ast.TypeSpec{} },
		TypeSwitchStmt: func() ast.Node { return &ast.TypeSwitchStmt{} },
		UnaryExpr:      func() ast.Node { return &ast.UnaryExpr{} },
		ValueSpec:      func() ast.Node { return &ast.ValueSpec{} },

		Expr: func() ast.Node { return forTestingInterfaceFields.Expr },
		Stmt: func() ast.Node { return forTestingInterfaceFields.Stmt },
		Decl: func() ast.Node { return forTestingInterfaceFields.Decl },
		Spec: func() ast.Node { return forTestingInterfaceFields.Spec },
	}

	testCases := []ast.Node{
		&ast.Comment{},
		&ast.CommentGroup{},
		&ast.Field{},
		&ast.FieldList{},
		&ast.BadExpr{},
		&ast.Ident{},
		&ast.BasicLit{},
		&ast.Ellipsis{},
		&ast.FuncLit{},
		&ast.CompositeLit{},
		&ast.ParenExpr{},
		&ast.SelectorExpr{},
		&ast.IndexExpr{},
		&ast.IndexListExpr{},
		&ast.SliceExpr{},
		&ast.TypeAssertExpr{},
		&ast.CallExpr{},
		&ast.StarExpr{},
		&ast.UnaryExpr{},
		&ast.BinaryExpr{},
		&ast.KeyValueExpr{},
		&ast.ArrayType{},
		&ast.StructType{},
		&ast.FuncType{},
		&ast.InterfaceType{},
		&ast.MapType{},
		&ast.ChanType{},
		&ast.BadStmt{},
		&ast.DeclStmt{},
		&ast.EmptyStmt{},
		&ast.LabeledStmt{},
		&ast.ExprStmt{},
		&ast.SendStmt{},
		&ast.IncDecStmt{},
		&ast.AssignStmt{},
		&ast.GoStmt{},
		&ast.DeferStmt{},
		&ast.ReturnStmt{},
		&ast.BranchStmt{},
		&ast.BlockStmt{},
		&ast.IfStmt{},
		&ast.CaseClause{},
		&ast.SwitchStmt{},
		&ast.TypeSwitchStmt{},
		&ast.CommClause{},
		&ast.SelectStmt{},
		&ast.ForStmt{},
		&ast.RangeStmt{},
		&ast.ImportSpec{},
		&ast.ValueSpec{},
		&ast.TypeSpec{},
		&ast.BadDecl{},
		&ast.GenDecl{},
		&ast.FuncDecl{},
		&ast.File{},
		&ast.Package{},
	}

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		t.Error("Failed on checking if returned type is the right one for field: ", r)
	// 	}
	// }()

	for _, testCase := range testCases {
		fields, types := Fields(testCase)
		for i := 0; i < len(types); i++ {
			fieldType := types[i]

			if fieldType.IsSliceType() {
				switch fieldType {
				case ExprSlice:
					fields[i] = []ast.Expr{}
					if _, ok := fields[i].([]ast.Expr); !ok {
						t.Errorf("Failed on case ExprSlice, testCase type %s: ", reflect.TypeOf(testCase))
					}
				case StmtSlice:
					fields[i] = []ast.Stmt{}
					if _, ok := fields[i].([]ast.Stmt); !ok {
						t.Errorf("Failed on case StmtSlice, testCase type %s: ", reflect.TypeOf(testCase))
					}
				case DeclSlice:
					fields[i] = []ast.Decl{}
					if _, ok := fields[i].([]ast.Decl); !ok {
						t.Errorf("Failed on case DeclSlice, testCase type %s: ", reflect.TypeOf(testCase))
					}
				case SpecSlice:
					fields[i] = []ast.Spec{}
					if _, ok := fields[i].([]ast.Spec); !ok {
						t.Errorf("Failed on case SpecSlice, testCase type %s: ", reflect.TypeOf(testCase))
					}
				case IdentSlice:
					fields[i] = []*ast.Ident{}
					if _, ok := fields[i].([]*ast.Ident); !ok {
						t.Errorf("Failed on case IdentSlice, testCase type %s: ", reflect.TypeOf(testCase))
					}
				}
			} else if fieldType.IsInterfaceType() {
				switch fieldType {
				case Expr:
					fields[i] = forTestingInterfaceFields.Expr
					if _, ok := fields[i].(ast.Expr); !ok {
						t.Errorf("Failed on case Expr, testCase type %s: ", reflect.TypeOf(testCase))
					}
				case Stmt:
					fields[i] = forTestingInterfaceFields.Stmt
					if _, ok := fields[i].(ast.Stmt); !ok {
						t.Errorf("Failed on case Stmt, testCase type %s: ", reflect.TypeOf(testCase))
					}
				case Decl:
					fields[i] = forTestingInterfaceFields.Decl
					if _, ok := fields[i].(ast.Decl); !ok {
						t.Errorf("Failed on case Decl, testCase type %s: ", reflect.TypeOf(testCase))
					}
				case Spec:
					fields[i] = forTestingInterfaceFields.Spec
					if _, ok := fields[i].(ast.Spec); !ok {
						t.Errorf("Failed on case Spec, testCase type %s: ", reflect.TypeOf(testCase))
					}
				}

			} else {
				fields[i] = (constructors[types[i]])()
			}
		}
	}

}
