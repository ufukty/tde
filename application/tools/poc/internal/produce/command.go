package produce

import (
	astwutl "tde/internal/astw/astwutl"
	"tde/internal/command"
	"tde/internal/evolution"
	"tde/internal/folders/discovery"
	"tde/internal/folders/evaluation"
	"tde/internal/folders/inject"
	"tde/internal/folders/list"
	"tde/internal/folders/slotmgr"

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

func NewEvolutionTarget(modulePath string, packagePath string, pkgInfo *list.Package, funcName string) (*evolution.Target, error) {
	_, pkgs, err := astwutl.LoadDir(packagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create ast representation for the directory %q: %w", packagePath, err)
	}
	var pkgAst, ok = pkgs[pkgInfo.ImportPath]
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

func (c *Command) Run() {
	modPath, pkgInMod, pkg, err := discovery.WhereAmI()
	if err != nil {
		log.Fatalln("Could not find module root or package import path. Are you in a Go package and in subdir of a Go module?", err)
	}

	prepPath, err := inject.WithCreatingSample(modPath, pkgInMod, pkg, c.TestName)
	if err != nil {
		log.Fatalln("Could not prepare the module", err)
	}

	testDetails, err := discovery.FindTest(c.TestName)
	if err != nil {
		log.Fatalln("Could not find test details")
	}

	evolutionTarget, err := NewEvolutionTarget(modPath, pkgInMod, pkg, testDetails.ImplFuncName)
	if err != nil {
		log.Fatalln("Failed in slot_manager.NewSession()", err)
	}

	var session = slotmgr.New(prepPath, testDetails)
	var evaluator = evaluation.NewEvaluator(session)
	var evolution = evolution.NewManager(evolutionTarget)

	evolution.InitPopulation(c.Population)

	for i := 0; i < c.Iterate; i++ {
		fmt.Printf("Iteration: %d\n", i)
		evolution.IterateLoop()
	}

	evaluator.Sm.Print()
}
