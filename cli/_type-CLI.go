package cli

import "os"
import "fmt"

type CLI struct {
  FlagSet
  //Usage *func()
  //name string
  params map[string]string
  flags map[string]FlagValue
  args []string
  fn commandFn
  err ErrorHandling
  flagSet FlagSet
  subCommands map[string]*SubCommand
}

func NewCLI (fn commandFn, args ...string) *CLI {
  var cli CLI
  cli.name = os.Args[0]
  cli.args = args
  cli.fn = fn
  cli.initFlagSet()
  return &cli
}

func (this *CLI) initFlagSet() {
  this.flagSet = *NewFlagSet(this.name, 1)
  this.flagSet.Usage = this.defaultUsage
  this.Usage = &this.flagSet.Usage
}

func (this CLI) defaultUsage() {
  stdout.Println("Usage:")
  stdout.Println("", this.CommandUsage())
  stdout.Println("")
  stdout.Println("Options:")
  stdout.Println(this.FlagsUsage())
}

func (this *CLI) CommandUsage() string {
  wrappedArgs := fmt.Sprintf("%s [options]", this.name)
  for i := 0 ; i < len(this.args) ; i++ {
    wrappedArgs = fmt.Sprintf("%s <%s>", wrappedArgs, this.args[i])
  }
  return wrappedArgs
}

func (this *CLI) FlagsUsage() string {
  fmt.Println(this.flagSet.Lookup("good"))
  return "foo"
}

func (this CLI) Flag(key string) FlagValue {
  return this.flags[key]
}

func (this CLI) Flags() map[string]FlagValue {
  return this.flags
}

func (this CLI) Param(key string) string {
  return this.params[key]
}

func (this CLI) Params() map[string]string {
  return this.params
}

func (this *CLI) Name() string {
  return this.name
}

func (this *CLI) NewSubCommand( name string, fn commandFn, args ...string ) *SubCommand {
  subcmd := NewSubCommand(name, fn, args...)
  subcmd.parents = append(subcmd.parents, this)
  this.subCommands[name] = subcmd
  return subcmd
}

func (this *CLI) Run() {
  this.Start(os.Args)
}

func (this *CLI) Start(args []string) {
  args = args[1:]               // extract command
  args = this.parseFlags(args)   // extract flags
  args = this.parseParams(args)  // extract params
  this.invoke(args)
}

func (this *CLI) invoke(args []string) {
  if len(args) == 0 {
    this.fn(this)
  } else {
    this.invokeSubCommand(args)
  }
}

func (this *CLI) invokeSubCommand(args []string) {
  subcmd, ok := this.subCommands[args[0]]
  if !ok { fail("invalid command") }
  subcmd.Start(args[1:])
}

func (this *CLI) parseFlags(args []string) []string {
  this.flags = make(map[string]FlagValue)
  this.flagSet.Parse(args)
  this.flagSet.VisitAll(func(f *Flag){
    this.flags[f.Name] = f.Value
  })
  return this.flagSet.Args()
}

func (this *CLI) parseParams(args []string) []string {
  this.params = make(map[string]string)
  argLen := len(this.args)
  paramLen := len(args)
  for i := 0 ; i < argLen && i < paramLen ; i++ {
    key := this.args[i]
    value := args[i]
    this.params[key] = value
  }
  if paramLen > argLen {
    return args[argLen:]
  } else {
    return make([]string, 0)
  }
}
