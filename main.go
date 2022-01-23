package main

import (
	"gftp/commands"
	"gftp/commands/handlers"
	"gftp/server"
)

func main() {
	resolver := handlers.CommandResolver{}
	cmdProcessor := commands.NewCommandProcessor(resolver)
	server := server.NewFTPServer(2121, "localhost", cmdProcessor)
	server.Start()
}
