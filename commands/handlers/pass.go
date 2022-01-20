package handlers

import (
	"gftp/commands"
	"gftp/server"
)

type passCmdHandler struct {
}

func (h passCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	return &server.Response{
		Code:    server.ReplyUserLoggedIn,
		Message: "Ready",
	}
}
