package server

type ConnContext struct {
	State    string
	ResChan  chan Response
	Dtp      *Dtp
	DtpConn  *DtpConn
	Pwd      string
	UserRoot string
}

func NewConnContext(dtp *Dtp) ConnContext {
	connContext := ConnContext{
		State:   "",
		ResChan: make(chan Response),
		Dtp:     dtp,
		DtpConn: nil,
	}

	return connContext
}
