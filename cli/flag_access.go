package cli

import (
	"fmt"

	"github.com/jwaldrip/odin/cli/values"
)

// Flag returns the Value interface to the value of the named flag,
// returning nil if none exists.
func (cmd *CLI) Flag(name string) values.Value {
	flag := cmd.getFlag(name)
	value := cmd.flagValues[flag]
	return value
}

// Flags returns the flags as a map of strings with Values
func (cmd *CLI) Flags() values.Map {
	flags := make(values.Map)
	for name := range cmd.inheritedFlags.Merge(cmd.flags) {
		flags[name] = cmd.Flag(name)
	}
	return flags
}

func (cmd *CLI) getFlag(name string) *Flag {
	flag, exists := cmd.inheritedFlags.Merge(cmd.flags)[name]
	if !exists {
		panic(fmt.Sprintf("flag not defined %v", name))
	}
	return flag
}

func (cmd *CLI) hasFlags() bool {
	var internalFlagCount int
	if flag, ok := cmd.flags["help"]; ok && flag == cmd.flagHelp {
		internalFlagCount++
	}
	if flag, ok := cmd.flags["version"]; ok && flag == cmd.flagVersion {
		internalFlagCount++
	}
	return len(cmd.flags) > internalFlagCount
}
