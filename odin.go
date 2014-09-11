package odin

import "flag"
import "log"

type commandFn func(*Command) error

type Command struct {
  Name string
  Params map[string]string
  Flags map[string]flag.Value

  args []string
  commandFn commandFn
  err flag.ErrorHandling
  flagSet flag.FlagSet
  parent *Command
  subCommands map[string]*Command
}

// NewCommand("say", *say, "greeting", "object")
func NewCommand( name string, fn *commandFn, args ...string ) *Command {
  var cmd Command
  cmd.Name = name
  cmd.args = args
  cmd.flagSet = *flag.NewFlagSet(name, cmd.err)
  return &cmd
}

func (cmd *Command) NewSubCommand( name string, fn *commandFn, args ...string ) *Command {
  subcmd := NewCommand(name, fn, args...)
  cmd.subCommands[name] = subcmd
  return subcmd
}

func (cmd *Command) Parents() []*Command {
  var parents []*Command
  for {
    if len(cmd.parent.Name) == 0 { break }
    parents = append(parents, cmd.parent)
    cmd = cmd.parent
  }
  return parents
}

func (cmd *Command) Start(args []string){
  if cmd.Name != args[0] { log.Panicln("invalid command") }
  cmd.parseArgs(args[1:])
}

func (cmd *Command) StartDefault(args []string){
  args = append([]string{cmd.Name}, args...)
  cmd.Start(args)
}

func (cmd *Command) parseArgs(args []string) []string {
  cmd.flagSet.Parse(args)
  argLen := len(cmd.args)
  for i := 0 ; i <= argLen ; i++ {
    cmd.Params[cmd.args[i]] = args[i]
  }
  cmd.flagSet.VisitAll(func(f *flag.Flag){
    cmd.Flags[f.Name] = f.Value
  })
  return cmd.flagSet.Args()[argLen+1:]
}
