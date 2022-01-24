package server

const (
	StateReady = iota
	StateWaitForPassword
	StateAuthenticated
)

type ConnContext struct {
	State      int
	ResChan    chan Response
	Dtp        *Dtp
	DtpConn    *DtpConn
	Pwd        string
	Username   string
	UserRoot   string
	ServerRoot string
}

func NewConnContext(dtp *Dtp, serverRoot string) ConnContext {
	connContext := ConnContext{
		State:      StateReady,
		ResChan:    make(chan Response),
		Dtp:        dtp,
		DtpConn:    nil,
		ServerRoot: serverRoot,
	}

	return connContext
}
