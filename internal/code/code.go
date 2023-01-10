package code

import (
	"tde/internal/ast_wrapper"
	"tde/internal/utilities"

	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

type Code struct {
	fset                      *token.FileSet
	astFile                   *ast.File
	overwrittenFunctionBodies map[*ast.FuncDecl]string
}

func NewCode() *Code {
	return &Code{
		overwrittenFunctionBodies: map[*ast.FuncDecl]string{},
	}
}

func (c *Code) LoadFromFile(path string) error {
	var err error
	c.fset = token.NewFileSet()
	c.astFile, err = parser.ParseFile(c.fset, path, nil, parser.AllErrors)
	if err != nil {
		return errors.Wrap(err, "could not parse file")
	}
	return nil
}

func (c *Code) LoadFromString(content string) error {
	var (
		err error
	)
	c.fset = token.NewFileSet()
	c.astFile, err = parser.ParseFile(c.fset, "", content, parser.AllErrors)
	if err != nil {
		return errors.Wrap(err, "could not parse string")
	}
	return nil
}

func (c *Code) RenameFunction(oldName, newName string) error {
	found := false
	ast.Inspect(c.astFile, func(n ast.Node) bool {
		if n != nil {
			if fd, ok := n.(*ast.FuncDecl); ok {
				if fd.Name.Name == oldName {
					fd.Name.Name = newName
					found = true
				}
				return false // no nested-function-defs anyways
			}
		}
		return !found // continue DFS if not found
	})
	if !found {
		return fmt.Errorf("no function named %s in the tree", oldName)
	}
	return nil
}

func (c *Code) FindFunction(funcName string) (*ast.FuncDecl, error) {
	return ast_wrapper.FindFuncDecl(c.astFile, funcName)
}

func (c *Code) ReplaceFunctionName(n *ast.FuncDecl, newName string) {
	n.Name.Name = newName
}

func (c *Code) InspectAST(f func(n ast.Node) bool) {
	ast.Inspect(c.astFile, f)
}

func (c *Code) InspectExceptRoot(f func(n ast.Node) bool) {
	ast.Inspect(c.astFile, func(n ast.Node) bool {
		if n == c.astFile {
			return true
		}
		return f(n)
	})
}
func (c *Code) PartialString(node ast.Node) string {
	sw := utilities.NewStringWriter()
	printer.Fprint(sw, c.fset, node)
	export := sw.String()
	insertedCharacters := 0
	ast.Inspect(node, func(n ast.Node) bool {
		if n, ok := n.(*ast.FuncDecl); ok {
			if overwrittenBody, ok := c.overwrittenFunctionBodies[n]; ok {
				export = export[:int(n.Body.Lbrace)+insertedCharacters] + " " + overwrittenBody + " " + export[int(n.Body.Rbrace)+insertedCharacters-1:]
				insertedCharacters += len(overwrittenBody)
			}
		}
		return true
	})
	return export
}

func (c *Code) String() string {
	return c.PartialString(c.astFile)
}

// ignores overwritten bodies
func (c *Code) PrintListOfTokens() {
	ast.Inspect(c.astFile, func(n ast.Node) bool {
		if n != nil {
			repr := strings.ReplaceAll(c.PartialString(n), "\n", "\\n")
			fmt.Printf("[%d:%d] %-20s \t %s\n", n.Pos(), n.End(), reflect.TypeOf(n), repr)
		}
		return true
	})
}

func (c *Code) Print(w io.Writer) {
	// printer.Fprint(w, c.fset, c.astFile)
	fmt.Print(c.String())
}

// Use this to refresh token positions after tree manipulation
// Performance notice: This will print&parse the tree, avoid using excessively
func (c *Code) Reload() error {
	return c.LoadFromString(c.String())
}

func (c *Code) OverwriteFunction(f *ast.FuncDecl, content string) {
	c.overwrittenFunctionBodies[f] = content
}

func (c *Code) LineNumberOfPosition(pos token.Pos) int {
	return c.fset.Position(pos).Line
}
