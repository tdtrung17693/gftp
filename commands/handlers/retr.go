package handlers

import (
	"gftp/commands"
	"gftp/server"
	"log"
	"path"
)

type retrCmdHandler struct {
}

func (h retrCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	dtpRequest := server.DtpTransferRequest{
		FilePath:     path.Join(ctx.ServerRoot, ctx.Pwd, cmd.Params[0]),
		TransferMode: server.TransferModeRetrieve,
		TransferType: server.TransferTypeASCII,
	}
	log.Print("[RETR] Sending file request")
	ctx.DtpConn.MsgChan <- dtpRequest

	select {
	case <-ctx.DtpConn.MsgChan:
		return &server.Response{
			Code:    server.ReplyFileStatusOk,
			Message: "OK",
		}
	case err := <-ctx.DtpConn.ErrChan:
		return &server.Response{
			Code:    server.ReplyRequestedFileUnavailableOrBusy,
			Message: err.Error(),
		}
	}
}
