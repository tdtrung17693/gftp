package handlers

import (
	"fmt"
	"gftp/commands"
	"gftp/server"
)

type typeCmdHandler struct {
}

func (h typeCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	fmt.Printf("TYPE command: %q", cmd.Params)
	return &server.Response{
		Code:    server.ReplySystemStatus,
		Message: "UNIX",
	}
}
