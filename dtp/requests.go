package dtp

import (
	"errors"
	"io"
	"os"
	"os/exec"
)

type DtpListRequest struct {
	Path string
}

type DtpTransferRequest struct {
	FilePath       string
	TransferAction TransferAction
	TransferType   TransferType
}

type DtpMkdRequest struct {
	DirPath string
}

type dtpRequestHandler interface {
	handle(dtpConn *DtpConn, request interface{})
}

type dtpTransferRequestHandler struct{}

func (h *dtpTransferRequestHandler) handle(dtpConn *DtpConn, request interface{}) {
	transferRequest := request.(DtpTransferRequest)
	dtpConn.logf("RETR command received")
	switch transferRequest.TransferAction {
	case TransferActionRetrieve:
		err := transferActionRetrieve(dtpConn, transferRequest.FilePath)
		if err != nil {
			dtpConn.logf("error: %q", err)
			dtpConn.sendError(err)
		}
		dtpConn.logf("Finish RETR")
	case TransferActionStore:
		err := transferActionStore(dtpConn, transferRequest.FilePath)
		if err != nil {
			dtpConn.logf("error: %q", err)
			dtpConn.sendError(err)
		}
		dtpConn.logf("Finish STOR")
	}

	dtpConn.msgChan <- dtpSuccessResponse{}
}

func transferActionRetrieve(dtpConn *DtpConn, filePath string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = dtpConn.write(file)
	if err != nil {
		return err
	}

	return nil
}

func transferActionStore(dtpConn *DtpConn, filePath string) error {
	file, err := os.Create(filePath)
	defer func() {
		err := file.Close()
		if err != nil {
			dtpConn.logf("error closing file %s - err: %q", filePath, err)
		}
	}()

	if err != nil {
		return err
	}

	_, err = io.Copy(file, dtpConn.tcpConn)

	if err != nil {
		return err
	}

	if err := file.Sync(); err != nil {
		return err
	}

	return nil
}

type dtpListRequestHandler struct {
}

func listFile(path string) ([]byte, error) {
	cmd := exec.Command("ls", "-ll")

	// Old LIST command is the output of `ls` command
	// New way for listing a directory is MLS*, described in RFC3659
	cmd.Env = append(cmd.Env, "LANG=en_US.utf-8")
	cmd.Dir = path
	return cmd.Output()
}

func (h *dtpListRequestHandler) handle(dtpConn *DtpConn, request interface{}) {
	listRequest := request.(DtpListRequest)
	dtpConn.logf("LIST command received.\n")
	res, err := listFile(listRequest.Path)
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

	dtpConn.msgChan <- dtpSuccessResponse{}

	dtpConn.logf("Finish LIST")
}

var errInvalidDtpRequest = errors.New("invalid request")

func getDtpRequestHandler(request interface{}) (dtpRequestHandler, error) {
	switch request.(type) {
	case DtpListRequest:
		return &dtpListRequestHandler{}, nil
	case DtpTransferRequest:
		return &dtpTransferRequestHandler{}, nil
	}

	return nil, errInvalidDtpRequest
}
