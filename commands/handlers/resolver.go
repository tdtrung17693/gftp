package handlers

import "gftp/commands"

var cmdHandlerMap map[string](interface{})

func init() {
	cmdHandlerMap = make(map[string]interface{})

	cmdHandlerMap["user"] = newCommandHandler(&userCmdHandler{})
	cmdHandlerMap["pass"] = newCommandHandler(&passCmdHandler{})
	cmdHandlerMap["feat"] = newCommandHandler(&featCmdHandler{})
	cmdHandlerMap["syst"] = newCommandHandler(&systCmdHandler{})
	cmdHandlerMap["pasv"] = newCommandHandler(&pasvCmdHandler{})
	cmdHandlerMap["pwd"] = newCommandHandler(&pwdCmdHandler{})
	cmdHandlerMap["type"] = newCommandHandler(&typeCmdHandler{})
	cmdHandlerMap["list"] = newCommandHandler(&listCmdHandler{})
	cmdHandlerMap["unknown"] = newCommandHandler(&unknownCmdHandler{})
}

func newCommandHandler(c interface{}) interface{} {
	return c
}

type CommandResolver struct {
}

func (r CommandResolver) Resolve(cmd *commands.Command) commands.CommandHandler {
	if val, ok := cmdHandlerMap[cmd.Name]; ok {
		ret := val.(commands.CommandHandler)
		return ret
	}

	ret := cmdHandlerMap["unknown"].(commands.CommandHandler)
	return ret
}
