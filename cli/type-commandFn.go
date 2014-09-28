package cli

// CommandFn is an alias for func(Command). It is the function type any CLI or
// SubCommand must use.
type CommandFn func(Command)
