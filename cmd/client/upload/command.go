package upload

import (
	"flag"
	"fmt"
	"os"
	"tde/internal/command"

	"github.com/pkg/errors"
)

type Command struct {
	ExcludeDirs command.MultiString
}

func (c *Command) BindFlagsToArgs(f *flag.FlagSet) {}

func (c *Command) ManualParser(f *flag.FlagSet) error {
	f.Var(&c.ExcludeDirs, "exclude-dir", "")

	err := f.Parse(os.Args[2:])
	if err != nil {
		return errors.Wrap(err, "Could not parse arguments. Run \"tde help\"")
	}

	return nil
}

func (c *Command) Run() {
	for i, dir := range c.ExcludeDirs {
		fmt.Println("i:", i, "dir:", dir)
	}
}
