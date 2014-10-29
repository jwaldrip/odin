package cli

func (cmd *CLI) parseSubCommands(args []string) ([]string, bool) {
	if len(args) == 0 || len(cmd.subCommands) == 0 {
		return args, false
	}
	name := args[0]
	subcmd, err := cmd.subCommands.Get(name)
	if err != nil {
		cmd.errf("invalid command: %s", name)
		return args, false
	}

	// Inherit Outputs
	if subcmd.errOutput == nil {
		subcmd.errOutput = cmd.errOutput
	}
	if subcmd.stdOutput == nil {
		subcmd.stdOutput = cmd.stdOutput
	}

	subcmd.Start(args...)

	return []string{}, true
}
