package cli

type subcommandable struct {
  *writer

  subCommands map[string]*SubCommand
  subCommandsParsed bool
}

func (this *subcommandable) parseSubCommands(args []string) bool {
  return false
}

func (this *subcommandable) SubCommandsUsage() string {
  return "  command usage"
}

func (this *subcommandable) DefineSubCommand(name string, fn commandFn, paramNames ...string) {
  if this.subCommands == nil {
    this.subCommands = make(map[string]*SubCommand)
  }
  var cmd SubCommand
  cmd.init(name, fn, paramNames...)
  this.subCommands[name] = &cmd
}
