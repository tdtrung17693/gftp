package dtp

import "errors"

type TransferMode string

const (
	TransferModePassive = "transfer_mode_passive"
	TransferModeActive  = "transfer_mode_active"
)

var errInvalidTransferMode = errors.New("invalid transfer mode")

func (transferMode TransferMode) IsValid() error {
	switch transferMode {
	case TransferModeActive, TransferModePassive:
		return nil
	}

	return errInvalidTransferMode
}
