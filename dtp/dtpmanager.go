package dtp

import (
	"fmt"
	"log"
)

type dtpOpenConnRequest struct {
	transferType TransferType
	responseChan chan interface{}
}

type Dtp struct {
	listenerChan chan interface{}
	errorChan    chan error
}

func (dtp *Dtp) OpenNewConn(transferType TransferType) (*DtpConn, error) {
	responseChan := make(chan interface{})
	dtp.listenerChan <- dtpOpenConnRequest{
		transferType: transferType,
		responseChan: responseChan,
	}

	response := <-responseChan

	switch val := response.(type) {
	case *DtpConn:
		return val, nil
	}

	// if all other response type is not matched,
	// then an error has been occurred
	return nil, response.(error)
}

func NewDTP() *Dtp {
	dtpChan := make(chan interface{}, 5)
	errChan := make(chan error)

	go func() {
		for msg := range dtpChan {
			switch val := msg.(type) {
			case dtpOpenConnRequest:
				res, err := newDtpConn()

				if err != nil {
					val.responseChan <- err
					log.Printf("[NewDTP] Error while creating DTP conn: %q", err)
				}
				go res.serve()
				val.responseChan <- res
			default:
				errChan <- fmt.Errorf("unknown DTP request (Received: %q)", msg)
			}
		}
	}()

	return &Dtp{
		listenerChan: dtpChan,
		errorChan:    errChan,
	}
}
