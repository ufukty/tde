package main

import (
	"tde/cmd/poc/internal/produce"
	"tde/internal/command"
)

func main() {
	command.RegisterCommand("", &produce.Command{})
}
