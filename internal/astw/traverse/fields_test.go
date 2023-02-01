package traverse

import (
	ast_types "tde/internal/astw/types"
	
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

	var constructors = map[ast_types.NodeType]func() ast.Node{
		ast_types.ArrayType:      func() ast.Node { return &ast.ArrayType{} },
		ast_types.AssignStmt:     func() ast.Node { return &ast.AssignStmt{} },
		ast_types.BadDecl:        func() ast.Node { return &ast.BadDecl{} },
		ast_types.BadExpr:        func() ast.Node { return &ast.BadExpr{} },
		ast_types.BadStmt:        func() ast.Node { return &ast.BadStmt{} },
		ast_types.BasicLit:       func() ast.Node { return &ast.BasicLit{} },
		ast_types.BinaryExpr:     func() ast.Node { return &ast.BinaryExpr{} },
		ast_types.BlockStmt:      func() ast.Node { return &ast.BlockStmt{} },
		ast_types.BranchStmt:     func() ast.Node { return &ast.BranchStmt{} },
		ast_types.CallExpr:       func() ast.Node { return &ast.CallExpr{} },
		ast_types.CaseClause:     func() ast.Node { return &ast.CaseClause{} },
		ast_types.ChanType:       func() ast.Node { return &ast.ChanType{} },
		ast_types.CommClause:     func() ast.Node { return &ast.CommClause{} },
		ast_types.Comment:        func() ast.Node { return &ast.Comment{} },
		ast_types.CommentGroup:   func() ast.Node { return &ast.CommentGroup{} },
		ast_types.CompositeLit:   func() ast.Node { return &ast.CompositeLit{} },
		ast_types.DeclStmt:       func() ast.Node { return &ast.DeclStmt{} },
		ast_types.DeferStmt:      func() ast.Node { return &ast.DeferStmt{} },
		ast_types.Ellipsis:       func() ast.Node { return &ast.Ellipsis{} },
		ast_types.EmptyStmt:      func() ast.Node { return &ast.EmptyStmt{} },
		ast_types.ExprStmt:       func() ast.Node { return &ast.ExprStmt{} },
		ast_types.Field:          func() ast.Node { return &ast.Field{} },
		ast_types.FieldList:      func() ast.Node { return &ast.FieldList{} },
		ast_types.File:           func() ast.Node { return &ast.File{} },
		ast_types.ForStmt:        func() ast.Node { return &ast.ForStmt{} },
		ast_types.FuncDecl:       func() ast.Node { return &ast.FuncDecl{} },
		ast_types.FuncLit:        func() ast.Node { return &ast.FuncLit{} },
		ast_types.FuncType:       func() ast.Node { return &ast.FuncType{} },
		ast_types.GenDecl:        func() ast.Node { return &ast.GenDecl{} },
		ast_types.GoStmt:         func() ast.Node { return &ast.GoStmt{} },
		ast_types.Ident:          func() ast.Node { return &ast.Ident{} },
		ast_types.IfStmt:         func() ast.Node { return &ast.IfStmt{} },
		ast_types.ImportSpec:     func() ast.Node { return &ast.ImportSpec{} },
		ast_types.IncDecStmt:     func() ast.Node { return &ast.IncDecStmt{} },
		ast_types.IndexExpr:      func() ast.Node { return &ast.IndexExpr{} },
		ast_types.IndexListExpr:  func() ast.Node { return &ast.IndexListExpr{} },
		ast_types.InterfaceType:  func() ast.Node { return &ast.InterfaceType{} },
		ast_types.KeyValueExpr:   func() ast.Node { return &ast.KeyValueExpr{} },
		ast_types.LabeledStmt:    func() ast.Node { return &ast.LabeledStmt{} },
		ast_types.MapType:        func() ast.Node { return &ast.MapType{} },
		ast_types.Package:        func() ast.Node { return &ast.Package{} },
		ast_types.ParenExpr:      func() ast.Node { return &ast.ParenExpr{} },
		ast_types.RangeStmt:      func() ast.Node { return &ast.RangeStmt{} },
		ast_types.ReturnStmt:     func() ast.Node { return &ast.ReturnStmt{} },
		ast_types.SelectorExpr:   func() ast.Node { return &ast.SelectorExpr{} },
		ast_types.SelectStmt:     func() ast.Node { return &ast.SelectStmt{} },
		ast_types.SendStmt:       func() ast.Node { return &ast.SendStmt{} },
		ast_types.SliceExpr:      func() ast.Node { return &ast.SliceExpr{} },
		ast_types.StarExpr:       func() ast.Node { return &ast.StarExpr{} },
		ast_types.StructType:     func() ast.Node { return &ast.StructType{} },
		ast_types.SwitchStmt:     func() ast.Node { return &ast.SwitchStmt{} },
		ast_types.TypeAssertExpr: func() ast.Node { return &ast.TypeAssertExpr{} },
		ast_types.TypeSpec:       func() ast.Node { return &ast.TypeSpec{} },
		ast_types.TypeSwitchStmt: func() ast.Node { return &ast.TypeSwitchStmt{} },
		ast_types.UnaryExpr:      func() ast.Node { return &ast.UnaryExpr{} },
		ast_types.ValueSpec:      func() ast.Node { return &ast.ValueSpec{} },

		ast_types.Expr: func() ast.Node { return forTestingInterfaceFields.Expr },
		ast_types.Stmt: func() ast.Node { return forTestingInterfaceFields.Stmt },
		ast_types.Decl: func() ast.Node { return forTestingInterfaceFields.Decl },
		ast_types.Spec: func() ast.Node { return forTestingInterfaceFields.Spec },
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
		fields := NodeFields(testCase)
		for _, field := range fields {

			if field.Type.IsSliceType() {
				switch field.Type {
				case ast_types.ExprSlice:
					field.Value = []ast.Expr{}
					if _, ok := field.Value.([]ast.Expr); !ok {
						t.Errorf("Failed on case ExprSlice, testCase type %s: ", reflect.TypeOf(testCase))
					}
				case ast_types.StmtSlice:
					field.Value = []ast.Stmt{}
					if _, ok := field.Value.([]ast.Stmt); !ok {
						t.Errorf("Failed on case StmtSlice, testCase type %s: ", reflect.TypeOf(testCase))
					}
				case ast_types.DeclSlice:
					field.Value = []ast.Decl{}
					if _, ok := field.Value.([]ast.Decl); !ok {
						t.Errorf("Failed on case DeclSlice, testCase type %s: ", reflect.TypeOf(testCase))
					}
				case ast_types.SpecSlice:
					field.Value = []ast.Spec{}
					if _, ok := field.Value.([]ast.Spec); !ok {
						t.Errorf("Failed on case SpecSlice, testCase type %s: ", reflect.TypeOf(testCase))
					}
				case ast_types.IdentSlice:
					field.Value = []*ast.Ident{}
					if _, ok := field.Value.([]*ast.Ident); !ok {
						t.Errorf("Failed on case IdentSlice, testCase type %s: ", reflect.TypeOf(testCase))
					}
				}
			} else if field.Type.IsInterfaceType() {
				switch field.Type {
				case ast_types.Expr:
					field.Value = forTestingInterfaceFields.Expr
					if _, ok := field.Value.(ast.Expr); !ok {
						t.Errorf("Failed on case Expr, testCase type %s: ", reflect.TypeOf(testCase))
					}
				case ast_types.Stmt:
					field.Value = forTestingInterfaceFields.Stmt
					if _, ok := field.Value.(ast.Stmt); !ok {
						t.Errorf("Failed on case Stmt, testCase type %s: ", reflect.TypeOf(testCase))
					}
				case ast_types.Decl:
					field.Value = forTestingInterfaceFields.Decl
					if _, ok := field.Value.(ast.Decl); !ok {
						t.Errorf("Failed on case Decl, testCase type %s: ", reflect.TypeOf(testCase))
					}
				case ast_types.Spec:
					field.Value = forTestingInterfaceFields.Spec
					if _, ok := field.Value.(ast.Spec); !ok {
						t.Errorf("Failed on case Spec, testCase type %s: ", reflect.TypeOf(testCase))
					}
				}

			} else {
				field.Value = (constructors[field.Type])()
			}
		}
	}

}
