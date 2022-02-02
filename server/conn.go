package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type ConnHandler struct {
	Conn             net.Conn
	commandProcessor CommandProcessor
	Scanner          *bufio.Scanner
	Context          ConnContext
}

func NewConnHandler(conn net.Conn, connContext ConnContext, commandProcessor CommandProcessor) *ConnHandler {
	s := bufio.NewScanner(conn)

	c := &ConnHandler{
		Conn:             conn,
		Scanner:          s,
		Context:          connContext,
		commandProcessor: commandProcessor,
	}

	return c
}

func (h *ConnHandler) logf(p string, args ...interface{}) {
	log.Printf(
		fmt.Sprintf("[interpreter] [%s] %s", h.Conn.RemoteAddr(), p),
		args...)
}

func (h *ConnHandler) Handle() bool {
	scanner := h.Scanner

	h.logf("Handling connection...")
	fmt.Fprint(h.Conn, "200 - READY\n")
	for {
		if !scanner.Scan() {
			err := scanner.Err()

			if err != nil {
				fmt.Println("Error")
				fmt.Print(scanner.Err())
			}

			break
		}

		res := h.commandProcessor.Handle(scanner.Text(), &h.Context)

		fmt.Fprintf(h.Conn, "%d %s\n", res.Code, res.Message)
	}

	return true
}
