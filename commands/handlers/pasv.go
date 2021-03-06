package handlers

import (
	"gftp/commands"
	"gftp/dtp"
	"gftp/server"
)

type pasvCmdHandler struct {
}

func (h pasvCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	dtpConn, err := ctx.Dtp.OpenNewConn(dtp.TransferModePassive)

	if err != nil {
		return &server.Response{
			Code:    server.ReplyCannotOpenDataConnection,
			Message: "Cannot open data connection",
		}
	}

	ctx.DtpConn = dtpConn

	return &server.Response{
		Code:    server.ReplyPassiveMode,
		Message: dtpConn.GetPASVMsg(),
	}
}
