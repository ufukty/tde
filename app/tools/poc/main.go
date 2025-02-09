package main

import (
	"fmt"
	"os"
	"strings"
	"tde/tools/poc/commands/produce"

	"golang.org/x/exp/maps"
)

func Main() error {
	commands := map[string]func() error{
		"produce": produce.Run,
	}

	if len(os.Args) < 2 {
		return fmt.Errorf("subcommands: %s", strings.Join(maps.Keys(commands), ", "))
	}

	cmd := os.Args[1]
	command, ok := commands[cmd]
	if !ok {
		return fmt.Errorf("available subcommands: %s", strings.Join(maps.Keys(commands), ", "))
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
