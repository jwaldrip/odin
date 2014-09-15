package cli

type Command interface {
  Name() string
  Flag(string) FlagValue
  Flags() map[string]FlagValue
  Param(string) string
  Params() map[string]string
  Start(args []string)
  // NewSubCommand(string, *commandFn, ...string ) *SubCommand
}
