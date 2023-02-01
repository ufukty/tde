package node_constructor

import (
	"tde/internal/cfg/context"

	"go/ast"
	"reflect"
)

var (
	ConstructorsByReflectType          map[reflect.Type]func(ctx *context.Context, limit int) ast.Node
	InterfaceConstructorsByReflectType map[reflect.Type]func(ctx *context.Context, limit int) ast.Node
)

func init() {
	ConstructorsByReflectType = map[reflect.Type]func(ctx *context.Context, limit int) ast.Node{
		reflect.TypeOf(&ast.ArrayType{}):      func(ctx *context.Context, limit int) ast.Node { return ArrayType(ctx, limit) },
		reflect.TypeOf(&ast.AssignStmt{}):     func(ctx *context.Context, limit int) ast.Node { return AssignStmt(ctx, limit) },
		reflect.TypeOf(&ast.BasicLit{}):       func(ctx *context.Context, limit int) ast.Node { return BasicLit(ctx, limit) },
		reflect.TypeOf(&ast.BinaryExpr{}):     func(ctx *context.Context, limit int) ast.Node { return BinaryExpr(ctx, limit) },
		reflect.TypeOf(&ast.BlockStmt{}):      func(ctx *context.Context, limit int) ast.Node { return BlockStmt(ctx, limit) },
		reflect.TypeOf(&ast.BranchStmt{}):     func(ctx *context.Context, limit int) ast.Node { return BranchStmt(ctx, limit) },
		reflect.TypeOf(&ast.CallExpr{}):       func(ctx *context.Context, limit int) ast.Node { return CallExpr(ctx, limit) },
		reflect.TypeOf(&ast.CaseClause{}):     func(ctx *context.Context, limit int) ast.Node { return CaseClause(ctx, limit) },
		reflect.TypeOf(&ast.ChanType{}):       func(ctx *context.Context, limit int) ast.Node { return ChanType(ctx, limit) },
		reflect.TypeOf(&ast.CommClause{}):     func(ctx *context.Context, limit int) ast.Node { return CommClause(ctx, limit) },
		reflect.TypeOf(&ast.CompositeLit{}):   func(ctx *context.Context, limit int) ast.Node { return CompositeLit(ctx, limit) },
		reflect.TypeOf(&ast.DeclStmt{}):       func(ctx *context.Context, limit int) ast.Node { return DeclStmt(ctx, limit) },
		reflect.TypeOf(&ast.DeferStmt{}):      func(ctx *context.Context, limit int) ast.Node { return DeferStmt(ctx, limit) },
		reflect.TypeOf(&ast.Ellipsis{}):       func(ctx *context.Context, limit int) ast.Node { return Ellipsis(ctx, limit) },
		reflect.TypeOf(&ast.EmptyStmt{}):      func(ctx *context.Context, limit int) ast.Node { return EmptyStmt(ctx, limit) },
		reflect.TypeOf(&ast.ExprStmt{}):       func(ctx *context.Context, limit int) ast.Node { return ExprStmt(ctx, limit) },
		reflect.TypeOf(&ast.Field{}):          func(ctx *context.Context, limit int) ast.Node { return Field(ctx, limit) },
		reflect.TypeOf(&ast.FieldList{}):      func(ctx *context.Context, limit int) ast.Node { return FieldList(ctx, limit) },
		reflect.TypeOf(&ast.ForStmt{}):        func(ctx *context.Context, limit int) ast.Node { return ForStmt(ctx, limit) },
		reflect.TypeOf(&ast.FuncDecl{}):       func(ctx *context.Context, limit int) ast.Node { return FuncDecl(ctx, limit) },
		reflect.TypeOf(&ast.FuncLit{}):        func(ctx *context.Context, limit int) ast.Node { return FuncLit(ctx, limit) },
		reflect.TypeOf(&ast.FuncType{}):       func(ctx *context.Context, limit int) ast.Node { return FuncType(ctx, limit) },
		reflect.TypeOf(&ast.GenDecl{}):        func(ctx *context.Context, limit int) ast.Node { return GenDecl(ctx, limit) },
		reflect.TypeOf(&ast.GoStmt{}):         func(ctx *context.Context, limit int) ast.Node { return GoStmt(ctx, limit) },
		reflect.TypeOf(&ast.Ident{}):          func(ctx *context.Context, limit int) ast.Node { return Ident(ctx, limit) },
		reflect.TypeOf(&ast.IfStmt{}):         func(ctx *context.Context, limit int) ast.Node { return IfStmt(ctx, limit) },
		reflect.TypeOf(&ast.ImportSpec{}):     func(ctx *context.Context, limit int) ast.Node { return ImportSpec(ctx, limit) },
		reflect.TypeOf(&ast.IncDecStmt{}):     func(ctx *context.Context, limit int) ast.Node { return IncDecStmt(ctx, limit) },
		reflect.TypeOf(&ast.IndexExpr{}):      func(ctx *context.Context, limit int) ast.Node { return IndexExpr(ctx, limit) },
		reflect.TypeOf(&ast.IndexListExpr{}):  func(ctx *context.Context, limit int) ast.Node { return IndexListExpr(ctx, limit) },
		reflect.TypeOf(&ast.InterfaceType{}):  func(ctx *context.Context, limit int) ast.Node { return InterfaceType(ctx, limit) },
		reflect.TypeOf(&ast.KeyValueExpr{}):   func(ctx *context.Context, limit int) ast.Node { return KeyValueExpr(ctx, limit) },
		reflect.TypeOf(&ast.LabeledStmt{}):    func(ctx *context.Context, limit int) ast.Node { return LabeledStmt(ctx, limit) },
		reflect.TypeOf(&ast.MapType{}):        func(ctx *context.Context, limit int) ast.Node { return MapType(ctx, limit) },
		reflect.TypeOf(&ast.ParenExpr{}):      func(ctx *context.Context, limit int) ast.Node { return ParenExpr(ctx, limit) },
		reflect.TypeOf(&ast.RangeStmt{}):      func(ctx *context.Context, limit int) ast.Node { return RangeStmt(ctx, limit) },
		reflect.TypeOf(&ast.ReturnStmt{}):     func(ctx *context.Context, limit int) ast.Node { return ReturnStmt(ctx, limit) },
		reflect.TypeOf(&ast.SelectorExpr{}):   func(ctx *context.Context, limit int) ast.Node { return SelectorExpr(ctx, limit) },
		reflect.TypeOf(&ast.SelectStmt{}):     func(ctx *context.Context, limit int) ast.Node { return SelectStmt(ctx, limit) },
		reflect.TypeOf(&ast.SendStmt{}):       func(ctx *context.Context, limit int) ast.Node { return SendStmt(ctx, limit) },
		reflect.TypeOf(&ast.SliceExpr{}):      func(ctx *context.Context, limit int) ast.Node { return SliceExpr(ctx, limit) },
		reflect.TypeOf(&ast.StarExpr{}):       func(ctx *context.Context, limit int) ast.Node { return StarExpr(ctx, limit) },
		reflect.TypeOf(&ast.StructType{}):     func(ctx *context.Context, limit int) ast.Node { return StructType(ctx, limit) },
		reflect.TypeOf(&ast.SwitchStmt{}):     func(ctx *context.Context, limit int) ast.Node { return SwitchStmt(ctx, limit) },
		reflect.TypeOf(&ast.TypeAssertExpr{}): func(ctx *context.Context, limit int) ast.Node { return TypeAssertExpr(ctx, limit) },
		reflect.TypeOf(&ast.TypeSpec{}):       func(ctx *context.Context, limit int) ast.Node { return TypeSpec(ctx, limit) },
		reflect.TypeOf(&ast.TypeSwitchStmt{}): func(ctx *context.Context, limit int) ast.Node { return TypeSwitchStmt(ctx, limit) },
		reflect.TypeOf(&ast.UnaryExpr{}):      func(ctx *context.Context, limit int) ast.Node { return UnaryExpr(ctx, limit) },
		reflect.TypeOf(&ast.ValueSpec{}):      func(ctx *context.Context, limit int) ast.Node { return ValueSpec(ctx, limit) },
	}

	var s = struct {
		Spec ast.Spec
		Decl ast.Decl
		Expr ast.Expr
		Stmt ast.Stmt
	}{}

	xSpec, _ := reflect.TypeOf(s).FieldByName("Spec")
	xDecl, _ := reflect.TypeOf(s).FieldByName("Decl")
	xExpr, _ := reflect.TypeOf(s).FieldByName("Expr")
	xStmt, _ := reflect.TypeOf(s).FieldByName("Stmt")

	InterfaceConstructorsByReflectType = map[reflect.Type]func(ctx *context.Context, limit int) ast.Node{
		xSpec.Type: func(ctx *context.Context, limit int) ast.Node { return Spec(ctx, limit) },
		xDecl.Type: func(ctx *context.Context, limit int) ast.Node { return Decl(ctx, limit) },
		xExpr.Type: func(ctx *context.Context, limit int) ast.Node { return Expr(ctx, limit) },
		xStmt.Type: func(ctx *context.Context, limit int) ast.Node { return Stmt(ctx, limit) },
	}
}
