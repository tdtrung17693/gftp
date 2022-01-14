package commands

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type ConnHandler struct {
	conn    net.Conn
	scanner bufio.Scanner
}

type ConnContext struct {
	State   string
	ResChan chan Response
}

type Command struct {
	Name   string
	Params []string
}

type CommandHandler interface {
	Handle(cmd *Command, ctx *ConnContext) Response
}

type Response struct {
	code    int
	message string
}

func ParseCommand(requestLine string) *Command {
	components := strings.Split(requestLine, " ")

	var params []string
	if len(components) > 1 {
		params = components[1:]
	}

	return &Command{
		Name:   components[0],
		Params: params,
	}
}

func NewConnHandler(conn net.Conn) *ConnHandler {
	s := *bufio.NewScanner(conn)
	// s.Split(bufio.ScanWords)

	c := &ConnHandler{
		conn:    conn,
		scanner: s,
	}

	return c
}

func (handler *ConnHandler) Handle() bool {
	if handler.scanner.Scan() {
		return false
	}

	cmd := ParseCommand(handler.scanner.Text())

	fmt.Println(cmd)

	return true
}
