package produce

import (
	"flag"

	"github.com/kr/pretty"
)

type Args struct {
	Continue   int
	Iterate    int
	Model      string
	Package    string
	Population int
	Ratios     string
	Runner     string
	Size       int
	TestName   string
	Timeout    int
}

func Run() error {
	args := &Args{}
	flag.IntVar(&args.Continue, "continue", 10, "")
	flag.IntVar(&args.Iterate, "iterate", 10, "")
	flag.IntVar(&args.Population, "population", 1000, "")
	flag.IntVar(&args.Size, "size", 1000, "")
	flag.IntVar(&args.Timeout, "timeout", 10, "")
	flag.StringVar(&args.Model, "model", "0.1", "")
	flag.StringVar(&args.Package, "package", "", "")
	flag.StringVar(&args.Ratios, "ratios", "10/1", "")
	flag.StringVar(&args.Runner, "runner", "", "")
	flag.StringVar(&args.TestName, "testname", "", "")
	flag.Parse()

	pretty.Println(args)

	return nil
}
