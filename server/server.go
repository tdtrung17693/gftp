package server

import (
	"fmt"
	"gftp/dtp"
	"log"
	"net"
)

type Server struct {
	port             int
	host             string
	dataRoot         string
	commandProcessor CommandProcessor
}

func NewFTPServer(port int, host string, dataRoot string, commandProcessor CommandProcessor) *Server {
	c := &Server{
		port:             port,
		host:             host,
		dataRoot:         dataRoot,
		commandProcessor: commandProcessor,
	}

	return c
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))

	if err != nil {
		log.Fatal(err)
	}

	newDTP := dtp.NewDTP()

	fmt.Printf("Listening on %d...\n", s.port)
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		go s.handleConn(conn, newDTP, s.commandProcessor)
	}
}

func (s *Server) handleConn(c net.Conn, dtp *dtp.Dtp, cmdProcessor CommandProcessor) {
	context := NewConnContext(dtp, s.dataRoot)
	handler := NewConnHandler(c, context, cmdProcessor)
	handler.Handle()
}
