package main

import (
	"fmt"
	"gftp/commands"
	"gftp/commands/handlers"
	"gftp/server"
	"log"
	"net"
)

type Server struct {
	port int
	host string
}

func NewFTPServer(port int, host string) *Server {
	c := &Server{
		port: port,
		host: host,
	}

	return c
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))

	if err != nil {
		log.Fatal(err)
	}

	dtp := server.NewDTP()

	fmt.Printf("Listening on %d...\n", s.port)
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn, dtp)
	}
}

func handleConn(c net.Conn, dtp *server.Dtp) {
	log.Printf("[%s] Client connected.\n", c.RemoteAddr())
	var resolver = handlers.CommandResolver{}
	var cmdProcessor = commands.NewCommandProcessor(resolver)
	handler := server.NewConnHandler(c, dtp, cmdProcessor)
	handler.Handle()
}
