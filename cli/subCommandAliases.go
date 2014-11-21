package cli

// AliasName allows you to call the subcommand by other names
func (cmd *CLI) AliasName(alias, subCommand string) {
	if cmd.nameAliases == nil {
		cmd.nameAliases = make(map[string]string)
	}
	cmd.nameAliases[alias] = subCommand
}
