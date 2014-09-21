package cli

type subcommandable struct {
  subCommands map[string]*SubCommand
  subCommandsParsed bool
}

func (this *subcommandable) parseSubCommands(args []string) ([]string, error){
  return args, nil
}
