package server

type CommandProcessor interface {
	Handle(string, *ConnContext) *Response
}
