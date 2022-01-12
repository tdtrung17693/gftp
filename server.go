package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	log.Printf("[%s] Client connected.\n", c.RemoteAddr())
	reader := bufio.NewScanner(c)
	fmt.Fprint(c, "200 - READY\n")
	var msg string
	for {
		if !reader.Scan() {
			continue
		}

		b := reader.Text()

		msg = string(b)
		fmt.Println(msg)

		if strings.HasPrefix(msg, "USER") {
			fmt.Fprint(c, "331 - Guest login ok, send your complete e-mail address as password.\n")
		} else if strings.HasPrefix(msg, "PASS") {
			fmt.Fprint(c, "230 - Greeting\n")
		} else if strings.HasPrefix(msg, "SYST") {
			fmt.Fprint(c, `
215 UNIX Type: L8
Remote system type is UNIX.
Using binary mode to transfer files.
			`)
		} else if strings.HasPrefix(msg, "PWD") {
			fmt.Fprint(c, "257 \"/usr/home/ixl\" is cwd.\n")
		} else if strings.HasPrefix(msg, "TYPE") {
			fmt.Fprint(c, "200 OK\n")
		} else if strings.HasPrefix(msg, "PASV") {
			fmt.Fprint(c, "227 Entering Passive Mode\n")
		} else if strings.HasPrefix(msg, "PORT") {
			fmt.Fprint(c, "200 OK\n")
		}

	}
}
