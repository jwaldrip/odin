package cli

// SubCommand is a subcommand to a cli
type SubCommand struct {
	CLI
}

// NewSubCommand Create a new subcommand instance
func NewSubCommand(name string, desc string, fn func(Command), paramNames ...string) *SubCommand {
	var cmd SubCommand
	cmd.init(name, desc, fn, paramNames...)
	return &cmd
}
