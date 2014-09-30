package cli

// Command represents a readable command
type Command interface {
	Usage()
	Arg(int) Value
	Args() ValueList
	Parent() Command
	Name() string
	Flag(string) Value
	Flags() ValueMap
	Param(string) Value
	Params() ValueMap
	Start(...string)
}
