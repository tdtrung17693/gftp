package dtp

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type DtpConn struct {
	msgChan     chan interface{}
	errChan     chan error
	tcpConn     net.Conn
	tcpListener net.Listener
	port        int
}

type dtpSuccessResponse struct{}

func newDtpConn() (*DtpConn, error) {
	listener, err := net.Listen("tcp", ":0")

	if err != nil {
		return nil, err
	}

	addr := listener.Addr().String()
	lstColon := strings.LastIndex(addr, ":")
	portStr := addr[lstColon+1:]
	port, err := strconv.Atoi(portStr)

	if err != nil {
		return nil, err
	}

	dtpConnChan := make(chan interface{})
	errConnChan := make(chan error)

	dtpConn := DtpConn{
		port:        port,
		msgChan:     dtpConnChan,
		errChan:     errConnChan,
		tcpListener: listener,
	}
	return &dtpConn, nil
}

func (dtpConn *DtpConn) SendMessage(dtpCmd interface{}) error {
	log.Println(dtpConn.msgChan)
	dtpConn.msgChan <- dtpCmd

	select {
	case <-dtpConn.msgChan:
		return nil
	case err := <-dtpConn.errChan:
		return err
	}
}

func (dtpConn *DtpConn) serve() {
	// only 1 client for 1 server dtp,
	// so just need 1 goroutine to handle
	conn, err := dtpConn.tcpListener.Accept()

	if err != nil {
		dtpConn.errChan <- err
		return
	}

	dtpConn.tcpConn = conn

	dtpConn.logf("Client connected.")
	dtpConn.logf("Waiting for command...")
	cmd := <-dtpConn.msgChan

	switch a := cmd.(type) {
	case DtpListRequest:
		dtpConn.logf("LIST command received.\n")
		cmd := exec.Command("ls", "-ll")
		cmd.Dir = a.Path
		res, err := cmd.Output()

		if err != nil {
			dtpConn.logf("Error: %s", err)
			dtpConn.sendError(err)
			return
		}

		err = dtpConn.write(res)

		if err != nil {
			dtpConn.logf("Error: %s", err)
			dtpConn.sendError(err)
			return
		}

		conn.Close()
		dtpConn.msgChan <- dtpSuccessResponse{}

		dtpConn.logf("Finish LIST")
	case DtpTransferRequest:
		dtpConn.logf("RETR command received")
		file, err := os.ReadFile(a.FilePath)
		if err != nil {
			dtpConn.logf("Error: %s", err)
		}

		_, err = fmt.Fprint(conn, string(file))
		if err != nil {
			dtpConn.logf("Error: %s", err)
			dtpConn.sendError(err)
		}

		conn.Close()
		dtpConn.msgChan <- dtpSuccessResponse{}
		dtpConn.logf("Finish RETR")
	}
}

func (dtpConn *DtpConn) sendError(e error) {
	dtpConn.errChan <- e
}

func (dtpConn *DtpConn) GetPASVMsg() string {
	p1 := dtpConn.port >> 8
	p2 := dtpConn.port & 0xFF

	return fmt.Sprintf("127,0,0,1,%d,%d", p1, p2)
}

func (dtpConn *DtpConn) GetEPSVMsg() string {
	return fmt.Sprintf(" Entering Extended Passive Mode (|||%d|)", dtpConn.port)
}

func (dtpConn *DtpConn) write(d []byte) error {
	_, err := fmt.Fprint(dtpConn.tcpConn, string(d))

	return err
}

func (dtpConn *DtpConn) logf(p string, args ...interface{}) {
	log.Printf(
		fmt.Sprintf("[DTP] [%s] %s", dtpConn.tcpConn.RemoteAddr(), p),
		args...)
}
