package main

import (
	"tde/internal/command"
	"tde/tools/poc/internal/produce"
)

func main() {
	command.RegisterCommand("produce", &produce.Command{})

	command.Route()
}
