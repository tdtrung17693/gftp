package handlers

import (
	"fmt"
	"gftp/commands"
	"gftp/server"
)

type cwdCmdHandler struct {
}

func (h cwdCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	ctx.Pwd = cmd.Params[0]

	return &server.Response{
		Code:    server.ReplyRequestedFileOk,
		Message: fmt.Sprintf("\"%s\" is cwd.", ctx.Pwd),
	}
}
