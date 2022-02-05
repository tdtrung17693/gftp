package handlers

import (
	"fmt"
	"gftp/commands"
	"gftp/server"
	"os/exec"
	"path"
	"strings"
)

type mkdCmdHandler struct {
}

func makeDir(path string) error {
	cmd := exec.Command("mkdir", path)

	return cmd.Run()
}

func (h mkdCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	// Mark the transfer starting point
	clientDirName := strings.Trim(strings.Join(cmd.Params, " "), "/.")

	dirPath := ctx.GetServerPath(path.Join(ctx.Pwd, clientDirName))
	err := makeDir(dirPath)

	if err != nil {
		return &server.Response{
			Code:    server.ReplyRequestedFileUnavailableOrBusy,
			Message: err.Error(),
		}
	}

	return &server.Response{
		Code:    server.ReplyPathNameCreated,
		Message: fmt.Sprintf("%q is cwd", dirPath),
	}
}
