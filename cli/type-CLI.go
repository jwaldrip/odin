package cli

import "os"
import "fmt"
import "bytes"

// A FlagSet represents a set of defined flags.  The zero value of a FlagSet
// has no name and has ContinueOnError error handling.
type CLI struct {
  flagable
  paramable
  subcommandable
  *writer

  fn            commandFn
  name          string
  parsed        bool
}

// NewFlagSet returns a new, empty flag set with the specified name and
// error handling property.
func NewCLI(fn commandFn, paramNames ...string) *CLI {
  cli := new(CLI)
  cli.init(os.Args[0], fn, paramNames...)
  cli.Version = "v0.0.1"
  return cli
}

func (this *CLI) init(name string, fn commandFn, paramNames ...string){
  writer := &writer{}
  this.writer         = writer
  this.flagable       = flagable{writer: writer}
  this.paramable      = paramable{writer: writer}
  this.subcommandable = subcommandable{writer: writer}
  this.name           = name
  this.fn             = fn
  this.setParams(paramNames...)
  this.usage = func() { fmt.Println(this.CLIUsage()) }
}

func (this *CLI) Start(args ...string) {
  if args == nil {
    args = os.Args[1:]
  }
  args = this.parseFlags(args)
  args = this.parseParams(args)
  if this.Flag("version").Get() == true {
    fmt.Println(this.Name(), this.Version)
    return
  }
  if this.Flag("help").Get() == true{
    this.Usage()
    return
  }
  ransubcommand := this.parseSubCommands(args)
  if !ransubcommand {
    this.fn(this)
  }
}

func (this *CLI) setName(name string) {
  this.name = name
}

func (this *CLI) Name() string {
  return this.name
}

func (this *CLI) CLIUsage() string {
  var buff bytes.Buffer
  hasSubCommands := len(this.subCommands) > 0
  buff.WriteString("usage:\n")
  if hasSubCommands {
    buff.WriteString(fmt.Sprintf("  %s [options...] %s <command> [arg...]\n", this.Name(), this.ParamsUsage()))
  } else {
    buff.WriteString(fmt.Sprintf("  %s [options...] %s\n", this.Name(), this.ParamsUsage()))
  }
  buff.WriteString("\n")
  buff.WriteString("options:\n")
  buff.WriteString(this.FlagsUsage())
  if hasSubCommands {
    buff.WriteString("\n\n")
    buff.WriteString("commands:\n")
    buff.WriteString(this.SubCommandsUsage())
  }
  return buff.String()
}

// Parsed reports whether f.Parse has been called.
func (this *CLI) Parsed() bool {
  this.parsed = this.flagsParsed && this.paramsParsed && this.subCommandsParsed
  return this.parsed
}
