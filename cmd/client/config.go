package main

type Config struct {
	File             string
	TestFile         string
	TestFunctionName string
	Population       int
	Generation       int
	SizeLimit        int
}

var config = Config{}
