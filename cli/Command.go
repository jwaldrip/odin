package cli

import (
	"io"

	"github.com/jwaldrip/odin/cli/values"
)

// Command represents a readable command
type Command interface {
	// Freeform Arguments
	Arg(int) values.Value
	Args() values.List

	// Attributes
	Name() string
	NameAliases() map[string]string
	Description() string
	Version() string
	Hidden() bool

	// Flags
	Flag(string) values.Value
	Flags() values.Map

	// Params
	Param(string) values.Value
	Params() values.Map

	// Outputs
	ErrPrint(...interface{})
	ErrPrintf(string, ...interface{})
	ErrPrintln(...interface{})
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
	StdOut() io.Writer
	StdErr() io.Writer

	// Misc
	Parent() Command
	Usage()
}
