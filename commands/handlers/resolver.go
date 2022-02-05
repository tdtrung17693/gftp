package handlers

import (
	"gftp/commands"
	"log"
)

var cmdHandlerMap map[string](interface{})

func init() {
	cmdHandlerMap = make(map[string]interface{})

	cmdHandlerMap["user"] = newCommandHandler(&userCmdHandler{})
	cmdHandlerMap["pass"] = newCommandHandler(&passCmdHandler{})
	cmdHandlerMap["feat"] = newCommandHandler(&featCmdHandler{})
	cmdHandlerMap["syst"] = newCommandHandler(&systCmdHandler{})
	cmdHandlerMap["pasv"] = newCommandHandler(&pasvCmdHandler{})
	cmdHandlerMap["epsv"] = newCommandHandler(&epsvCmdHandler{})
	cmdHandlerMap["pwd"] = newCommandHandler(&pwdCmdHandler{})
	cmdHandlerMap["cwd"] = newCommandHandler(&cwdCmdHandler{})
	cmdHandlerMap["type"] = newCommandHandler(&typeCmdHandler{})
	cmdHandlerMap["list"] = newCommandHandler(&listCmdHandler{})
	cmdHandlerMap["retr"] = newCommandHandler(&retrCmdHandler{})
	cmdHandlerMap["stor"] = newCommandHandler(&storCmdHandler{})
	cmdHandlerMap["mkd"] = newCommandHandler(&mkdCmdHandler{})
	cmdHandlerMap["rmd"] = newCommandHandler(&rmdCmdHandler{})
	cmdHandlerMap["dele"] = newCommandHandler(&deleCmdHandler{})
	cmdHandlerMap["unknown"] = newCommandHandler(&unknownCmdHandler{})
}

func newCommandHandler(c interface{}) interface{} {
	return c
}

type CommandResolver struct {
}

func (r CommandResolver) Resolve(cmd *commands.Command) commands.CommandHandler {
	log.Printf("[command resolver] handling command %s", cmd.Name)
	if val, ok := cmdHandlerMap[cmd.Name]; ok {
		ret := val.(commands.CommandHandler)
		return ret
	}

	ret := cmdHandlerMap["unknown"].(commands.CommandHandler)
	return ret
}
