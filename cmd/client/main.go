package main

/*
Usage:
	cd ./example
	gp Knappsack pop=1000 gen=10 bloat=1024

Main progress:
	Discovery -> starting and ending lines of
				 the blocks in the implementation
				 file and test file
	Initialization : Random code segments will be
	      		     created for empty bodies when
					 they are given as input.
	Evolve : Iterates the evolution by 1 generation.
*/

/*
Functions:
	- zips a module (or package specified by user) and uploads to server with access token
	- orders from evolver server to *start* evolution
	- orders from evolver server to *continue* evolution
	- tracks versions for working directory, when user changes
*/

import (
	"tde/cmd/client/internal/help"
	"tde/cmd/client/internal/list_test"
	"tde/cmd/client/internal/produce"
	"tde/cmd/client/internal/upload"

	"tde/internal/command"
)

func main() {
	command.RegisterCommand("help", &help.Command{})
	command.RegisterCommand("upload", &upload.Command{})
	command.RegisterCommand("produce", &produce.Command{})
	command.RegisterCommand("list-test", &list_test.Command{})

	command.Route()
}
