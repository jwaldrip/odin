package cli

type subcommandable struct {
  *writer

  subCommands map[string]*SubCommand
  subCommandsParsed bool
}

func (this *subcommandable) parseSubCommands(args []string) bool {
  return false
}
