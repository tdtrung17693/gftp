package handlers

import (
	"gftp/commands"
	"gftp/server"
)

type systCmdHandler struct {
}

func (h systCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	return &server.Response{
		Code:    server.ReplySystemStatus,
		Message: "UNIX",
	}
}
