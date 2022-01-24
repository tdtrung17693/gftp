package handlers

import (
	"fmt"
	"gftp/commands"
	"gftp/server"
)

type pwdCmdHandler struct {
}

func (h pwdCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	return &server.Response{
		Code:    server.ReplyPathNameCreated,
		Message: fmt.Sprintf("\"%s\" is cwd.", ctx.Pwd),
	}
}
