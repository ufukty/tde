package main

import (
	"flag"
	"fmt"
)

type Config struct {
	File             string
	TestFile         string
	TestFunctionName string
}

var config = Config{}

func init() {
	flag.Usage = func() {
		fmt.Print("GoGP is a program to produce human-competitive, developer-readable Go code from a user provided test function which is common to have when TDD practices are followed.\n\n")
		fmt.Print("Usage:\n\n")
		fmt.Print("\tgogp -file=./path/to/file.go -test-file=./path/to/file_test.go -test-function=TestKnappsack\n\n")
		fmt.Print("Arguments:\n\n")
		flag.PrintDefaults()
	}

	config.File = *flag.String(
		"file", "path/to/file.go", "(Required) Relative path to the file that contains the function body you want to evolve.")
	config.TestFile = *flag.String(
		"test-file", "path/to/file_test.go", "(Required) Relative path to the test file that will be used as fitness measurement.")
	config.TestFunctionName = *flag.String(
		"test-function", "TestKnappsack", "(Required) Name of the test function that will be used as fitness measurement")

	flag.Parse()
}
