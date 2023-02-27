package commands

import (
	"tde/cmd/client/produce"
	"tde/cmd/client/upload"

	"fmt"
	"os"
)

type Command interface {
	Setup() // Sets flags, "Usage" function
	PrintGenericHelp()
	Parse() int
	Run() // Called when user runs the program with this command
}

// Doesn't include "help" command
var Commands = map[string]Command{
	"help":    NewHelp(),
	"produce": produce.NewProduce(),
	"upload":  upload.NewUpload(),
	// "login": ,
	// "logout": ,
	// "credits": ,
	// "recheck": ,
}

func exitIfFailed(statusCode int) {
	if statusCode != 0 {
		os.Exit(statusCode)
	}
}

func init() {

	if len(os.Args) < 2 {
		fmt.Println("Expected command. Run \"tde help\"")
		os.Exit(1)
	}

	for _, cmd := range Commands {
		cmd.Setup()
	}

	cmdToRun := os.Args[1]
	if cmd, ok := Commands[cmdToRun]; ok {
		exitIfFailed(cmd.Parse())
		cmd.Run()
		os.Exit(0)
	} else {
		fmt.Println("Unrecognized command. Run \"tde help\"")
		os.Exit(1)
	}
}
