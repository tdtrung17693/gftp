package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Dtp struct {
	ListenerChan chan interface{}
	ErrorChan    chan error
}

type DtpConn struct {
	MsgChan chan interface{}
	ErrChan chan error
}

type DtpOpenConnRequest struct {
	TransferType string
	ResponseChan chan interface{}
}

type DtpResponse struct {
	Port   int
	MsgBus chan interface{}
	ErrBus chan error
}

type DtpListRequest struct {
	Path   string
	MsgBus chan interface{}
}

type DtpTransferRequest struct {
	FilePath     string
	TransferMode int
	TransferType int
}

const (
	TransferTypeASCII = iota
	TransferTypeEBCDIC
	TransferTypeImage
	TransferTypeLocal
)

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
		ErrBus: errConnChan,
	}

	return &dtpResponse, nil
}

func dtpListener(l net.Listener, dtpChan chan interface{}, errChan chan error) {
	// only 1 client for 1 server dtp,
	// so just need 1 goroutine to handle
	defer l.Close()
	conn, err := l.Accept()

	if err != nil {
		fmt.Println(err)
		errChan <- err
		return
	}

	fmt.Printf("[DTP] [%s] Client connected.\n", conn.RemoteAddr())
	cmd := <-dtpChan
	switch a := cmd.(type) {
	case DtpListRequest:
		log.Printf("[DTP] [%s] LIST command received.\n", conn.RemoteAddr())
		cmd := exec.Command("ls", "-ll")
		cmd.Dir = a.Path
		res, err := cmd.Output()
		if err != nil {
			log.Printf("[DTP] [%s] Error: %s", conn.RemoteAddr(), err)
			errChan <- err
			return
		}
		_, err = fmt.Fprint(conn, string(res))
		if err != nil {
			log.Printf("[DTP] [%s] Error: %s", conn.RemoteAddr(), err)
			errChan <- err
			return
		}

		conn.Close()
		dtpChan <- DtpResponse{}
		log.Printf("[DTP] [%s] Finish LIST\n", conn.RemoteAddr())
	case DtpTransferRequest:
		log.Printf("[DTP] [%s] RETR command received.\n", conn.RemoteAddr())
		file, err := os.ReadFile(a.FilePath)
		if err != nil {
			log.Printf("[DTP] [%s] Error: %s", conn.RemoteAddr(), err)
		}

		_, err = fmt.Fprint(conn, string(file))
		if err != nil {
			log.Printf("[DTP] [%s] Error: %s", conn.RemoteAddr(), err)
			errChan <- err
			return
		}

		conn.Close()
		dtpChan <- DtpResponse{}
		log.Printf("[DTP] [%s] Finish RETR\n", conn.RemoteAddr())
	}
}
