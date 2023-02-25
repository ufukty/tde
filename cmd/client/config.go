package main

type ProduceCommandFlags struct {
	IncludeModule bool
	RunnerAddress string

	// File             string
	// TestFile         string
	TestFunctionName string

	// Population       int
	// Generation       int
	// SizeLimit        int

	Timeout int
}

var produceCommandFlags = ProduceCommandFlags{}
