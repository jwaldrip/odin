package cli

import "fmt"
import "time"
import "os"

type flags struct {
  writer
  flags           map[string]*Flag
  flagValues      map[*Flag]Value
  flagsTerminated bool
  flagsParsed     bool
  ErrorHandling ErrorHandling
}

// VisitAll visits the flags in lexicographical order, calling fn for each.
// It visits all flags, even those not set.
// Really want EachFlagWithDefaults
func (f *flags) EachFlag(fn func(*Flag, Value)) {
  for name, flag := range f.flags {
    value := f.Flag(name)
    fn(flag, value)
  }
}

func (f *flags) AliasFlag(newname string, oldname string){
  flag, ok := f.flags[oldname]
  if !ok {
    panic(fmt.Sprintf("flag not defined %v", oldname))
  }
  f.flags[newname] = flag
}

// Lookup returns the Flag structure of the named flag, returning nil if none exists.
func (f *flags) Flag(name string) Value {
  flag, ok := f.flags[name]
  if !ok {
    panic(fmt.Sprintf("flag not defined %v", name))
  }
  value, ok := f.flagValues[flag]
  if ok {
    return value
  } else {
    return nil
  }
}

// Set sets the value of the named flag.
func (f *flags) SetFlag(name, value string) error {
  flag, ok := f.flags[name]
  if !ok {
    return fmt.Errorf("no such flag -%v", name)
  }
  err := flag.value.Set(value)
  if err != nil {
    return err
  }
  if f.flagValues == nil {
    f.flagValues = make(map[*Flag]Value)
  }
  f.flagValues[flag] = flag.value
  return nil
}

// NFlag returns the number of flags that have been set.
func (f *flags) FlagCount() int { return len(f.flagValues) }

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (f *flags) DefineBoolFlagVar(p *bool, name string, value bool, usage string) {
  f.DefineFlag(newBoolValue(value, p), name, usage)
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (f *flags) DefineBoolFlag(name string, value bool, usage string) *bool {
  p := new(bool)
  f.DefineBoolFlagVar(p, name, value, usage)
  return p
}

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func (f *flags) DefineIntFlagVar(p *int, name string, value int, usage string) {
  f.DefineFlag(newIntValue(value, p), name, usage)
}

// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func (f *flags) DefineIntFlag(name string, value int, usage string) *int {
  p := new(int)
  f.DefineIntFlagVar(p, name, value, usage)
  return p
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func (f *flags) DefineInt64FlagVar(p *int64, name string, value int64, usage string) {
  f.DefineFlag(newInt64Value(value, p), name, usage)
}


// Int64 defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func (f *flags) DefineInt64Flag(name string, value int64, usage string) *int64 {
  p := new(int64)
  f.DefineInt64FlagVar(p, name, value, usage)
  return p
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *flags) DefineUintFlagVar(p *uint, name string, value uint, usage string) {
  f.DefineFlag(newUintValue(value, p), name, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (f *flags) DefineUintFlag(name string, value uint, usage string) *uint {
  p := new(uint)
  f.DefineUintFlagVar(p, name, value, usage)
  return p
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (f *flags) DefineUint64FlagVar(p *uint64, name string, value uint64, usage string) {
  f.DefineFlag(newUint64Value(value, p), name, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func (f *flags) DefineUint64Flag(name string, value uint64, usage string) *uint64 {
  p := new(uint64)
  f.DefineUint64FlagVar(p, name, value, usage)
  return p
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (f *flags) DefineStringFlagVar(p *string, name string, value string, usage string) {
  f.DefineFlag(newStringValue(value, p), name, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (f *flags) DefineStringFlag(name string, value string, usage string) *string {
  p := new(string)
  f.DefineStringFlagVar(p, name, value, usage)
  return p
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func (f *flags) DefineFloat64FlagVar(p *float64, name string, value float64, usage string) {
  f.DefineFlag(newFloat64Value(value, p), name, usage)
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func (f *flags) DefineFloat64Flag(name string, value float64, usage string) *float64 {
  p := new(float64)
  f.DefineFloat64FlagVar(p, name, value, usage)
  return p
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
func (f *flags) DefineDurationFlagVar(p *time.Duration, name string, value time.Duration, usage string) {
  f.DefineFlag(newDurationValue(value, p), name, usage)
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
func (f *flags) DefineDurationFlag(name string, value time.Duration, usage string) *time.Duration {
  p := new(time.Duration)
  f.DefineDurationFlagVar(p, name, value, usage)
  return p
}

// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
func (f *flags) DefineFlag(value Value, name string, usage string) {
  // Remember the default value as a string; it won't change.
  flag := &Flag{
    Name: name,
    Usage: usage,
    value: value,
    DefValue: value.String()}
  _, alreadythere := f.flags[name]
  if alreadythere {
    msg := fmt.Sprintf("flag redefined: %s", name)
    fmt.Fprintln(f.out(), msg)
    panic(msg) // Happens only if flags are declared with identical names
  }
  if f.flags == nil {
    f.flags = make(map[string]*Flag)
  }
  f.flags[name] = flag
}


// parseOne parses one flag. It reports whether a flag was seen.
func (f *flags) parseOneFlag(args ...string) (bool, error) {
  arg := args[0]
  if f.flagsTerminated {
    return false, nil
  }
  if len(arg) == 0 || arg[0] != '-' || len(arg) == 1 {
    return false, nil
  }
  num_minuses := 1
  if arg[1] == '-' {
    num_minuses++
    if len(arg) == 2 { // "--" terminates the flags
      f.flagsTerminated = true
      return false, nil
    }
  }
  name := arg[num_minuses:]
  if len(name) == 0 || name[0] == '-' || name[0] == '=' {
    return false, f.failf("bad flag syntax: %s", arg)
  }

  // it's a flag. does it have an argument?
  // f.args = f.args[1:] // Not sure what this does
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
  m := f.flags
  flag, alreadythere := m[name] // BUG
  if !alreadythere {
    if name == "help" || name == "h" { // special case for nice help message.
      f.usage()
      return false, ErrHelp
    }
    return false, f.failf("flag provided but not defined: -%s", name)
  }
  if fv, ok := flag.value.(boolFlag); ok && fv.IsBoolFlag() { // special case: doesn't need an arg
    if has_value {
      if err := fv.Set(value); err != nil {
        return false, f.failf("invalid boolean value %q for -%s: %v", value, name, err)
      }
    } else {
      fv.Set("true")
    }
  } else {
    // It must have a value, which might be the next argument.
    if !has_value && len(args) > 0 {
      // value is the next arg
      has_value = true
      value, args = args[0], args[1:]
    }
    if !has_value {
      return false, f.failf("flag needs an argument: -%s", name)
    }
    if err := flag.value.Set(value); err != nil {
      return false, f.failf("invalid value %q for flag -%s: %v", value, name, err)
    }
  }
  if f.flagValues == nil {
    f.flagValues = make(map[*Flag]Value)
  }
  f.flagValues[flag] = flag.value
  return true, nil
}

// Parse parses flag definitions from the argument list, which should not
// include the command name.  Must be called after all flags in the FlagSet
// are defined and before flags are accessed by the program.
// The return value will be ErrHelp if -help was set but not defined.
func (f *flags) parseFlags(arguments []string) error {
  f.flagsParsed = true
  for i := 0 ; i < len(arguments) ; i++ {
    seen, err := f.parseOneFlag(arguments[i:i+2]...) // take at most 3 args
    if seen {
      continue
    }
    if err == nil {
      break
    }
    switch f.ErrorHandling {
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

func (f *flags) usage() string {
  return f.flagsUsage()
}

func (f *flags) flagsUsage() string {
  return "Flags Usage"
}
