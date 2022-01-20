package handlers

import (
	"gftp/commands"
	"gftp/server"
)

type unknownCmdHandler struct {
}

func (h unknownCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	return &server.Response{
		Code:    server.ReplyCommandNotImplemented,
		Message: "Command not implemented",
	}
}
