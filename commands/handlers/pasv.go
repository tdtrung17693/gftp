package handlers

import (
	"fmt"
	"gftp/commands"
	"gftp/server"
)

type pasvCmdHandler struct {
}

func (h pasvCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	resChan := make(chan interface{})
	request := server.DtpOpenConnRequest{
		TransferType: "passive",
		ResponseChan: resChan,
	}
	var port int
	ctx.Dtp.ListenerChan <- request
	select {
	case res := <-resChan:
		switch v := res.(type) {
		case *server.DtpResponse:
			ctx.ConnDtpChan = v.MsgBus
			port = v.Port
		case error:
			return &server.Response{
				Code:    server.ReplyCannotOpenDataConnection,
				Message: "Cannot open data connection",
			}
		default:
			fmt.Println(v)
		}
	}
	p1 := port >> 8
	p2 := port & 0xFF
	fmt.Printf("DTP Port: %d\n", port)
	return &server.Response{
		Code:    server.ReplyPassiveMode,
		Message: fmt.Sprintf("127,0,0,1,%d,%d", p1, p2),
	}
}
