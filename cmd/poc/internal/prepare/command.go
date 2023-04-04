package prepare

import (
	"context"
	"tde/internal/command"
	"tde/internal/evolution"
	"tde/internal/folders/discovery"
	"tde/internal/folders/preparation"
	"tde/internal/folders/slot_manager"
	"tde/internal/folders/types"
	"tde/internal/utilities"

	"github.com/kr/pretty"
)

type Command struct {
	N        int                 `short:"n" default:"10"`
	Exclude  command.MultiString `long:"exclude" short:"e"` // TODO:
	TestName string              `precedence:"0"`

	Timeout    int                 `long:"timeout" default:"10"`
	Runner     string              `long:"runner"`
	Continue   string              `long:"continue" short:"c" default:"10"`
	Model      string              `long:"model" default:"0.1"`
	Ratios     string              `long:"ratios" default:"10/1"`
	Population int                 `long:"population" default:"1000"`
	Iterate    int                 `long:"iterate" default:"10"`
	Size       int                 `long:"size" default:"1000"`
	Package    command.MultiString `long:"package" short:"p"` // packages allowed to import
}

func (c *Command) Run() {
	modPath, pkgInMod, importPath, err := discovery.WhereAmI()
	if err != nil {
		utilities.Terminate("Could not find module root or package import path. Are you in a Go package and in subdir of a Go module?", err)
	}

	prepPath, err := preparation.Prepare(types.AbsolutePath(modPath), types.InModulePath(pkgInMod), importPath, c.TestName)
	if err != nil {
		utilities.Terminate("Could not prepare the module", err)
	}

	pretty.Println(c)
	pretty.Println(prepPath)

	testDetails, err := discovery.ResolveTestDetailsInCurrentDir(c.TestName)
	if err != nil {
		utilities.Terminate("Could not find test details")
	}

	var config = &types.TestDetails{
		PackagePath:   types.InModulePath(pkgInMod),
		PackageImport: importPath,
		ImplFuncFile:  testDetails.ImplFuncFile,
		TestFuncFile:  testDetails.TestFuncFile,
		TestFuncName:  c.TestName,
	}
	var session = slot_manager.NewSession(prepPath, config)
	var evolution = evolution.NewEvolution(session)
	evolution.InitPopulation(c.N)

	evolution.IterateLoop(context.Background())

	pretty.Print(evolution.Evaluation.SlotManagerSession)
}
