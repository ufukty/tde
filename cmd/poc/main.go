package main

import (
	"tde/cmd/poc/internal/prepare"
	"tde/internal/command"
)

func main() {
	command.RegisterCommand("prepare", &prepare.Command{})

	command.Route()
}
