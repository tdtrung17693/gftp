package handlers

import (
	. "gftp/commands"
)

type passCmdHandler struct {
	command Command
}

func (h *passCmdHandler) Handle(cmd *Command, ctx *ConnContext) *Response {

}
