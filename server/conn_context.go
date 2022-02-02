package server

import "gftp/dtp"

const (
	StateReady = iota
	StateWaitForPassword
	StateAuthenticated
)

type ConnContext struct {
	State           int
	ResChan         chan Response
	Dtp             *dtp.Dtp
	DtpConn         *dtp.DtpConn
	DtpTransferType dtp.TransferType
	Pwd             string
	Username        string
	UserRoot        string
	ServerRoot      string
}

func NewConnContext(dtp *dtp.Dtp, serverRoot string) ConnContext {
	connContext := ConnContext{
		State:      StateReady,
		ResChan:    make(chan Response),
		Dtp:        dtp,
		DtpConn:    nil,
		ServerRoot: serverRoot,
	}

	return connContext
}
