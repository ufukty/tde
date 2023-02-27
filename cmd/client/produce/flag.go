package produce

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Flags struct {
	fs *flag.FlagSet

	TestName string

	IncludeModule bool
	RunnerAddress string

	Timeout int

	// Population       int
	// Generation       int
	// SizeLimit        int
}

func init() {
	
}

func NewProduce() *Flags {
	produce := Flags{
		fs: flag.NewFlagSet("produce", flag.ExitOnError),
	}
	return &produce
}

func (f *Flags) Setup() {
	f.fs.Usage = func() {
		fmt.Println(strings.Join([]string{
			"Always run produce command on directory contains the package.",
			"",
			"Usage:",
			"",
			"	tde produce [ -include-module ] TestName",
			"",
			"Arguments:",
		}, "\n"))
		f.fs.PrintDefaults()
	}

	f.IncludeModule = *f.fs.Bool(
		"include-module", false, "(Optional) Enable full module upload. It is necessary if the package contains imports from rest of the module.")
	f.RunnerAddress = *f.fs.String(
		"runner-address", "", "(Optional) Not needed when tde's runners wanted to use by user for compiling and running candidates. Needed for using custom runner. IP address and port in format: 127.0.0.1.26900")
	f.Timeout = *f.fs.Int(
		"timeout", 10, "Minutes to wait before terminate request.")

	// config.Population = *flag.Int(
	// 	"population", 1000, "Number of candidates generated at start and tested at each iteration at each generation")
	// config.Generation = *flag.Int(
	// 	"generation", 10, "Number of generations which the evolution will be stopped.")
	// config.SizeLimit = *flag.Int(
	// 	"size-limit", 1000, "Character size limit for any candidate.")

}

func (p *Flags) Parse() int {
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

func (f *Flags) PrintGenericHelp() {
	f.fs.Usage()
}

func (f *Flags) Run() {
	fmt.Println(f.TestName)
}
