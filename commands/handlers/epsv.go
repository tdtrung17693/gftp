package handlers

import (
	"gftp/commands"
	"gftp/dtp"
	"gftp/server"
)

type epsvCmdHandler struct {
}

func (h epsvCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	dtpConn, err := ctx.Dtp.OpenNewConn(dtp.TransferModePassive)

	if err != nil {
		return &server.Response{
			Code:    server.ReplyCannotOpenDataConnection,
			Message: "Cannot open data connection",
		}
	}

	ctx.DtpConn = dtpConn

	return &server.Response{
		Code:    server.ReplyExtendedPassiveMode,
		Message: dtpConn.GetEPSVMsg(),
	}
}
