package handlers

import (
	"gftp/commands"
	"gftp/server"
)

type typeCmdHandler struct {
}

func (h typeCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	return &server.Response{
		Code:    server.ReplySystemStatus,
		Message: "UNIX",
	}
}
