package cli

type SubCommand struct {
  CLI
  parents []Command
}

func NewSubCommand(name string, fn commandFn, args ...string ) *SubCommand {
  var cmd SubCommand
  cmd.name = name
  cmd.args = args
  cmd.fn = fn
  cmd.initFlagSet()
  return &cmd
}
