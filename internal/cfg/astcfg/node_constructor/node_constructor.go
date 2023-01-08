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
		func(counter int) ast.Expr { return ArrayType(counter) },
		func(counter int) ast.Expr { return BasicLit(counter) },
		func(counter int) ast.Expr { return BinaryExpr(counter) },
		func(counter int) ast.Expr { return CallExpr(counter) },
		func(counter int) ast.Expr { return ChanType(counter) },
		func(counter int) ast.Expr { return CompositeLit(counter) },
		func(counter int) ast.Expr { return Ellipsis(counter) },
		func(counter int) ast.Expr { return FuncLit(counter) },
		func(counter int) ast.Expr { return FuncType(counter) },
		func(counter int) ast.Expr { return Ident(counter) },
		func(counter int) ast.Expr { return IndexExpr(counter) },
		func(counter int) ast.Expr { return IndexListExpr(counter) },
		func(counter int) ast.Expr { return InterfaceType(counter) },
		func(counter int) ast.Expr { return KeyValueExpr(counter) },
		func(counter int) ast.Expr { return MapType(counter) },
		func(counter int) ast.Expr { return ParenExpr(counter) },
		func(counter int) ast.Expr { return SelectorExpr(counter) },
		func(counter int) ast.Expr { return SliceExpr(counter) },
		func(counter int) ast.Expr { return StarExpr(counter) },
		func(counter int) ast.Expr { return StructType(counter) },
		func(counter int) ast.Expr { return TypeAssertExpr(counter) },
		func(counter int) ast.Expr { return UnaryExpr(counter) },
	}

	StatementConstructors = []func(int) ast.Stmt{
		func(counter int) ast.Stmt { return AssignStmt(counter) },
		func(counter int) ast.Stmt { return BlockStmt(counter) },
		func(counter int) ast.Stmt { return BranchStmt(counter) },
		func(counter int) ast.Stmt { return CaseClause(counter) },
		func(counter int) ast.Stmt { return CommClause(counter) },
		func(counter int) ast.Stmt { return DeclStmt(counter) },
		func(counter int) ast.Stmt { return DeferStmt(counter) },
		func(counter int) ast.Stmt { return EmptyStmt(counter) },
		func(counter int) ast.Stmt { return ExprStmt(counter) },
		func(counter int) ast.Stmt { return ForStmt(counter) },
		func(counter int) ast.Stmt { return GoStmt(counter) },
		func(counter int) ast.Stmt { return IfStmt(counter) },
		func(counter int) ast.Stmt { return IncDecStmt(counter) },
		func(counter int) ast.Stmt { return LabeledStmt(counter) },
		func(counter int) ast.Stmt { return RangeStmt(counter) },
		func(counter int) ast.Stmt { return ReturnStmt(counter) },
		func(counter int) ast.Stmt { return SelectStmt(counter) },
		func(counter int) ast.Stmt { return SendStmt(counter) },
		func(counter int) ast.Stmt { return SwitchStmt(counter) },
		func(counter int) ast.Stmt { return TypeSwitchStmt(counter) },
	}

	SpecificationConstructors = []func(int) ast.Spec{
		func(counter int) ast.Spec { return ImportSpec(counter) },
		func(counter int) ast.Spec { return TypeSpec(counter) },
		func(counter int) ast.Spec { return ValueSpec(counter) },
	}

	DeclarationConstructors = []func(int) ast.Decl{
		func(counter int) ast.Decl { return FuncDecl(counter) },
		func(counter int) ast.Decl { return GenDecl(counter) },
	}

	TypeDeclarationConstructors = []func(int) ast.Expr{
		func(counter int) ast.Expr { return ArrayType(counter) },
		func(counter int) ast.Expr { return ChanType(counter) },
		func(counter int) ast.Expr { return FuncType(counter) },
		func(counter int) ast.Expr { return InterfaceType(counter) },
		func(counter int) ast.Expr { return MapType(counter) },
		func(counter int) ast.Expr { return StructType(counter) },
		func(counter int) ast.Expr { return IdentType(counter) },
		func(counter int) ast.Expr { return ParenExpr(counter) },
		func(counter int) ast.Expr { return SelectorExpr(counter) },
		func(counter int) ast.Expr { return StarExpr(counter) },
	}
}
