package commands

import "gftp/server"

type CommandHandler interface {
	Handle(cmd *Command, ctx *server.ConnContext) *server.Response
}

type HandlerResolver interface {
	Resolve(cmd *Command) CommandHandler
}
