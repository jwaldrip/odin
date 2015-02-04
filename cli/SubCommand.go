package cli

// SubCommand is a subcommand to a cli
type SubCommand struct {
	CLI
}

// NewSubCommand Create a new subcommand instance. It takes a name, desc, command
// and params. If desc is equal to "none" or "hidden", this will not generate
// documentation for this command.
func NewSubCommand(name string, desc string, fn func(Command), paramNames ...string) *SubCommand {
	var cmd SubCommand
	cmd.init(name, desc, fn, paramNames...)
	return &cmd
}
