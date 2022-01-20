package server

import (
	"bufio"
	"fmt"
	"net"
)

type ConnHandler struct {
	Conn             net.Conn
	commandProcessor CommandProcessor
	Scanner          *bufio.Scanner
	Context          ConnContext
}

type ConnContext struct {
	State       string
	ResChan     chan Response
	Dtp         *Dtp
	ConnDtpChan chan interface{}
	Pwd         string
	UserRoot    string
}

func NewConnHandler(conn net.Conn, dtp *Dtp, commandProcessor CommandProcessor) *ConnHandler {
	s := bufio.NewScanner(conn)

	context := ConnContext{
		State:       "",
		ResChan:     make(chan Response),
		Dtp:         dtp,
		ConnDtpChan: nil,
	}

	c := &ConnHandler{
		Conn:             conn,
		Scanner:          s,
		Context:          context,
		commandProcessor: commandProcessor,
	}

	return c
}

func (h *ConnHandler) Handle() bool {
	scanner := h.Scanner

	fmt.Println("Handling connection...")
	fmt.Fprint(h.Conn, "200 - READY\n")
	for {
		if !scanner.Scan() {
			fmt.Println("Error")
			fmt.Print(scanner.Err())
			break
		}

		res := h.commandProcessor.Handle(scanner.Text(), &h.Context)

		h.Conn.Write([]byte(fmt.Sprintf("%d: %s\n", res.Code, res.Message)))

	}

	return true
}
