package cli

import "io"
import "time"
import "os"
import "fmt"

// A FlagSet represents a set of defined flags.  The zero value of a FlagSet
// has no name and has ContinueOnError error handling.
type FlagSet struct {
  // Usage is the function called when an error occurs while parsing flags.
  // The field is a function (not a method) that may be changed to point to
  // a custom error handler.
  Usage func()

  name          string
  parsed        bool
  actual        map[string]*Flag
  formal        map[string]*Flag
  args          []string // arguments after flags
  errorHandling ErrorHandling
  output        io.Writer // nil means stderr; use out() accessor
}

// NewFlagSet returns a new, empty flag set with the specified name and
// error handling property.
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet {
  f := &FlagSet{
    name:          name,
    errorHandling: errorHandling,
  }
  return f
}

func (f *FlagSet) out() io.Writer {
  if f.output == nil {
    return os.Stderr
  }
  return f.output
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, os.Stderr is used.
func (f *FlagSet) SetOutput(output io.Writer) {
  f.output = output
}

// VisitAll visits the flags in lexicographical order, calling fn for each.
// It visits all flags, even those not set.
func (f *FlagSet) VisitAll(fn func(*Flag)) {
  for _, flag := range sortFlags(f.formal) {
    fn(flag)
  }
}

// Visit visits the flags in lexicographical order, calling fn for each.
// It visits only those flags that have been set.
func (f *FlagSet) Visit(fn func(*Flag)) {
  for _, flag := range sortFlags(f.actual) {
    fn(flag)
  }
}

// Lookup returns the Flag structure of the named flag, returning nil if none exists.
func (f *FlagSet) Lookup(name string) *Flag {
  return f.formal[name]
}

// Set sets the value of the named flag.
func (f *FlagSet) Set(name, value string) error {
  flag, ok := f.formal[name]
  if !ok {
    return fmt.Errorf("no such flag -%v", name)
  }
  err := flag.Value.Set(value)
  if err != nil {
    return err
  }
  if f.actual == nil {
    f.actual = make(map[string]*Flag)
  }
  f.actual[name] = flag
  return nil
}

// PrintDefaults prints, to standard error unless configured
// otherwise, the default values of all defined flags in the set.
func (f *FlagSet) PrintDefaults() {
  f.VisitAll(func(flag *Flag) {
    format := "  -%s=%s: %s\n"
    if _, ok := flag.Value.(*stringValue); ok {
      // put quotes on the value
      format = "  -%s=%q: %s\n"
    }
    fmt.Fprintf(f.out(), format, flag.Name, flag.DefValue, flag.Usage)
  })
}

// NFlag returns the number of flags that have been set.
func (f *FlagSet) NFlag() int { return len(f.actual) }

// Arg returns the i'th argument.  Arg(0) is the first remaining argument
// after flags have been processed.
func (f *FlagSet) Arg(i int) string {
  if i < 0 || i >= len(f.args) {
    return ""
  }
  return f.args[i]
}

// NArg is the number of arguments remaining after flags have been processed.
func (f *FlagSet) NArg() int { return len(f.args) }

// Args returns the non-flag arguments.
func (f *FlagSet) Args() []string { return f.args }

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string) {
  f.Var(newBoolValue(value, p), name, usage)
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (f *FlagSet) Bool(name string, value bool, usage string) *bool {
  p := new(bool)
  f.BoolVar(p, name, value, usage)
  return p
}

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func (f *FlagSet) IntVar(p *int, name string, value int, usage string) {
  f.Var(newIntValue(value, p), name, usage)
}

// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func (f *FlagSet) Int(name string, value int, usage string) *int {
  p := new(int)
  f.IntVar(p, name, value, usage)
  return p
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string) {
  f.Var(newInt64Value(value, p), name, usage)
}


// Int64 defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func (f *FlagSet) Int64(name string, value int64, usage string) *int64 {
  p := new(int64)
  f.Int64Var(p, name, value, usage)
  return p
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string) {
  f.Var(newUintValue(value, p), name, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (f *FlagSet) Uint(name string, value uint, usage string) *uint {
  p := new(uint)
  f.UintVar(p, name, value, usage)
  return p
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string) {
  f.Var(newUint64Value(value, p), name, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64 {
  p := new(uint64)
  f.Uint64Var(p, name, value, usage)
  return p
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (f *FlagSet) StringVar(p *string, name string, value string, usage string) {
  f.Var(newStringValue(value, p), name, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (f *FlagSet) String(name string, value string, usage string) *string {
  p := new(string)
  f.StringVar(p, name, value, usage)
  return p
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string) {
  f.Var(newFloat64Value(value, p), name, usage)
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func (f *FlagSet) Float64(name string, value float64, usage string) *float64 {
  p := new(float64)
  f.Float64Var(p, name, value, usage)
  return p
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
  f.Var(newDurationValue(value, p), name, usage)
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration {
  p := new(time.Duration)
  f.DurationVar(p, name, value, usage)
  return p
}

// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
func (f *FlagSet) Var(value FlagValue, name string, usage string) {
  // Remember the default value as a string; it won't change.
  flag := &Flag{name, usage, value, value.String()}
  _, alreadythere := f.formal[name]
  if alreadythere {
    var msg string
    if f.name == "" {
      msg = fmt.Sprintf("flag redefined: %s", name)
    } else {
      msg = fmt.Sprintf("%s flag redefined: %s", f.name, name)
    }
    fmt.Fprintln(f.out(), msg)
    panic(msg) // Happens only if flags are declared with identical names
  }
  if f.formal == nil {
    f.formal = make(map[string]*Flag)
  }
  f.formal[name] = flag
}

// failf prints to standard error a formatted error and usage message and
// returns the error.
func (f *FlagSet) failf(format string, a ...interface{}) error {
  err := fmt.Errorf(format, a...)
  fmt.Fprintln(f.out(), err)
  f.usage()
  return err
}

// usage calls the Usage method for the flag set
func (f *FlagSet) usage() {
  if f.Usage == nil {
    defaultUsage(f)
  } else {
    f.Usage()
  }
}

// parseOne parses one flag. It reports whether a flag was seen.
func (f *FlagSet) parseOne() (bool, error) {
  if len(f.args) == 0 {
    return false, nil
  }
  s := f.args[0]
  if len(s) == 0 || s[0] != '-' || len(s) == 1 {
    return false, nil
  }
  num_minuses := 1
  if s[1] == '-' {
    num_minuses++
    if len(s) == 2 { // "--" terminates the flags
      f.args = f.args[1:]
      return false, nil
    }
  }
  name := s[num_minuses:]
  if len(name) == 0 || name[0] == '-' || name[0] == '=' {
    return false, f.failf("bad flag syntax: %s", s)
  }

  // it's a flag. does it have an argument?
  f.args = f.args[1:]
  has_value := false
  value := ""
  for i := 1; i < len(name); i++ { // equals cannot be first
    if name[i] == '=' {
      value = name[i+1:]
      has_value = true
      name = name[0:i]
      break
    }
  }
  m := f.formal
  flag, alreadythere := m[name] // BUG
  if !alreadythere {
    if name == "help" || name == "h" { // special case for nice help message.
      f.usage()
      return false, ErrHelp
    }
    return false, f.failf("flag provided but not defined: -%s", name)
  }
  if fv, ok := flag.Value.(boolFlag); ok && fv.IsBoolFlag() { // special case: doesn't need an arg
    if has_value {
      if err := fv.Set(value); err != nil {
        return false, f.failf("invalid boolean value %q for -%s: %v", value, name, err)
      }
    } else {
      fv.Set("true")
    }
  } else {
    // It must have a value, which might be the next argument.
    if !has_value && len(f.args) > 0 {
      // value is the next arg
      has_value = true
      value, f.args = f.args[0], f.args[1:]
    }
    if !has_value {
      return false, f.failf("flag needs an argument: -%s", name)
    }
    if err := flag.Value.Set(value); err != nil {
      return false, f.failf("invalid value %q for flag -%s: %v", value, name, err)
    }
  }
  if f.actual == nil {
    f.actual = make(map[string]*Flag)
  }
  f.actual[name] = flag
  return true, nil
}

// Parse parses flag definitions from the argument list, which should not
// include the command name.  Must be called after all flags in the FlagSet
// are defined and before flags are accessed by the program.
// The return value will be ErrHelp if -help was set but not defined.
func (f *FlagSet) Parse(arguments []string) error {
  f.parsed = true
  f.args = arguments
  for {
    seen, err := f.parseOne()
    if seen {
      continue
    }
    if err == nil {
      break
    }
    switch f.errorHandling {
    case ContinueOnError:
      return err
    case ExitOnError:
      os.Exit(2)
    case PanicOnError:
      panic(err)
    }
  }
  return nil
}

// Parsed reports whether f.Parse has been called.
func (f *FlagSet) Parsed() bool {
  return f.parsed
}

// Init sets the name and error handling property for a flag set.
// By default, the zero FlagSet uses an empty name and the
// ContinueOnError error handling policy.
func (f *FlagSet) Init(name string, errorHandling ErrorHandling) {
  f.name = name
  f.errorHandling = errorHandling
}
