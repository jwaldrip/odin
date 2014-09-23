package cli

import "os"
import "fmt"
import "bytes"
import "strings"

// A FlagSet represents a set of defined flags.  The zero value of a FlagSet
// has no name and has ContinueOnError error handling.
type CLI struct {
  flagable
  paramable
  subcommandable
  *writer

  fn            commandFn
  name          string
  description   string
  parsed        bool
}

// NewFlagSet returns a new, empty flag set with the specified name and
// error handling property.
func NewCLI(fn commandFn, paramNames ...string) *CLI {
  cli := new(CLI)
  cli.init(os.Args[0], fn, paramNames...)
  cli.SetVersion("v0.0.1")
  return cli
}

func (this *CLI) init(name string, fn commandFn, paramNames ...string){
  writer := &writer{ErrorHandling: ExitOnError}
  this.writer         = writer
  this.flagable       = flagable{writer: writer}
  this.paramable      = paramable{writer: writer}
  this.subcommandable = subcommandable{writer: writer}
  this.name           = name
  this.fn             = fn
  this.setParams(paramNames...)
  this.usage = func() { fmt.Println(this.UsageString()) }
}

func (this *CLI) Start(args ...string) {
  if args == nil {
    args = os.Args
  }

  if len(args) > 1 {
    args = args[1:]
    args = this.parseFlags(args)
    args = this.parseParams(args)

    // Show a version
    if len(this.Version()) > 0 && this.Flag("version").Get() == true {
      fmt.Println(this.Name(), this.Version())
      return
    }

    // Show Help
    if this.Flag("help").Get() == true{
      this.Usage()
      return
    }

    // run subcommands
    ransubcommand := this.parseSubCommands(args)

    if ransubcommand {
      return
    }
  }

  // Run the function
  this.fn(this)
}

func (this *CLI) setName(name string) {
  this.name = name
}

func (this *CLI) SetDescription(desc string){
  this.description = desc
}

func (this *CLI) Name() string {
  var name string
  if this.parent != nil {
    name = strings.Join([]string{this.parent.Name(), this.name}, " ")
  } else {
    name = this.name
  }
  return name
}

func (this *CLI) Description() string {
  return this.description
}

func (this *CLI) UsageString() string {
  hasSubCommands := len(this.subCommands) > 0
  hasParams := len(this.Params()) > 0
  hasDescription := len(this.description) > 0

  // Start the Buffer
  var buff bytes.Buffer

  buff.WriteString("Usage:\n")
  buff.WriteString( fmt.Sprintf( "  %s [options...]", this.Name() ) )

  // Write Param Syntax
  if hasParams {
    buff.WriteString( fmt.Sprintf(" %s", this.paramable.UsageString() ) )
  }

  // Write Sub Command Syntax
  if hasSubCommands {
    buff.WriteString(" <command> [arg...]")
  }

  if hasDescription {
    buff.WriteString( fmt.Sprintf("\n\n%s", this.Description()) )
  }

  // Write Flags Syntax
  buff.WriteString("\n\nOptions:\n")
  buff.WriteString(this.flagable.UsageString())

  // Write Sub Command List
  if hasSubCommands {
    buff.WriteString("\n\nCommands:\n")
    buff.WriteString(this.subcommandable.UsageString())
  }

  // Return buffer as string
  return buff.String()
}

// Parsed reports whether f.Parse has been called.
func (this *CLI) Parsed() bool {
  this.parsed = this.flagsParsed && this.paramsParsed && this.subCommandsParsed
  return this.parsed
}

func (this *CLI) DefineSubCommand(name string, desc string, fn commandFn, paramNames ...string) {
  cmd := this.subcommandable.DefineSubCommand(name, desc, fn, paramNames...)
  cmd.parent = this
}
