package handlers

import (
	"gftp/commands"
	"gftp/dtp"
	"gftp/server"
	"path"
	"strings"
)

type storCmdHandler struct {
}

func (h storCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	// Mark the transfer starting point
	ctx.WriteResponse(&server.Response{
		Code:    server.ReplyFileStatusOk,
		Message: "",
	})

	err := ctx.DtpConn.SendMessage(dtp.DtpTransferRequest{
		FilePath:       path.Join(ctx.ServerRoot, ctx.UserRoot, ctx.Pwd, strings.Join(cmd.Params, " ")),
		TransferAction: dtp.TransferActionStore,
		TransferType:   ctx.DtpTransferType,
	})

	if err != nil {
		return &server.Response{
			Code:    server.ReplyRequestedFileUnavailableOrBusy,
			Message: err.Error(),
		}
	}

	return &server.Response{
		Code:    server.ReplyClosingDataConnection,
		Message: "OK",
	}
}
