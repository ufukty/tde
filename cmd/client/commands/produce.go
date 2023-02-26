package commands

import (
	"flag"
	"fmt"
	"os"
)

type Produce struct {
	fs *flag.FlagSet

	TestName string

	IncludeModule bool
	RunnerAddress string

	Timeout int

	// Population       int
	// Generation       int
	// SizeLimit        int
}

func NewProduce() *Produce {
	produce := Produce{
		fs: flag.NewFlagSet("produce", flag.ExitOnError),
	}
	return &produce
}

func (p *Produce) Setup() {
	p.fs.Usage = func() {
		fmt.Print(`
Always run produce command on directory contains the package.

Usage:

	tde produce [ -include-module ] TestName

Arguments:

`)
		p.fs.PrintDefaults()
	}

	p.IncludeModule = *p.fs.Bool(
		"include-module", false, "(Optional) Enable full module upload. It is necessary if the package contains imports from rest of the module.")
	p.RunnerAddress = *p.fs.String(
		"runner-address", "", "(Optional) Not needed when tde's runners wanted to use by user for compiling and running candidates. Needed for using custom runner. IP address and port in format: 127.0.0.1.26900")
	p.Timeout = *p.fs.Int(
		"timeout", 10, "Minutes to wait before terminate request.")

	// config.Population = *flag.Int(
	// 	"population", 1000, "Number of candidates generated at start and tested at each iteration at each generation")
	// config.Generation = *flag.Int(
	// 	"generation", 10, "Number of generations which the evolution will be stopped.")
	// config.SizeLimit = *flag.Int(
	// 	"size-limit", 1000, "Character size limit for any candidate.")

}

func (p *Produce) Parse() int {
	err := p.fs.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("Could not parse arguments. Run \"tde help\"")
		return 1
	}

	tailingArguments := p.fs.Args()
	if len(tailingArguments) == 0 {
		fmt.Println("Test name is missing. Run \"tde help\"")
		return 1
	}
	p.TestName = tailingArguments[0]

	return 0
}

func (p *Produce) PrintGenericHelp() {
	p.fs.Usage()
}

func (p *Produce) Run() int {
	fmt.Println(p.TestName)
	return 0
}
