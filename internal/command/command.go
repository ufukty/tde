package command

import (
	"flag"
	"fmt"
	"os"
)

type Command interface {
	BindFlagsToArgs(*flag.FlagSet)    // Registers arguments to flags and assigns return values to implementing struct
	ManualParser(*flag.FlagSet) error // Tailing arguments are not processed with BindFlagsToArgs
	Run()                             // Called when user runs the program with this command
}

var (
	commands = map[string]Command{}
	flagSets = map[string]*flag.FlagSet{}
)

func RegisterCommand(commandName string, cmd Command) {
	flagSet := flag.NewFlagSet(commandName, flag.ExitOnError)
	flagSet.Usage = func() {
		fmt.Printf("Run help: \"tde help %s\"\n", commandName)
	}

	commands[commandName] = cmd
	flagSets[commandName] = flagSet
}

func terminate(msg any) {
	fmt.Println(msg)
	os.Exit(1)
}

func Route() {
	if len(os.Args) < 2 {
		terminate("Expected command. Run \"tde help\"")
	}

	for name, cmd := range commands {
		cmd.BindFlagsToArgs(flagSets[name])
	}

	cmdToRun := os.Args[1]
	if cmd, ok := commands[cmdToRun]; ok {
		if err := cmd.ManualParser(flagSets[cmdToRun]); err != nil {
			terminate(err)
		}
		cmd.Run()
		os.Exit(0)
	} else {
		terminate("Unrecognized command. Run \"tde help\"")
	}

}
