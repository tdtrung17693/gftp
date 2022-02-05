package handlers

import (
	"gftp/commands"
	"gftp/server"
	"log"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type rmdCmdHandler struct {
}

func deleteDir(path string) error {
	cmd := exec.Command("rm", "-rf", path)

	return cmd.Run()
}

func isSubDir(parentDir string, testDir string) bool {
	rel, err := filepath.Rel(parentDir, testDir)
	if err != nil {
		log.Printf("error: %q", err)
		return false
	}

	return !strings.Contains(rel, "..")
}

func (h rmdCmdHandler) Handle(cmd *commands.Command, ctx *server.ConnContext) *server.Response {
	filePath := ctx.GetServerPath(path.Join(ctx.Pwd, strings.Join(cmd.Params[0:], " ")))

	if !isSubDir(ctx.GetUserServerPath(), filePath) {
		return &server.Response{
			Code:    server.ReplyArgumentSyntaxError,
			Message: "invalid path",
		}
	}

	err := deleteDir(filePath)

	if err != nil {
		return &server.Response{
			Code:    server.ReplyRequestedFileUnavailableOrBusy,
			Message: err.Error(),
		}
	}

	return &server.Response{
		Code:    server.ReplyPathNameCreated,
		Message: "file has been rmdted",
	}
}
