package handlers

import (
	"gftp/commands"
	"gftp/server"
)

type pwdCmdHandler struct {
}

func (h pwdCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	return &server.Response{
		Code:    server.ReplyPathNameCreated,
		Message: "\"/home/anonymous\" is cwd.",
	}
}
