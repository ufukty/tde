package node_constructor

import (
	"go/ast"
	"tde/internal/cfg/astcfg/context"
)

var AllowedPackagesToImport = []string{"fmt", "strings", "math"}

var ExpressionConstructors []func(context.Context, int) ast.Expr
var StatementConstructors []func(context.Context, int) ast.Stmt
var SpecificationConstructors []func(context.Context, int) ast.Spec
var DeclarationConstructors []func(context.Context, int) ast.Decl
var TypeDeclarationConstructors []func(context.Context, int) ast.Expr

func init() {
	ExpressionConstructors = []func(context.Context, int) ast.Expr{
		func(ctx context.Context, limit int) ast.Expr { return ArrayType(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return BasicLit(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return BinaryExpr(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return CallExpr(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return ChanType(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return CompositeLit(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return Ellipsis(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return FuncLit(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return FuncType(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return Ident(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return IndexExpr(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return IndexListExpr(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return InterfaceType(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return KeyValueExpr(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return MapType(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return ParenExpr(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return SelectorExpr(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return SliceExpr(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return StarExpr(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return StructType(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return TypeAssertExpr(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return UnaryExpr(ctx, limit) },
	}

	StatementConstructors = []func(context.Context, int) ast.Stmt{
		func(ctx context.Context, limit int) ast.Stmt { return AssignStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return BlockStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return BranchStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return CaseClause(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return CommClause(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return DeclStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return DeferStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return EmptyStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return ExprStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return ForStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return GoStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return IfStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return IncDecStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return LabeledStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return RangeStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return ReturnStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return SelectStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return SendStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return SwitchStmt(ctx, limit) },
		func(ctx context.Context, limit int) ast.Stmt { return TypeSwitchStmt(ctx, limit) },
	}

	SpecificationConstructors = []func(context.Context, int) ast.Spec{
		func(ctx context.Context, limit int) ast.Spec { return ImportSpec(ctx, limit) },
		func(ctx context.Context, limit int) ast.Spec { return TypeSpec(ctx, limit) },
		func(ctx context.Context, limit int) ast.Spec { return ValueSpec(ctx, limit) },
	}

	DeclarationConstructors = []func(context.Context, int) ast.Decl{
		func(ctx context.Context, limit int) ast.Decl { return FuncDecl(ctx, limit) },
		func(ctx context.Context, limit int) ast.Decl { return GenDecl(ctx, limit) },
	}

	TypeDeclarationConstructors = []func(context.Context, int) ast.Expr{
		func(ctx context.Context, limit int) ast.Expr { return ArrayType(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return ChanType(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return FuncType(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return InterfaceType(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return MapType(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return StructType(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return IdentType(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return ParenExpr(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return SelectorExpr(ctx, limit) },
		func(ctx context.Context, limit int) ast.Expr { return StarExpr(ctx, limit) },
	}
}
