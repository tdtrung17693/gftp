package commands

type unknownCommand struct{}

func (c *unknownCommand) Handle() {

}

func GetCommandHandler(cmd *Command) *CommandHandler {}
