package produce

import (
	"tde/internal/command"
	"tde/internal/evaluation"
	"tde/internal/evolution"
	"tde/internal/folders/discovery"
	"tde/internal/folders/preparation"
	"tde/internal/folders/slot_manager"
	"tde/internal/folders/types"
	"tde/internal/utilities"
	"tde/models/common_models"

	"fmt"

	"github.com/davecgh/go-spew/spew"
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

func (c *Command) Run() {
	modPath, pkgInMod, importPath, err := discovery.WhereAmI()
	if err != nil {
		utilities.Terminate("Could not find module root or package import path. Are you in a Go package and in subdir of a Go module?", err)
	}

	prepPath, err := preparation.Prepare(types.AbsolutePath(modPath), types.InModulePath(pkgInMod), importPath, c.TestName)
	if err != nil {
		utilities.Terminate("Could not prepare the module", err)
	}

	testDetails, err := discovery.ResolveTestDetailsInCurrentDir(c.TestName)
	if err != nil {
		utilities.Terminate("Could not find test details")
	}

	evolutionTarget, err := common_models.NewEvolutionTarget(types.AbsolutePath(modPath), types.InModulePath(pkgInMod), importPath, testDetails.ImplFuncName)
	if err != nil {
		utilities.Terminate("Failed in slot_manager.NewSession()", err)
	}

	var session = slot_manager.NewSession(prepPath, testDetails)
	var evaluator = evaluation.NewEvaluator(session)
	var evolution = evolution.NewEvolutionManager(evolutionTarget)

	evolution.InitPopulation(c.Population)

	for i := 0; i < c.Iterate; i++ {
		fmt.Printf("Iteration: %d\n", i)
		evolution.IterateLoop()
	}

	spew.Dump(evaluator.SlotManagerSession)
}
