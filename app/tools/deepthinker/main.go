package main

/*
Usage:
	cd ./example
	gp Knappsack pop=1000 gen=10 bloat=1024

Main progress:
	Discovery -> starting and ending lines of
				 the blocks in the implementation
				 file and test file
	Initialization : Random code segments will be
	      		     created for empty bodies when
					 they are given as input.
	Evolve : Iterates the evolution by 1 generation.
*/

/*
Functions:
	- zips a module (or package specified by user) and uploads to server with access token
	- orders from evolver server to *start* evolution
	- orders from evolver server to *continue* evolution
	- tracks versions for working directory, when user changes
*/

import (
	"fmt"
	"os"
	"strings"
	"tde/tools/deepthinker/internal/help"
	"tde/tools/deepthinker/internal/listtest"
	"tde/tools/deepthinker/internal/produce"
	"tde/tools/deepthinker/internal/upload"

	"golang.org/x/exp/maps"
)

func Main() error {
	commands := map[string]func() error{
		"help":      help.Run,
		"upload":    upload.Run,
		"produce":   produce.Run,
		"list-test": listtest.Run,
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
