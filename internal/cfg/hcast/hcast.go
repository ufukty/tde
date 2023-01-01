package hcast

import "go/ast"

type File *ast.File
type Scope *ast.BlockStmt
type Offset int

type Declaration struct {
	Ident  *ast.DeclStmt
	Scope  Scope  // Declared scope
	Offset Offset // Child's offset within its parent
}

type Blocks struct {
	Parent *ast.Block
}

func NewVariable(file File, scope Scope, line Offset)    {}
func VariableScopeUp(decl Declaration)                   {}
func VariableScopeDown(decl Declaration)                 {}
func EmbedInNewFunction(children []ast.Node)             {}
func MoveIntoExistingFunction()                          {}
func GenerateErrorCheckBlock(err *ast.Ident) *ast.IfStmt { return &ast.IfStmt{} }
func ConditionalizeScope() *ast.IfStmt                   { return &ast.IfStmt{} }
func IterateOverArray(ds *ast.Ident)                     { return &ast.ForStmt{Init: ast.} }
