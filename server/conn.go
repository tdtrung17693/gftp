package server

import (
	"bufio"
	"fmt"
	"net"
)

type ConnHandler struct {
	commandProcessor CommandProcessor
	Scanner          *bufio.Scanner
	Context          ConnContext
}

func NewConnHandler(conn net.Conn, connContext ConnContext, commandProcessor CommandProcessor) *ConnHandler {
	s := bufio.NewScanner(conn)

	c := &ConnHandler{
		Scanner:          s,
		Context:          connContext,
		commandProcessor: commandProcessor,
	}

	return c
}

func (h *ConnHandler) logf(p string, args ...interface{}) {
	h.Context.logf(
		"%s", args...)
}

func (h *ConnHandler) Handle() bool {
	scanner := h.Scanner

	h.logf("Handling connection...")
	h.Context.WriteResponse(&Response{Code: ReplyCommandOk, Message: "READY"})

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

		// This line just handles the final stage of the command
		// Some commands require several stages, hence its handler
		// could also send response to the client
		h.Context.WriteResponse(res)
	}

	return true
}
