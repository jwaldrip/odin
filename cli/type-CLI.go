package cli

import "fmt"
import "os"

// A FlagSet represents a set of defined flags.  The zero value of a FlagSet
// has no name and has ContinueOnError error handling.
type CLI struct {
  flagable
  paramable
  subcommandable
  writer

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
  cli := &CLI{name: os.Args[0]}
  cli.setFn(fn)
  cli.setParams(paramNames...)
  return cli
}

func (this *CLI) Start(args ...string) {
  if args == nil {
    args = os.Args
  }
  var err error
  args, err = this.parseFlags(args)
  exitIfError(err)
  args, err = this.parseParams(args)
  exitIfError(err)
  args, err = this.parseSubCommands(args)
  exitIfError(err)
}

func (this *CLI) setName(name string) {
  this.name = name
}

func (this *CLI) setFn(fn commandFn){
  this.fn = fn
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

// usage calls the Usage method for the flag set
func (this *CLI) usage() {
  if this.Usage == nil {
    this.defaultUsage()
  } else {
    this.Usage()
  }
}

// Parsed reports whether f.Parse has been called.
func (this *CLI) Parsed() bool {
  allParsed := this.flagsParsed && this.paramsParsed && this.subCommandsParsed
  return allParsed
}

// Parsed reports whether f.Parse has been called.
func (this *CLI) defaultUsage() {
  fmt.Println("Default usage")
}

// Init sets the name and error handling property for a flag set.
// By default, the zero CLI uses an empty name and the
// ContinueOnError error handling policy.
func (this *CLI) Init(name string, errorHandling ErrorHandling) {
  this.name = name
  this.ErrorHandling = errorHandling
}
