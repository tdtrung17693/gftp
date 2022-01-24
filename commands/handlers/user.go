package handlers

import (
	"gftp/commands"
	"gftp/server"
)

type userCmdHandler struct {
}

func (h userCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	username := cmd.Params[0]
	exists, err := server.DefaultAuth.UserProvider.Exists(username)
	if err != nil {
		return &server.Response{
			Code:    server.ReplyNotLoggedIn,
			Message: err.Error(),
		}
	}

	if !exists {
		return &server.Response{
			Code:    server.ReplyNotLoggedIn,
			Message: "User not exist",
		}
	}

	ctx.Username = username
	msg := "Accepted. Wait for password"
	if username == "anonymous" {
		msg = "Anonymous mode. Provide email to continue"
	}
	return &server.Response{
		Code:    server.ReplyNeedPassword,
		Message: msg,
	}
}
