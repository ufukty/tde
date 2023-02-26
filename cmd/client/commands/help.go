package commands

import (
	"flag"
	"fmt"
	"os"
)

type Help struct {
	fs *flag.FlagSet

	IncludeModule bool
	RunnerAddress string
	Timeout       int
}

func NewHelp() *Help {
	help := Help{
		fs: flag.NewFlagSet("help", flag.ExitOnError),
	}
	return &help
}

func (h *Help) Setup() {
}

func (h *Help) PrintGenericHelp() {
	fmt.Print(`Use tde to produce implementation that complies test.

Usage:

	tde <command> [arguments]

Commands:

	login		Required on first run. Asks username and pass.
	logout		Deletes session on this device.
	credits 	Prints info about remaining credits for account.
	
	produce 	Always run from directory contains the package.

Use "tde help <command>" for more information about a command.
`)
}

func (h *Help) Parse() int {
	return 0
}

func (h *Help) Run() int {
	if len(os.Args) < 3 { // means running as "tde help"
		h.PrintGenericHelp()
		return 0
	}

	cmdToRun := os.Args[2] // command part in "tde help <command>"

	if cmdToRun == "help" {
		fmt.Println("Help command helps you. Run \"tde help\" to learn more, instead \"tde help help\"")
		return 1 // avoid infinite recursion
	}

	if cmd, ok := Commands[cmdToRun]; ok {
		cmd.PrintGenericHelp()
		return 0
	} else {
		fmt.Println("Unrecognized command for help. Run \"tde help\"")
		return 1
	}
}
