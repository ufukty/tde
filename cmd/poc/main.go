package main

import (
	"tde/cmd/poc/internal/produce"
	"tde/cmd/poc/internal/prepare"
	"tde/internal/command"
)

func main() {
	command.RegisterCommand("produce", &produce.Command{})
	command.RegisterCommand("prepare", &prepare.Command{})

	command.Route()
}
