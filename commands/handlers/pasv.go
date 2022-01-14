package handlers

import (
	. "gftp/commands"
)

type pasvCmdHandler struct {
	command Command
}

func (h *pasvCmdHandler) Handle(cmd *Command, ctx *ConnContext) *Response {

}
