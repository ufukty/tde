package node_construtor

import (
	"go/ast"
)

var AllowedPackagesToImport = []string{"fmt", "strings", "math"}

var ExpressionConstructors []func(int) ast.Expr
var StatementConstructors []func(int) ast.Stmt
var SpecificationConstructors []func(int) ast.Spec
var DeclarationConstructors []func(int) ast.Decl
var TypeDeclarationConstructors []func(int) ast.Expr

func init() {
	ExpressionConstructors = []func(int) ast.Expr{
		func(limit int) ast.Expr { return ArrayType(limit) },
		func(limit int) ast.Expr { return BasicLit(limit) },
		func(limit int) ast.Expr { return BinaryExpr(limit) },
		func(limit int) ast.Expr { return CallExpr(limit) },
		func(limit int) ast.Expr { return ChanType(limit) },
		func(limit int) ast.Expr { return CompositeLit(limit) },
		func(limit int) ast.Expr { return Ellipsis(limit) },
		func(limit int) ast.Expr { return FuncLit(limit) },
		func(limit int) ast.Expr { return FuncType(limit) },
		func(limit int) ast.Expr { return Ident(limit) },
		func(limit int) ast.Expr { return IndexExpr(limit) },
		func(limit int) ast.Expr { return IndexListExpr(limit) },
		func(limit int) ast.Expr { return InterfaceType(limit) },
		func(limit int) ast.Expr { return KeyValueExpr(limit) },
		func(limit int) ast.Expr { return MapType(limit) },
		func(limit int) ast.Expr { return ParenExpr(limit) },
		func(limit int) ast.Expr { return SelectorExpr(limit) },
		func(limit int) ast.Expr { return SliceExpr(limit) },
		func(limit int) ast.Expr { return StarExpr(limit) },
		func(limit int) ast.Expr { return StructType(limit) },
		func(limit int) ast.Expr { return TypeAssertExpr(limit) },
		func(limit int) ast.Expr { return UnaryExpr(limit) },
	}

	StatementConstructors = []func(int) ast.Stmt{
		func(limit int) ast.Stmt { return AssignStmt(limit) },
		func(limit int) ast.Stmt { return BlockStmt(limit) },
		func(limit int) ast.Stmt { return BranchStmt(limit) },
		func(limit int) ast.Stmt { return CaseClause(limit) },
		func(limit int) ast.Stmt { return CommClause(limit) },
		func(limit int) ast.Stmt { return DeclStmt(limit) },
		func(limit int) ast.Stmt { return DeferStmt(limit) },
		func(limit int) ast.Stmt { return EmptyStmt(limit) },
		func(limit int) ast.Stmt { return ExprStmt(limit) },
		func(limit int) ast.Stmt { return ForStmt(limit) },
		func(limit int) ast.Stmt { return GoStmt(limit) },
		func(limit int) ast.Stmt { return IfStmt(limit) },
		func(limit int) ast.Stmt { return IncDecStmt(limit) },
		func(limit int) ast.Stmt { return LabeledStmt(limit) },
		func(limit int) ast.Stmt { return RangeStmt(limit) },
		func(limit int) ast.Stmt { return ReturnStmt(limit) },
		func(limit int) ast.Stmt { return SelectStmt(limit) },
		func(limit int) ast.Stmt { return SendStmt(limit) },
		func(limit int) ast.Stmt { return SwitchStmt(limit) },
		func(limit int) ast.Stmt { return TypeSwitchStmt(limit) },
	}

	SpecificationConstructors = []func(int) ast.Spec{
		func(limit int) ast.Spec { return ImportSpec(limit) },
		func(limit int) ast.Spec { return TypeSpec(limit) },
		func(limit int) ast.Spec { return ValueSpec(limit) },
	}

	DeclarationConstructors = []func(int) ast.Decl{
		func(limit int) ast.Decl { return FuncDecl(limit) },
		func(limit int) ast.Decl { return GenDecl(limit) },
	}

	TypeDeclarationConstructors = []func(int) ast.Expr{
		func(limit int) ast.Expr { return ArrayType(limit) },
		func(limit int) ast.Expr { return ChanType(limit) },
		func(limit int) ast.Expr { return FuncType(limit) },
		func(limit int) ast.Expr { return InterfaceType(limit) },
		func(limit int) ast.Expr { return MapType(limit) },
		func(limit int) ast.Expr { return StructType(limit) },
		func(limit int) ast.Expr { return IdentType(limit) },
		func(limit int) ast.Expr { return ParenExpr(limit) },
		func(limit int) ast.Expr { return SelectorExpr(limit) },
		func(limit int) ast.Expr { return StarExpr(limit) },
	}
}
