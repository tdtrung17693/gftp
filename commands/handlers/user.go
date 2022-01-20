package handlers

import (
	"gftp/commands"
	"gftp/server"
)

type userCmdHandler struct {
}

func (h userCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	return &server.Response{
		Code:    server.ReplyNeedPassword,
		Message: "Anonymous mode. Provide email to continue",
	}
}
