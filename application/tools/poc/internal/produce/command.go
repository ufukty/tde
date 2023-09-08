package produce

import (
	"tde/internal/astw/astwutl"
	"tde/internal/command"
	"tde/internal/evolution"
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/evaluation/discovery"
	"tde/internal/evolution/evaluation/inject"
	"tde/internal/evolution/evaluation/list"
	"tde/internal/evolution/evaluation/slotmgr"
	"tde/internal/utilities"

	"fmt"
	"log"

	"golang.org/x/exp/maps"
)

type Command struct {
	Timeout    int                 `long:"timeout" default:"10"`            // in seconds
	Runner     string              `long:"runner"`                          // ip address
	Continue   string              `long:"continue" short:"c" default:"10"` // session
	Model      string              `long:"model" default:"0.1"`             //
	Ratios     string              `long:"ratios" default:"10/1"`           //
	Population int                 `long:"population" default:"1000"`       //
	Iterate    int                 `long:"iterate" default:"10"`            //
	Size       int                 `long:"size" default:"1000"`             //
	Package    command.MultiString `long:"package" short:"p"`               // packages allowed to import
	Exclude    command.MultiString `long:"exclude" short:"e"`               // TODO:
	TestName   string              `precedence:"0"`
}

func newEvolutionTarget(modulePath string, pkgInfo *list.Package, funcName string) (*evolution.Target, error) {
	_, pkgs, err := astwutl.LoadDir(".")
	if err != nil {
		return nil, fmt.Errorf("failed on parsing files in current dir: %w", err)
	}
	var pkgAst, ok = pkgs[pkgInfo.Name]
	if !ok {
		return nil, fmt.Errorf("directory doesn't contain named package %q. Available packages are %v", pkgInfo.ImportPath, maps.Keys(pkgs))
	}
	fileAst, funcDeclAst, err := astwutl.FindFuncDeclInPkg(pkgAst, funcName)
	if err != nil {
		return nil, fmt.Errorf("directory doesn't contain named function %q: %w", funcName, err)
	}
	return &evolution.Target{
		Package:  pkgAst,
		File:     fileAst,
		FuncDecl: funcDeclAst,
	}, nil
}

func printDetails() {
	fmt.Println()
}

func (c *Command) Run() {
	mod, pkg, err := discovery.WhereAmI()
	if err != nil {
		log.Fatalln("Could not get the details about Go module and package in this directory:", err)
	}
	combined, err := discovery.CombinedDetailsForTest(".", c.TestName)
	if err != nil {
		log.Fatalln("Could not find test details:", err)
	}
	fmt.Println("Detected values:")
	fmt.Println(utilities.IndentLines(combined.String(), 4))

	evolutionTarget, err := newEvolutionTarget(mod, pkg, combined.Target.Name)
	if err != nil {
		log.Fatalln("Could not create evolution target:", err)
	}
	prepPath, err := inject.WithCreatingSample(mod, pkg, c.TestName)
	if err != nil {
		log.Fatalln("Could not prepare the module:", err)
	}

	var sm = slotmgr.New(prepPath, combined)
	sm.Print()

	var evaluator = evaluation.NewEvaluator(sm)
	var evolution = evolution.NewManager(evolutionTarget)

	evolution.InitPopulation(c.Population)

	for i := 0; i < c.Iterate; i++ {
		fmt.Printf("Iteration: %d\n", i)
		evolution.IterateLoop()
		evaluator.Pipeline(maps.Values(evolution.Candidates)) // TODO:
	}

}
