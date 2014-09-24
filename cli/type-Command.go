package cli

// Command represents a readable command
type Command interface {
	Usage()
	Parent() Command
	Name() string
	Flag(string) Getter
	Flags() map[string]Getter
	Param(string) Getter
	Params() map[string]Getter
	Start(...string)
}
