package cli

import "github.com/jwaldrip/odin/cli/values"

// Command represents a readable command
type Command interface {
	Arg(int) values.Value
	Args() values.List
	Description() string
	Flag(string) values.Value
	Flags() values.Map
	Name() string
	Param(string) values.Value
	Params() values.Map
	Parent() Command
	Start(...string)
	Usage()
	Version() string
}
