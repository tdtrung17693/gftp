package handlers

import (
	"gftp/commands"
	"gftp/server"
	"os/exec"
	"path"
	"strings"
)

type deleCmdHandler struct {
}

func deleteFile(path string) error {
	cmd := exec.Command("rm", path)

	return cmd.Run()
}

func (h deleCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	filePath := ctx.GetServerPath(path.Join(ctx.Pwd, strings.Join(cmd.Params[0:], " ")))

	err := deleteFile(filePath)

	if err != nil {
		return &server.Response{
			Code:    server.ReplyRequestedFileUnavailableOrBusy,
			Message: err.Error(),
		}
	}

	return &server.Response{
		Code:    server.ReplyPathNameCreated,
		Message: "file has been deleted",
	}
}
