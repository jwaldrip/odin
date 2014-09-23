package cli

import "fmt"
import "time"
import "strings"
import "bytes"

type flagable struct {
	*writer
	flags           map[string]*Flag
	aliases         map[rune]*Flag
	flagValues      map[*Flag]Getter
	flagsTerminated bool
	flagsParsed     bool
  Version         string
}

// VisitAll visits the flags in lexicographical order, calling fn for each.
// It visits all flags, even those not set.
// Really want EachFlagWithDefaults
func (this *flagable) EachFlag(fn func(*Flag, Value)) {
	for name, flag := range this.flags {
		value := this.Flag(name)
		fn(flag, value)
	}
}

func (this *flagable) AliasFlag(alias rune, flagname string) {
	flag, ok := this.flags[flagname]
	if !ok {
		panic(fmt.Sprintf("flag not defined %v", flagname))
	}
	if this.aliases == nil {
		this.aliases = make(map[rune]*Flag)
	}
	this.aliases[alias] = flag
}

// Lookup returns the Flag structure of the named flag, returning nil if none exists.
func (this *flagable) Flag(name string) Getter {
	flag, ok := this.flags[name]
	if !ok {
		panic(fmt.Sprintf("flag not defined %v", name))
	}
	value, ok := this.flagValues[flag]
	if ok {
		return value
	} else {
		return nil
	}
}

func (this *flagable) Flags() map[string]Getter {
	flags := make(map[string]Getter)
	for name, _ := range this.flags {
		flags[name] = this.Flag(name)
	}
	return flags
}

// Set sets the value of the named flag.
func (this *flagable) setFlag(flag *Flag, value string) error {
	// Verify the flag is a flag for this set
	flag, ok := this.flags[flag.Name]
	if !ok {
		return fmt.Errorf("no such flag -%v", flag.Name)
	}
	err := flag.value.Set(value)
	if err != nil {
		return err
	}
	if this.flagValues == nil {
		this.flagValues = make(map[*Flag]Getter)
	}
	this.flagValues[flag] = flag.value
	return nil
}

// NFlag returns the number of flags that have been set.
func (this *flagable) FlagCount() int { return len(this.flagValues) }

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (this *flagable) DefineBoolFlagVar(p *bool, name string, value bool, usage string) {
	this.DefineFlag(newBoolValue(value, p), name, usage)
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (this *flagable) DefineBoolFlag(name string, value bool, usage string) *bool {
	p := new(bool)
	this.DefineBoolFlagVar(p, name, value, usage)
	return p
}

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func (this *flagable) DefineIntFlagVar(p *int, name string, value int, usage string) {
	this.DefineFlag(newIntValue(value, p), name, usage)
}

// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func (this *flagable) DefineIntFlag(name string, value int, usage string) *int {
	p := new(int)
	this.DefineIntFlagVar(p, name, value, usage)
	return p
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func (this *flagable) DefineInt64FlagVar(p *int64, name string, value int64, usage string) {
	this.DefineFlag(newInt64Value(value, p), name, usage)
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func (this *flagable) DefineInt64Flag(name string, value int64, usage string) *int64 {
	p := new(int64)
	this.DefineInt64FlagVar(p, name, value, usage)
	return p
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (this *flagable) DefineUintFlagVar(p *uint, name string, value uint, usage string) {
	this.DefineFlag(newUintValue(value, p), name, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (this *flagable) DefineUintFlag(name string, value uint, usage string) *uint {
	p := new(uint)
	this.DefineUintFlagVar(p, name, value, usage)
	return p
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (this *flagable) DefineUint64FlagVar(p *uint64, name string, value uint64, usage string) {
	this.DefineFlag(newUint64Value(value, p), name, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func (this *flagable) DefineUint64Flag(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	this.DefineUint64FlagVar(p, name, value, usage)
	return p
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (this *flagable) DefineStringFlagVar(p *string, name string, value string, usage string) {
	this.DefineFlag(newStringValue(value, p), name, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (this *flagable) DefineStringFlag(name string, value string, usage string) *string {
	p := new(string)
	this.DefineStringFlagVar(p, name, value, usage)
	return p
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func (this *flagable) DefineFloat64FlagVar(p *float64, name string, value float64, usage string) {
	this.DefineFlag(newFloat64Value(value, p), name, usage)
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func (this *flagable) DefineFloat64Flag(name string, value float64, usage string) *float64 {
	p := new(float64)
	this.DefineFloat64FlagVar(p, name, value, usage)
	return p
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
func (this *flagable) DefineDurationFlagVar(p *time.Duration, name string, value time.Duration, usage string) {
	this.DefineFlag(newDurationValue(value, p), name, usage)
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
func (this *flagable) DefineDurationFlag(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	this.DefineDurationFlagVar(p, name, value, usage)
	return p
}

// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
func (this *flagable) DefineFlag(value Getter, name string, usage string) {
	// Remember the default value as a string; it won't change.
	flag := &Flag{
		Name:     name,
		Usage:    usage,
		value:    value,
		DefValue: value.String(),
	}
	_, alreadythere := this.flags[name]
	if alreadythere {
		this.panicf("flag redefined: %s", name)
	}
	if this.flags == nil {
		this.flags = make(map[string]*Flag)
	}
	this.flags[name] = flag
}

func (this *flagable) flagFromArg(arg string) (bool, []*Flag) {
	var flags []*Flag

	// Do nothing if flags terminated
	if this.flagsTerminated {
		return false, flags
	}
	if arg[len(arg)-1] == '=' {
		this.errf("invalid flag format")
	}
	arg = strings.Split(arg, "=")[0]

	// Determine if we need to terminate flags
	isFlag := arg[0] == '-'
	areAliases := isFlag && arg[1] != '-'
	isTerminator := !areAliases && len(arg) == 2

	if !isFlag || isTerminator {
		this.flagsTerminated = true
		return false, flags
	}

	// Determine if name or alias
	if areAliases {
		aliases := arg[1:]
		for _, c := range aliases {
			flag, ok := this.aliases[c]
			if !ok {
		    this.errf("invalid alias: %v", string(c))
			}
			flags = append(flags, flag)
		}
	} else {
		name := arg[2:]
		flag, ok := this.flags[name]
		if !ok {
			this.errf("invalid flag")
		}
		flags = append(flags, flag)
	}
	return areAliases, flags
}

func (this *flagable) setAliasValues(flags []*Flag, arg string) {
	args := strings.Split(arg, "=")
	hasvalue := len(args) > 1
	var lastflag *Flag

	// If a value is provided, set the last flag
	if hasvalue {
		lastflag = flags[len(flags)-1]
		flags = flags[:len(flags)-1]
		this.setFlag(lastflag, args[1])
	}

	for i := 0; i < len(flags); i++ {
		flag := flags[i]
		if fv, ok := flag.value.(boolFlag); ok && fv.IsBoolFlag() {
			this.setFlag(flag, "true")
		} else {
			this.panicf("flag %v missing value", flag.Name)
		}
	}
}

func (this *flagable) setFlagValue(flag *Flag, arg string) {
	args := strings.Split(arg, "=")
	hasvalue := len(args) > 1
	if hasvalue {
		value := args[1]
		this.setFlag(flag, value)
	} else {
		if fv, ok := flag.value.(boolFlag); ok && fv.IsBoolFlag() {
			this.setFlag(flag, "true")
		} else {
			this.panicf("flag %v missing value", flag.Name)
		}
	}
}

func (this *flagable) setFlagDefaults() {
	for name, flag := range this.flags {
		if this.Flag(name) == nil {
			this.setFlag(flag, flag.DefValue)
		}
	}
}

func (this *flagable) defineHelp(){
  if _, ok := this.flags["help"] ; !ok {
    this.DefineBoolFlag("help", false, "show help and exit")
    if _, ok := this.aliases['h'] ; !ok {
      this.AliasFlag('h', "help")
    }
  }
}

func (this *flagable) defineVersion(){
  if _, ok := this.flags["version"] ; !ok {
    this.DefineBoolFlag("version", false, "show version and exit")
    if _, ok := this.aliases['v'] ; !ok {
      this.AliasFlag('v', "version")
    }
  }
}

// Parse parses flag definitions from the argument list, returns any left over
// arguments after flags have been parsed.
func (this *flagable) parseFlags(args []string) []string {
  this.defineHelp()
  this.defineVersion()
	this.flagsParsed = true
	i := 0
	for i < len(args) {
		isAlias, flags := this.flagFromArg(args[i])
		if this.flagsTerminated {
			break
		}
		if isAlias {
			this.setAliasValues(flags, args[i])
		} else {
			this.setFlagValue(flags[0], args[i])
		}
		i++
	}
	// Set the remaining flags to defaults
	this.setFlagDefaults()
	// return the remaining unused args
	return args[i:]
}

func (this *flagable) FlagsUsage() string {
  var maxBufferLen int
  flagsUsages := make(map[*Flag]*bytes.Buffer)

  // init the map for each flag
  for _, flag := range this.aliases {
    flagsUsages[flag] = bytes.NewBufferString("")
  }

  // Get each flags aliases
	for r, flag := range this.aliases {
    alias := string(r)
    buffer := flagsUsages[flag]
    var err error
    if buffer.Len() == 0 {
      _, err = buffer.WriteString(fmt.Sprintf("-%s", alias))
    } else {
      _, err = buffer.WriteString(fmt.Sprintf(", -%s", alias))
    }
    exitIfError(err)
    buffLen := len(buffer.String())
    if buffLen > maxBufferLen {
      maxBufferLen = buffLen
    }
  }

  // Get each flags names
  for name, flag := range this.flags {
    buffer := flagsUsages[flag]
    var err error
    if buffer.Len() == 0 {
      _, err = buffer.WriteString(fmt.Sprintf("--%s", name))
    } else {
      _, err = buffer.WriteString(fmt.Sprintf(", --%s", name))
    }
    exitIfError(err)
    buffLen := len(buffer.String())
    if buffLen > maxBufferLen {
      maxBufferLen = buffLen
    }
  }

  // get the flag strings and append the usage info
  var outputLines []string
  for flag, buffer := range flagsUsages {
    for {
      buffLen := len(buffer.String())
      if buffLen > maxBufferLen {
        break
      }
      buffer.WriteString(" ")
    }
    outputLines = append(outputLines, fmt.Sprintf("  %s # %s", buffer.String(), flag.Usage))
  }

  return strings.Join(outputLines, "\n")
}
