package handlers

import (
	. "gftp/commands"
)

type systCmdHandler struct {
	command Command
}

func (h *systCmdHandler) Handle(cmd *Command, ctx *ConnContext) *Response {

}
