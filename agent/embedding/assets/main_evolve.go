//go:build tde
// +build tde

package main

// usage from commandline, after this file places into user's project root
//   go build -tags tdgp

import (
	"GoGP/testing/evolution"
	"flag"

	"fmt"
	"os"
)

// func checkFolder(path string) {
// 	dirEntries, err := os.ReadDir(path)
// 	if err != nil {
// 		errors.Wrap(err, fmt.Sprintf("Could not list the directory '%s'", path))
// 	}
// 	foldersToCheck := []string{}
// 	filesToCheck := []string{}
// 	for _, dirEntry := range dirEntries {
// 		path := filepath.Join(path, dirEntry.Name())
// 		if dirEntry.IsDir() {
// 			foldersToCheck = append(foldersToCheck, path)
// 		} else {
// 			filesToCheck = append(filesToCheck, path)
// 		}
// 	}

// 	for _, filePath := range filesToCheck {
// 		fmt.Println(filePath)
// 	}

// 	for _, folderPath := range foldersToCheck {
// 		fmt.Println(folderPath)
// 	}
// }

type Arguments struct {
	FunctionNameToRun string
}

func ParseArguments() *Arguments {
	arguments := Arguments{}

	flag.Usage = func() {
		fmt.Print("GoGP is a program to produce human-competitive, developer-readable Go code from a user provided test function which is common to have when TDD practices are followed.\n\n")
		fmt.Print("Usage:\n\n")
		fmt.Print("\tgogp -file=./path/to/file.go -test-file=./path/to/file_test.go -test-function=TestKnappsack\n\n")
		fmt.Print("Arguments:\n\n")
		flag.PrintDefaults()
	}

	arguments.FunctionNameToRun = *flag.String(
		"run", "ExamplePackage/TDE_Example", "(Required) Name of the test function that will be used as fitness measurement")

	flag.Parse()

	return &arguments
}

func init() {

	// targetFunc := os.Args[1]
	var evolveFunction func(*evolution.E)
	e := evolution.NewE(candidateRefs)
	evolveFunction(e)
	e.PrintStats()

	// search all files that ends with xxx_evolve.go for a function starts with "Evolve<Target>"

	os.Exit(0)
}

func main() {
	fmt.Println("Goodbye world")
}
