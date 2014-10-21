package cli

// DefineSubCommand defines and adds a SubCommand on the current command
func (cmd *CLI) DefineSubCommand(name string, desc string, fn func(Command), paramNames ...string) *SubCommand {
	return cmd.AddSubCommand(
		NewSubCommand(name, desc, fn, paramNames...),
	)
}

// AddSubCommands adds subcommands to a command
func (cmd *CLI) AddSubCommands(subcmds ...*SubCommand) {
	for _, subcmd := range subcmds {
		cmd.AddSubCommand(subcmd)
	}
}

// AddSubCommand adds a subcommand to a command
func (cmd *CLI) AddSubCommand(subcmd *SubCommand) *SubCommand {
	if cmd.subCommands == nil {
		cmd.subCommands = make(map[string]*SubCommand)
	}
	subcmd.errOutput = cmd.ErrOutput()
	subcmd.stdOutput = cmd.StdOutput()
	subcmd.ErrorHandling = cmd.ErrorHandling
	cmd.subCommands[subcmd.name] = subcmd
	if subcmd.parent != nil {
		panic("command already assigned")
	}
	if &subcmd.CLI == cmd {
		panic("cannot assign subcmd to itself as a subcmd")
	}
	subcmd.parent = cmd
	return subcmd
}
