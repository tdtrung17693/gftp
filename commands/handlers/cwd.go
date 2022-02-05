package handlers

import (
	"fmt"
	"gftp/commands"
	"gftp/server"
	"path"
	"strings"
)

type cwdCmdHandler struct {
}

func (h cwdCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	target := strings.Join(cmd.Params, " ")

	if target[0] != '/' && !strings.Contains(target, ctx.UserRoot) {
		target = fmt.Sprintf("%s/%s", ctx.Pwd, target)
	} else {
		target = strings.ReplaceAll(target, ctx.UserRoot, "")
	}

	if target[0:] == ".." && ctx.Pwd == "/" {
		target = "/"
	}
	ctx.Pwd = target

	return &server.Response{
		Code:    server.ReplyRequestedFileOk,
		Message: fmt.Sprintf("\"%s\" is cwd.", path.Join(ctx.UserRoot, ctx.Pwd)),
	}
}
