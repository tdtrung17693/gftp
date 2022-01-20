package server

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
)

type Dtp struct {
	ListenerChan chan interface{}
	ErrorChan    chan error
}

type DtpOpenConnRequest struct {
	TransferType string
	ResponseChan chan interface{}
}

type DtpResponse struct {
	Port   int
	MsgBus chan interface{}
}

type DtpListRequest struct {
	Path   string
	MsgBus chan interface{}
}

type DtpTransferRequest struct {
	FileName     string
	TransferMode string
	TransferType int
}

const (
	TransferModeStore = iota
	TransferModeRetrieve
)

func NewDTP() *Dtp {
	dtpChan := make(chan interface{})
	errChan := make(chan error)

	go func() {
		for {
			select {
			case a := <-dtpChan:
				switch v := a.(type) {
				case DtpOpenConnRequest:
					res, err := startDTPConnection()

					if err != nil {
						v.ResponseChan <- err
					}

					v.ResponseChan <- res
				default:
					errChan <- fmt.Errorf("unknown DTP request (Received: %q)", a)
				}
			}
		}
	}()

	return &Dtp{
		ListenerChan: dtpChan,
		ErrorChan:    errChan,
	}
}

func startDTPConnection() (*DtpResponse, error) {
	listener, err := net.Listen("tcp4", ":0")

	if err != nil {
		return nil, err
	}

	addr := listener.Addr().String()
	parts := strings.Split(addr, ":")
	portStr := parts[1]
	port, err := strconv.Atoi(portStr)

	if err != nil {
		return nil, err
	}

	dtpConnChan := make(chan interface{})
	errConnChan := make(chan error)

	go dtpListener(listener, dtpConnChan, errConnChan)

	dtpResponse := DtpResponse{
		Port:   port,
		MsgBus: dtpConnChan,
	}

	return &dtpResponse, nil
}

func dtpListener(l net.Listener, dtpChan chan interface{}, errChan chan error) {
	// only 1 client for 1 server dtp,
	// so just need 1 goroutine to handle
	defer l.Close()
	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println(err)
			errChan <- err
			return
		}

		fmt.Printf("[DTP] [%s] Client connected.\n", conn.RemoteAddr())
		select {
		case cmd := <-dtpChan:
			switch a := cmd.(type) {
			case DtpListRequest:
				writer := bufio.NewWriter(conn)
				fmt.Printf("[DTP] [%s] LIST command received.\n", conn.RemoteAddr())
				res, err := exec.Command("ls", "-l").Output()
				if err != nil {
					fmt.Printf("[DTP] [%s] Error: %s", conn.RemoteAddr(), err)
					continue
				}
				fmt.Print(string(res))
				writer.Write(res)
				conn.Close()
			default:
				fmt.Print(a, conn)
			}
		}

	}
}
