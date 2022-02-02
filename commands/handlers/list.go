package handlers

import (
	"gftp/commands"
	"gftp/dtp"
	"gftp/server"
	"path"
)

type listCmdHandler struct {
}

func (h listCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	currentDtpConn := ctx.DtpConn
	realPath := path.Join(ctx.ServerRoot, ctx.Pwd)

	err := currentDtpConn.SendMessage(dtp.DtpListRequest{
		Path: realPath,
	})

	if err != nil {
		return &server.Response{
			Code:    server.ReplyRequestedLocalProcesingError,
			Message: "Error in processing",
		}
	}

	return &server.Response{
		Code:    server.ReplyRequestedFileOk,
		Message: "Ok",
	}
}
