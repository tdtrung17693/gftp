package handlers

import (
	"gftp/commands"
	"gftp/server"
)

type featCmdHandler struct {
}

func (h featCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	return &server.Response{
		Code:    server.ReplyCommandNotImplemented,
		Message: "Command not implemented",
	}
}
