package cli

type subcommandable struct {
  subcommands []SubCommand
  subCommandsParsed bool
}

func (this *subcommandable) parseSubCommands(args []string) ([]string, error){
  return args, nil
}
