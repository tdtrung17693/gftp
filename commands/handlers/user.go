package handlers

import (
	. "gftp/commands"
)

type userCmdHandler struct {
	command Command
}

func (h *userCmdHandler) Handle(cmd *Command, ctx *ConnContext) *Response {

}
