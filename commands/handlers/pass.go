package handlers

import (
	"gftp/commands"
	"gftp/server"
	"log"
)

type passCmdHandler struct {
}

func (h passCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	password := cmd.Params[0]

	if ctx.Username != "anonymous" {
		isValid, err := server.DefaultAuth.AuthService.Verify(ctx.Username, password)

		if err != nil {
			return &server.Response{
				Code:    server.ReplyNotLoggedIn,
				Message: err.Error(),
			}
		}

		if !isValid {
			log.Print(err)
			return &server.Response{
				Code:    server.ReplyNotLoggedIn,
				Message: "Invalid credentials",
			}
		}
	}

	user, err := server.DefaultAuth.UserProvider.GetUser(ctx.Username)

	if err != nil {
		log.Print(err)
		return &server.Response{
			Code:    server.ReplyArgumentSyntaxError,
			Message: "Server error",
		}
	}

	ctx.UserRoot = user.RootPath
	ctx.Pwd = ctx.UserRoot

	return &server.Response{
		Code:    server.ReplyUserLoggedIn,
		Message: "Ready",
	}
}
