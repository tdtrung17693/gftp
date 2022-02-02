package dtp

import (
	"errors"
	"strings"
)

type TransferType string

const (
	TransferTypeASCII  = "ascii"
	TransferTypeEBCDIC = "ebcdic"
	TransferTypeImage  = "image"
	TransferTypeLocal  = "local"
)

func (transferType TransferType) IsValid() error {
	switch transferType {
	case TransferTypeASCII, TransferTypeImage, TransferTypeEBCDIC, TransferTypeLocal:
		return nil
	}

	return errors.New("invalid transfer type")
}

func TransferTypeFromParam(paramStr string) (TransferType, error) {
	switch strings.ToLower(paramStr) {
	case "i":
		return TransferTypeImage, nil
	case "e":
		return TransferTypeEBCDIC, nil
	case "a":
		return TransferTypeASCII, nil
	}

	return "", errors.New("invalid transfer type")
}
