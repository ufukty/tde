package produce

import (
	"flag"
	"os"
	"tde/internal/command"

	"github.com/kr/pretty"
	"github.com/pkg/errors"
)

type Command struct {
	Timeout       int
	TestName      string
	RunnerAddress string
	IncludeModule bool

	Population int
	Generation int
	Size       int
	Model      string
	Ratios     string
	Iterate    int
	Package    command.MultiString
}

func (c *Command) BindFlagsToArgs(f *flag.FlagSet) {
	c = &Command{
		Timeout:       *f.Int("timeout", 10, ""),
		IncludeModule: *f.Bool("include-module", false, ""),
		Population:    *f.Int("population", 1000, ""),
		Generation:    *f.Int("generation", 10, ""),
		Size:          *f.Int("size", 1000, ""),
		Iterate:       *f.Int("iterate", 10, ""),
		RunnerAddress: *f.String("runner-address", "", ""),
		Model:         *f.String("model", "0.1", ""),
		Ratios:        *f.String("ratios", "10/1", ""),
	}
}

func (c *Command) ManualParser(f *flag.FlagSet) error {
	f.Var(&c.Package, "package", "")

	err := f.Parse(os.Args[2:])
	if err != nil {
		return errors.Wrap(err, "Could not parse arguments. Run \"tde help produce\"")
	}

	switch tailingArguments := f.Args(); len(tailingArguments) {
	case 0:
		return errors.Wrap(err, "Test name is missing. Run \"tde help produce\"")
	case 1:
		c.TestName = tailingArguments[0]
	default:
		return errors.Wrap(err, "Could not parse arguments. More than one tailing arguments are found. Run \"tde help produce\"")
	}

	return nil
}

func (c *Command) Run() {
	pretty.Println(c)
}
