package commands

import "gftp/server"

type CommandProcessor struct {
	handlerResolver HandlerResolver
}

func NewCommandProcessor(handlerResolver HandlerResolver) server.CommandProcessor {
	return CommandProcessor{
		handlerResolver: handlerResolver,
	}
}

func (c CommandProcessor) Handle(requestLine string, ctx *server.ConnContext) *server.Response {
	cmd := parseCommand(requestLine)
	handler := c.handlerResolver.Resolve(cmd)

	return handler.Handle(cmd, ctx)
}
