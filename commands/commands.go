package commands

import (
	"strings"
)

type Command struct {
	Name   string
	Params []string
}

func parseCommand(requestLine string) *Command {
	components := strings.Split(requestLine, " ")

	var params []string
	if len(components) > 1 {
		params = components[1:]
	}

	return &Command{
		Name:   strings.ToLower(components[0]),
		Params: params,
	}
}
