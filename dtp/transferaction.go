package dtp

import "errors"

type TransferAction string

var errInvalidTransferAction = errors.New("invalid transfer action")

const (
	TransferActionStore    = "transfer_action_store"
	TransferActionRetrieve = "transfer_action_retrieve"
)

func (transferAction TransferAction) IsValid() error {
	switch transferAction {
	case TransferActionStore, TransferActionRetrieve:
		return nil
	}

	return errInvalidTransferAction
}
