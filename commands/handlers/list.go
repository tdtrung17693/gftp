package handlers

import (
	"gftp/commands"
	"gftp/server"
)

type listCmdHandler struct {
}

func (h listCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	currentDtpChan := ctx.ConnDtpChan

	currentDtpChan <- server.DtpListRequest{
		Path: ctx.Pwd,
	}

	return &server.Response{
		Code:    server.ReplyRequestedFileOk,
		Message: "Ok",
	}
}
