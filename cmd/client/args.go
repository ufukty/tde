package main

import (
	"flag"
	"fmt"
	"os"
)

func CommandHelp(produce *flag.FlagSet) {
	if len(os.Args) < 3 {
		PrintHelp()
		return
	}

	switch os.Args[2] {
	case "produce":
		produce.Usage()
		return
	case "login":

		return
	case "logout":

		return
	case "credits":

		return
	default:
		fmt.Println("Unrecognized command for help. Run \"tde help\"")
		os.Exit(1)
	}
}

func PrintHelp() {
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

// func produce.Usage(produce *flag.FlagSet) {

// }

func SetupProduce() *flag.FlagSet {
	produce := flag.NewFlagSet("produce", flag.ExitOnError)
	produce.Usage = func() {
		fmt.Print(`Always run from directory contains the package.

Usage:

	tde produce [ -include-module ] TestName

Arguments:

`)
		produce.PrintDefaults()
	}
	produceCommandFlags.IncludeModule = *produce.Bool(
		"include-module", false, "(Optional) Enable full module upload. It is necessary if the package contains imports from rest of the module.")
	produceCommandFlags.RunnerAddress = *produce.String(
		"runner-address", "", "(Optional) Not needed when tde's runners wanted to use by user for compiling and running candidates. Needed for using custom runner. IP address and port in format: 127.0.0.1.26900")
	produceCommandFlags.Timeout = *produce.Int(
		"timeout", 10, "Period of time to wait before terminate.")
	return produce
}

func CommandProduce(produce *flag.FlagSet) {
	err := produce.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	produce := SetupProduce()

	// config.Population = *flag.Int(
	// 	"population", 1000, "Number of candidates generated at start and tested at each iteration at each generation")
	// config.Generation = *flag.Int(
	// 	"generation", 10, "Number of generations which the evolution will be stopped.")
	// config.SizeLimit = *flag.Int(
	// 	"size-limit", 1000, "Character size limit for any candidate.")

	if len(os.Args) < 2 {
		fmt.Println("Expected command. Run \"tde help\"")
		os.Exit(1)
	}

	switch command := os.Args[1]; command {
	case "help":
		CommandHelp(produce)
		os.Exit(0)

	case "produce":
		CommandProduce(produce)

	default:
		fmt.Println("Unrecognized command. Run \"tde help\"")
		os.Exit(1)
	}

}
