package handlers

import (
	"gftp/commands"
	"gftp/dtp"
	"gftp/server"
	"path"
)

type retrCmdHandler struct {
}

func (h retrCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	err := ctx.DtpConn.SendMessage(dtp.DtpTransferRequest{
		FilePath:       path.Join(ctx.ServerRoot, ctx.Pwd, cmd.Params[0]),
		TransferAction: dtp.TransferActionRetrieve,
		TransferType:   ctx.DtpTransferType,
	})

	if err != nil {
		return &server.Response{
			Code:    server.ReplyRequestedFileUnavailableOrBusy,
			Message: err.Error(),
		}
	}

	return &server.Response{
		Code:    server.ReplyFileStatusOk,
		Message: "OK",
	}
}
