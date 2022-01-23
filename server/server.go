package server

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	port             int
	host             string
	commandProcessor CommandProcessor
}

func NewFTPServer(port int, host string, commandProcessor CommandProcessor) *Server {
	c := &Server{
		port:             port,
		host:             host,
		commandProcessor: commandProcessor,
	}

	return c
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))

	if err != nil {
		log.Fatal(err)
	}

	dtp := NewDTP()

	fmt.Printf("Listening on %d...\n", s.port)
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn, dtp, s.commandProcessor)
	}
}

func handleConn(c net.Conn, dtp *Dtp, cmdProcessor CommandProcessor) {
	log.Printf("[%s] Client connected.\n", c.RemoteAddr())
	handler := NewConnHandler(c, dtp, cmdProcessor)
	handler.Handle()
}
