package produce

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"tde/internal/evolution"
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/evaluation/discovery"
	"tde/internal/evolution/evaluation/inject"
	"tde/internal/evolution/evaluation/slotmgr"
	"tde/internal/evolution/models"
	"tde/internal/utilities/strw"
	"tde/internal/utilities/valuable"
)

type Args struct {
	Continue   int
	Dc         int
	Dp         int
	Ds         int
	Exclude    valuable.Strings
	Iterate    int
	Model      string
	Package    valuable.Strings
	Population int
	Ratios     string
	Runner     string
	Size       int
	TestName   string
	Timeout    int
}

func pipeline(args *Args) error {
	mod, pkg, err := discovery.WhereAmI()
	if err != nil {
		log.Fatalln("Could not get the details about Go module and package in this directory:", err)
	}
	combined, err := discovery.CombinedDetailsForTest(".", args.TestName)
	if err != nil {
		log.Fatalln("Could not find test details:", err)
	}
	fmt.Println("Detected values:")
	fmt.Println(strw.IndentLines(combined.String(), 4))
	prepPath, err := inject.WithCreatingSample(mod, pkg, args.TestName)
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

	return nil
}

func Run() error {
	args := &Args{}
	flag.IntVar(&args.Continue, "continue", 10, "session")
	flag.IntVar(&args.Dc, "dc", 0, "depth limit for code search")
	flag.IntVar(&args.Dp, "dp", 0, "depth limit for program search")
	flag.IntVar(&args.Ds, "ds", 0, "depth limit for solution search")
	flag.IntVar(&args.Iterate, "iterate", 10, "")
	flag.IntVar(&args.Population, "population", 1000, "")
	flag.IntVar(&args.Size, "size", 1000, "")
	flag.IntVar(&args.Timeout, "timeout", 10, "in seconds")
	flag.StringVar(&args.Model, "model", "0.1", "")
	flag.StringVar(&args.Ratios, "ratios", "10/1", "")
	flag.StringVar(&args.Runner, "runner", "", "ip address")
	flag.StringVar(&args.TestName, "0", "", "")
	flag.Var(&args.Exclude, "exclude", "")
	flag.Var(&args.Package, "package", "packages allowed to import")
	flag.Parse()
	return pipeline(args)
}
