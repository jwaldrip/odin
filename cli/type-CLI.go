package cli

import "os"

// A FlagSet represents a set of defined flags.  The zero value of a FlagSet
// has no name and has ContinueOnError error handling.
type CLI struct {
  flagable
  paramable
  subcommandable
  *writer

  // Usage is the function called when an error occurs while parsing flags.
  // The field is a function (not a method) that may be changed to point to
  // a custom error handler.
  Usage func()

  fn            commandFn
  name          string
  parsed        bool
  subCommands   []Command
}

// NewFlagSet returns a new, empty flag set with the specified name and
// error handling property.
func NewCLI(fn commandFn, paramNames ...string) *CLI {
  writer := &writer{}
  cli := &CLI{
    writer:         writer,
    flagable:       flagable{writer: writer},
    paramable:      paramable{writer: writer},
    subcommandable: subcommandable{writer: writer},
    name: os.Args[0],
  }
  cli.setFn(fn)
  cli.setParams(paramNames...)
  return cli
}

func (this *CLI) Start(args ...string) {
  if args == nil {
    args = os.Args[1:]
  }
  args = this.parseFlags(args)
  args = this.parseParams(args)
  ransubcommand := this.parseSubCommands(args)
  if !ransubcommand {
    this.fn(this)
  }
}

func (this *CLI) setName(name string) {
  this.name = name
}

func (this *CLI) setFn(fn commandFn){
  this.fn = fn
}

func (this *CLI) Name() string {
  return this.name
}

// PrintDefaults prints, to standard error unless configured
// otherwise, the default values of all defined flags in the set.
// func (this *FlagSet) PrintDefaults() {
//   f.VisitAll(func(flag *Flag) {
//     format := "  -%s=%s: %s\n"
//     if _, ok := flag.Value.(*stringValue); ok {
//       // put quotes on the value
//       format = "  -%s=%q: %s\n"
//     }
//     fmt.Fprintf(f.out(), format, flag.Name, flag.DefValue, flag.Usage)
//   })
// }

// Parsed reports whether f.Parse has been called.
func (this *CLI) Parsed() bool {
  allParsed := this.flagsParsed && this.paramsParsed && this.subCommandsParsed
  return allParsed
}
