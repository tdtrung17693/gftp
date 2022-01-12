package commands

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Command interface {
	Handle() Response
}

type CommandHandler struct {
	conn    net.Conn
	scanner bufio.Scanner
}

type Response struct {
	code    int
	message string
}

type Request struct {
	command    string
	parameters []string
}

func ParseRequest(requestLine string) Request {
	components := strings.Split(requestLine, " ")
        fmt.Print(components[0])
        return Request{}

}

func NewCommandHandler(conn net.Conn) *CommandHandler {
	s := *bufio.NewScanner(conn)
	// s.Split(bufio.ScanWords)

	c := &CommandHandler{
		conn:    conn,
		scanner: s,
	}

	return c
}

func (handler *CommandHandler) getCurrentCommand() *Command {
	var c interface{}
	msg := handler.scanner.Text()
	switch msg {
	case "USER":
		c = &userCommand{}
	case "PASS":
		c = &passCommand{}
	case "FEAT":
		c = &featCommand{}
	default:
		c = &unknownCommand{}
	}

	return c.(*Command)
}

func (handler *CommandHandler) Handle() bool {
	if handler.scanner.Scan() {
		return false
	}

	// cmd := handler.getCurrentCommand()

	return true
}
