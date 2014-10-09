package cli

import "strings"

// Description returns the command description
func (cmd *CLI) Description() string {
	return cmd.description
}

// Parent Returns the parent command
func (cmd *CLI) Parent() Command {
	return cmd.parent
}

// Name returns the command name
func (cmd *CLI) Name() string {
	var name string
	if cmd.parent != nil {
		name = strings.Join([]string{cmd.parent.Name(), cmd.name}, " ")
	} else {
		name = cmd.name
	}
	return name
}

// Version returns the command version
func (cmd *CLI) Version() string {
	return cmd.version
}
