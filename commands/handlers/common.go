package handlers

import (
	. "gftp/commands"
)

type CommandHandler interface {
	Handle(cmd *Command, ctx *ConnContext) Response
}

type cmdBuilder (func() interface{})

var cmdHandlerMap map[string](*CommandHandler)

func init() {
	cmdHandlerMap = make(map[string]*CommandHandler)

	cmdHandlerMap["user"] = newCommandHandler(&userCmdHandler{}).(*CommandHandler)
	cmdHandlerMap["pass"] = newCommandHandler(&passCmdHandler{}).(*CommandHandler)
	cmdHandlerMap["feat"] = newCommandHandler(&featCmdHandler{}).(*CommandHandler)
	cmdHandlerMap["syst"] = newCommandHandler(&systCmdHandler{}).(*CommandHandler)
	cmdHandlerMap["pasv"] = newCommandHandler(&pasvCmdHandler{}).(*CommandHandler)

}

func newCommandHandler(c interface{}) interface{} {
	return &c
}

func GetCommandHandler(cmd *Command) *CommandHandler {}
