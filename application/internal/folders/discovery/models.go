package discovery

import (
	"fmt"
	"go/ast"
	"tde/internal/folders/list"
	"tde/internal/utilities"
)

type TargetFunction struct {
	Name string
	Path string
	Line int
}

type TestFunction struct {
	Name  string
	Path  string
	Line  int
	Calls []*ast.CallExpr
}

type CombinedDetails struct {
	Package *list.Package
	Target  *TargetFunction
	Test    *TestFunction
}

func (t TestFunction) String() string {
	r := ""
	r += fmt.Sprintln("Name :", t.Name)
	r += fmt.Sprintln("Path :", t.Path)
	r += fmt.Sprintln("Line :", t.Line)
	return r
}

func (t TargetFunction) String() string {
	r := ""
	r += fmt.Sprintln("Name :", t.Name)
	r += fmt.Sprintln("Path :", t.Path)
	r += fmt.Sprintln("Line :", t.Line)
	return r
}

func (c CombinedDetails) String() string {
	r := ""
	r += fmt.Sprintln("Package:")
	r += fmt.Sprintln(utilities.IndentLines(c.Package.String(), 4))
	r += fmt.Sprintln("Test:")
	r += fmt.Sprintln(utilities.IndentLines(c.Test.String(), 4))
	r += fmt.Sprintln("Target:")
	r += fmt.Sprintln(utilities.IndentLines(c.Target.String(), 4))
	return r
}
