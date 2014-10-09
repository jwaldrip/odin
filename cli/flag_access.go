package cli

import (
	"fmt"

	"github.com/jwaldrip/odin/cli/values"
)

// Flag returns the Value interface to the value of the named flag,
// returning nil if none exists.
func (cmd *CLI) Flag(name string) values.Value {
	flag, ok := cmd.flags[name]
	if !ok {
		panic(fmt.Sprintf("flag not defined %v", name))
	}
	value := cmd.flagValues[flag]
	return value
}

// Flags returns the flags as a map of strings with Values
func (cmd *CLI) Flags() values.Map {
	flags := make(values.Map)
	for name := range cmd.flags {
		flags[name] = cmd.Flag(name)
	}
	return flags
}
