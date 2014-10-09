package cli

import "github.com/jwaldrip/odin/cli/values"

// Args returns any remaining args that were not parsed as params
func (cmd *CLI) Args() values.List {
	return cmd.unparsedArgs
}

// Arg takes a position of a remaining arg that was not parsed as a param
func (cmd *CLI) Arg(index int) values.Value {
	return cmd.Args()[index]
}
