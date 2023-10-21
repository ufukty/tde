package produce

import (
	"log"
	"path/filepath"
	"tde/internal/command"
	"tde/internal/evolution"
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/evaluation/discovery"
	"tde/internal/evolution/evaluation/inject"
	"tde/internal/evolution/evaluation/slotmgr"
	"tde/internal/utilities"
	models "tde/models/program"

	"fmt"
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
	Dc         int                 `short:"dc"`                             // depth limit for code search
	Dp         int                 `short:"dp"`                             // depth limit for program search
	Ds         int                 `short:"ds"`                             // depth limit for solution search
	Package    command.MultiString `long:"package" short:"p"`               // packages allowed to import
	Exclude    command.MultiString `long:"exclude" short:"e"`               // TODO:
	TestName   string              `precedence:"0"`
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
	prepPath, err := inject.WithCreatingSample(mod, pkg, c.TestName)
	if err != nil {
		log.Fatalln("Could not prepare the module:", err)
	}
	sm := slotmgr.New(prepPath, combined.Package.PathInModule(), filepath.Base(combined.Target.Path))
	sm.Print()
	ctx, err := models.LoadContext(mod, pkg.PathInModule(), combined.Target.Name)
	if err != nil {
		log.Fatalln("Could not load the context for package:", err)
	}
	evaluator := evaluation.NewEvaluator(sm, ctx)
	evolution := evolution.NewSolutionSearch(evaluator, &models.Parameters{}, ctx)

	if err := evolution.Loop(); err != nil {
		log.Fatalln("Could not complete the process:", err)
	}
}
