package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"

	"tde/tools/poc/commands/produce"
)

func Main() error {
	commands := map[string]func() error{
		"produce": produce.Run,
	}

	if len(os.Args) < 2 {
		return fmt.Errorf("subcommands: %s", strings.Join(slices.Collect(maps.Keys(commands)), ", "))
	}

	cmd := os.Args[1]
	command, ok := commands[cmd]
	if !ok {
		return fmt.Errorf("available subcommands: %s", strings.Join(slices.Collect(maps.Keys(commands)), ", "))
	}

	os.Args = os.Args[1:]
	err := command()
	if err != nil {
		return fmt.Errorf("%s: %w", cmd, err)
	}

	return nil
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
