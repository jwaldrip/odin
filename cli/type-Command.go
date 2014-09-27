package cli

// Command represents a readable command
type Command interface {
	Usage()
	Arg(int) Value
	Args() []Value
	Parent() Command
	Name() string
	Flag(string) Value
	Flags() map[string]Value
	Param(string) Value
	Params() map[string]Value
	Start(...string)
}
