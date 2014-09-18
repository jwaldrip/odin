package cli

type SubCommand struct {
  CLI
  parents []Command
}

func NewSubCommand(name string, fn commandFn, paramNames ...string) *SubCommand {
  var cmd SubCommand
  cmd.setName(name)
  cmd.setParams(paramNames...)
  cmd.setFn(fn)
  return &cmd
}
