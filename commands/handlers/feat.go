package handlers

import (
	. "gftp/commands"
)

type featCmdHandler struct {
	command Command
}

func (h *featCmdHandler) Handle(cmd *Command, ctx *ConnContext) *Response {

}
