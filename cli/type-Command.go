package cli

type Command interface {
  Name() string
  Flag(string) Value
  Flags() map[string]Value
  Param(string) Value
  Params() map[string]Value
  Start(...string)
  // NewSubCommand(string, *commandFn, ...string ) *SubCommand
}
