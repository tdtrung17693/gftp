package dtp

type DtpListRequest struct {
	Path   string
	MsgBus chan interface{}
}

type DtpTransferRequest struct {
	FilePath       string
	TransferAction TransferAction
	TransferType   TransferType
}
