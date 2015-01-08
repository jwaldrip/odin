package cli

import "strings"

// Description returns the command description
func (cmd *CLI) Description() string {
	return cmd.description
}

// LongDescription is the long description for a command.
func (cmd *CLI) LongDescription() string {
	return strings.TrimSpace(cmd.longDescription)
}

// Parent Returns the parent command
func (cmd *CLI) Parent() Command {
	return cmd.parent
}

// Name returns the command name
func (cmd *CLI) Name() string {
	return cmd.name
}

// SetLongDescription sets the long desription. This will will replace the
// description when viewing the --help for a command.
func (cmd *CLI) SetLongDescription(desc string) {
	cmd.longDescription = desc
}

//SetUsage lets your override the usagestring
func (cmd *CLI) SetUsage(use func()) {
	cmd.usage = use
}

// Version returns the command version
func (cmd *CLI) Version() string {
	return cmd.version
}

// NameAliases sets the name aliases
func (cmd *CLI) NameAliases() map[string]string {
	return cmd.nameAliases
}
