package cli

import "fmt"
import "os"

// A FlagSet represents a set of defined flags.  The zero value of a FlagSet
// has no name and has ContinueOnError error handling.
type CLI struct {
  flags
  params
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
  f := &CLI{name: os.Args[0]}
  f.setFn(fn)
  f.setParams(paramNames...)
  return f
}

func (f *CLI) Start(args ...string) {
  if args == nil {
    args = os.Args
  }
  ExitIfError(f.parseFlags(args))
  ExitIfError(f.parseParams(args))
}

func (f *CLI) setName(name string) {
  f.name = name
}

func (f *CLI) setFn(fn commandFn){
  f.fn = fn
}

// PrintDefaults prints, to standard error unless configured
// otherwise, the default values of all defined flags in the set.
// func (f *FlagSet) PrintDefaults() {
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
func (f *CLI) usage() {
  if f.Usage == nil {
    f.defaultUsage()
  } else {
    f.Usage()
  }
}

// Parsed reports whether f.Parse has been called.
func (f *CLI) Parsed() bool {
  return f.parsed
}

// Parsed reports whether f.Parse has been called.
func (f *CLI) defaultUsage() {
  fmt.Println("Default usage")
}

// Init sets the name and error handling property for a flag set.
// By default, the zero CLI uses an empty name and the
// ContinueOnError error handling policy.
func (f *CLI) Init(name string, errorHandling ErrorHandling) {
  f.name = name
  f.ErrorHandling = errorHandling
}
