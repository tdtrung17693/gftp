package handlers

import (
	"gftp/commands"
	"gftp/server"
	"log"
)

type listCmdHandler struct {
}

func (h listCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	currentDtpConn := ctx.DtpConn

	currentDtpConn.MsgChan <- server.DtpListRequest{
		Path: ctx.Pwd,
	}

	select {
	case value := <-currentDtpConn.MsgChan:
		log.Println(value)
		return &server.Response{
			Code:    server.ReplyRequestedFileOk,
			Message: "Ok",
		}
	case err := <-currentDtpConn.ErrChan:
		log.Println(err)
		return &server.Response{
			Code:    server.ReplyRequestedLocalProcesingError,
			Message: "Error in processing",
		}
	}

}
