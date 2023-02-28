package produce

import (
	"flag"
	"os"

	"github.com/pkg/errors"
)

type Command struct {
	Timeout       int
	TestName      string
	RunnerAddress string
	IncludeModule bool
	// Population       int
	// Generation       int
	// SizeLimit        int
}

func (c *Command) BindFlagsToArgs(f *flag.FlagSet) {
	c = &Command{
		Timeout:       *f.Int("timeout", 10, ""),
		IncludeModule: *f.Bool("include-module", false, ""),
		RunnerAddress: *f.String("runner-address", "", ""),
		// Population : *f.Int("population", 1000, "")
		// Generation : *f.Int("generation", 10, "")
		// SizeLimit  : *f.Int("size-limit", 1000, "")
	}
}

func (c *Command) ManualParser(f *flag.FlagSet) error {
	err := f.Parse(os.Args[2:])
	if err != nil {
		return errors.Wrap(err, "Could not parse arguments. Run \"tde help\"")
	}

	tailingArguments := f.Args()
	if len(tailingArguments) == 0 {
		return errors.Wrap(err, "Test name is missing. Run \"tde help\"")
	}
	c.TestName = tailingArguments[0]

	return nil
}

func (c *Command) Run() {

}
