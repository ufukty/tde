package code

import (
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
	fset    *token.FileSet
	astFile *ast.File
}

func (c *Code) LoadFromFile(path string) error {
	var (
		err error
	)
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

// func (c *Code) FindFunction(funcName string) (*ast.FuncDecl, error) {
// 	found := []*ast.FuncDecl{}

// 	ast.Inspect(c.astFile, func(n ast.Node) bool {
// 		if n != nil {
// 			if functionDeclaration, ok := n.(*ast.FuncDecl); ok {
// 				if functionDeclaration.Name.Name == funcName {
// 					found = append(found, functionDeclaration)
// 					return false
// 				}
// 			}
// 		}
// 		return true
// 	})

// 	if len(found) == 0 {
// 		return nil, fmt.Errorf("could not find function '%s'", funcName)
// 	} else if len(found) > 1 {
// 		return nil, fmt.Errorf("more than one function definition with the name  '%s'", funcName)
// 	} else {
// 		return found[0], nil
// 	}
// }

// func (c *Code) ReplaceFunctionName(n *ast.FuncDecl, newName string) {
// 	n.Name.Name = newName
// }

func (c *Code) InspectAST(f func(n ast.Node) bool) {
	ast.Inspect(c.astFile, f)
}

func (c *Code) PartialString(node any) string {
	sw := utilities.NewStringWriter()
	printer.Fprint(sw, c.fset, node)
	return sw.String()
}

func (c *Code) String() string {
	return c.PartialString(c.astFile)
}

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
	printer.Fprint(w, c.fset, c.astFile)
}

// Use this to refresh token positions after tree manipulation
// Performance notice: This will print&parse the tree, avoid using excessively
func (c *Code) Reload() error {
	return c.LoadFromString(c.String())
}
