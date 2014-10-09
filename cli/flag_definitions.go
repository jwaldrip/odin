package cli

import (
	"fmt"
	"time"

	"github.com/jwaldrip/odin/cli/values"
)

// AliasFlag creates an alias from a flag
func (cmd *CLI) AliasFlag(alias rune, flagname string) {
	flag, ok := cmd.flags[flagname]
	if !ok {
		panic(fmt.Sprintf("flag not defined %v", flagname))
	}
	if cmd.aliases == nil {
		cmd.aliases = make(map[rune]*Flag)
	}
	cmd.aliases[alias] = flag
}

// DefineBoolFlag defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (cmd *CLI) DefineBoolFlag(name string, value bool, usage string) *bool {
	p := new(bool)
	cmd.DefineBoolFlagVar(p, name, value, usage)
	return p
}

// DefineBoolFlagVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (cmd *CLI) DefineBoolFlagVar(p *bool, name string, value bool, usage string) {
	cmd.DefineFlag(values.NewBool(value, p), name, usage)
}

// DefineDurationFlag defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
func (cmd *CLI) DefineDurationFlag(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	cmd.DefineDurationFlagVar(p, name, value, usage)
	return p
}

// DefineDurationFlagVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
func (cmd *CLI) DefineDurationFlagVar(p *time.Duration, name string, value time.Duration, usage string) {
	cmd.DefineFlag(values.NewDuration(value, p), name, usage)
}

// DefineFloat64Flag defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func (cmd *CLI) DefineFloat64Flag(name string, value float64, usage string) *float64 {
	p := new(float64)
	cmd.DefineFloat64FlagVar(p, name, value, usage)
	return p
}

// DefineFloat64FlagVar defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func (cmd *CLI) DefineFloat64FlagVar(p *float64, name string, value float64, usage string) {
	cmd.DefineFlag(values.NewFloat64(value, p), name, usage)
}

// DefineInt64Flag defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func (cmd *CLI) DefineInt64Flag(name string, value int64, usage string) *int64 {
	p := new(int64)
	cmd.DefineInt64FlagVar(p, name, value, usage)
	return p
}

// DefineInt64FlagVar defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func (cmd *CLI) DefineInt64FlagVar(p *int64, name string, value int64, usage string) {
	cmd.DefineFlag(values.NewInt64(value, p), name, usage)
}

// DefineIntFlag defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func (cmd *CLI) DefineIntFlag(name string, value int, usage string) *int {
	p := new(int)
	cmd.DefineIntFlagVar(p, name, value, usage)
	return p
}

// DefineIntFlagVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func (cmd *CLI) DefineIntFlagVar(p *int, name string, value int, usage string) {
	cmd.DefineFlag(values.NewInt(value, p), name, usage)
}

// DefineStringFlag defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (cmd *CLI) DefineStringFlag(name string, value string, usage string) *string {
	p := new(string)
	cmd.DefineStringFlagVar(p, name, value, usage)
	return p
}

// DefineStringFlagVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (cmd *CLI) DefineStringFlagVar(p *string, name string, value string, usage string) {
	cmd.DefineFlag(values.NewString(value, p), name, usage)
}

// DefineUint64Flag defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func (cmd *CLI) DefineUint64Flag(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	cmd.DefineUint64FlagVar(p, name, value, usage)
	return p
}

// DefineUint64FlagVar defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (cmd *CLI) DefineUint64FlagVar(p *uint64, name string, value uint64, usage string) {
	cmd.DefineFlag(values.NewUint64(value, p), name, usage)
}

// DefineUintFlag defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (cmd *CLI) DefineUintFlag(name string, value uint, usage string) *uint {
	p := new(uint)
	cmd.DefineUintFlagVar(p, name, value, usage)
	return p
}

// DefineUintFlagVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (cmd *CLI) DefineUintFlagVar(p *uint, name string, value uint, usage string) {
	cmd.DefineFlag(values.NewUint(value, p), name, usage)
}

// DefineFlag defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
func (cmd *CLI) DefineFlag(value values.Value, name string, usage string) {
	// Remember the default value as a string; it won't change.
	flag := &Flag{
		Name:     name,
		Usage:    usage,
		value:    value,
		DefValue: value.String(),
	}
	_, alreadythere := cmd.flags[name]
	if alreadythere {
		cmd.panicf("flag redefined: %s", name)
	}
	if cmd.flags == nil {
		cmd.flags = make(flagMap)
	}
	cmd.flags[name] = flag
}
