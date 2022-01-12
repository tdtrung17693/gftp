package commands_test

import (
	"gftp/commands"
	"testing"
)

func TestParseRequest(t *testing.T) {
  commands.ParseRequest("USER abc")
}
