package discovery

import (
	"fmt"
	"go/ast"

	"tde/internal/evaluation/list"
	"tde/internal/utilities/indent"
)

type TargetFunction struct {
	Name      string
	Path      string
	LineStart int
	LineEnd   int
}

type TestFunction struct {
	Name      string
	Path      string
	LineStart int
	LineEnd   int
	Calls     []*ast.CallExpr
}

type CombinedDetails struct {
	Package *list.Package
	Target  *TargetFunction
	Test    *TestFunction
}

func (t TestFunction) String() string {
	r := ""
	r += fmt.Sprintln("Name      :", t.Name)
	r += fmt.Sprintln("Path      :", t.Path)
	r += fmt.Sprintln("LineStart :", t.LineStart)
	r += fmt.Sprintln("LineEnd   :", t.LineEnd)
	return r
}

func (t TargetFunction) String() string {
	r := ""
	r += fmt.Sprintln("Name      :", t.Name)
	r += fmt.Sprintln("Path      :", t.Path)
	r += fmt.Sprintln("LineStart :", t.LineStart)
	r += fmt.Sprintln("LineEnd   :", t.LineEnd)
	return r
}

func (c CombinedDetails) String() string {
	r := ""
	r += fmt.Sprintln("Package:")
	r += fmt.Sprintln(indent.Lines(c.Package.String(), 4))
	r += fmt.Sprintln("Test:")
	r += fmt.Sprintln(indent.Lines(c.Test.String(), 4))
	r += fmt.Sprintln("Target:")
	r += fmt.Sprintln(indent.Lines(c.Target.String(), 4))
	return r
}
