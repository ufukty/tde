package main

import (
	"tde/cmd/client-poc/internal/produce"
	"tde/internal/command"
)

func main() {
	command.RegisterCommand("", &produce.Command{})
}
