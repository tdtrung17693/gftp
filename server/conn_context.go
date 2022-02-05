package server

import (
	"fmt"
	"gftp/dtp"
	"log"
	"net"
	"path"
)

const (
	StateReady = iota
	StateWaitForPassword
	StateAuthenticated
)

type ConnContext struct {
	State           int
	tcpConn         net.Conn
	Dtp             *dtp.Dtp
	DtpConn         *dtp.DtpConn
	DtpTransferType dtp.TransferType
	Pwd             string
	Username        string
	UserRoot        string
	ServerRoot      string
}

func (ctx *ConnContext) GetServerPath(clientPath string) string {
	return path.Join(ctx.ServerRoot, ctx.UserRoot, clientPath)
}

func (ctx *ConnContext) GetUserServerPath() string {
	return path.Join(ctx.ServerRoot, ctx.UserRoot, ctx.Pwd)
}

func NewConnContext(tcpConn net.Conn, dtp *dtp.Dtp, serverRoot string) ConnContext {
	connContext := ConnContext{
		State:      StateReady,
		Dtp:        dtp,
		DtpConn:    nil,
		ServerRoot: serverRoot,
		tcpConn:    tcpConn,
	}

	return connContext
}

func (c *ConnContext) WriteResponse(res *Response) {
	fmt.Fprintf(c.tcpConn, "%d %s\n", res.Code, res.Message)
}

func (c *ConnContext) logf(p string, args ...interface{}) {
	log.Printf(
		fmt.Sprintf("[%s] %s", c.tcpConn.RemoteAddr(), p),
		args...)
}
