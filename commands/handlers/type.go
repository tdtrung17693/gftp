package handlers

import (
	"gftp/commands"
	"gftp/dtp"
	"gftp/server"
)

type typeCmdHandler struct {
}

func (h typeCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	// Currently not support format parameter
	transferType, err := dtp.TransferTypeFromParam(cmd.Params[0])
	if err != nil {
		return &server.Response{
			Code:    server.ReplyArgumentSyntaxError,
			Message: err.Error(),
		}
	}

	ctx.DtpTransferType = transferType

	return &server.Response{
		Code:    server.ReplyCommandOk,
		Message: "Transfer type has been set up",
	}
}
