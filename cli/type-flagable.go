package cli

import "fmt"
import "time"
import "strings"
import "bytes"

type flagable struct {
	*writer
	flags           map[string]*Flag
	aliases         map[rune]*Flag
	flagValues      map[*Flag]Value
	flagsTerminated bool
	parsed          bool
	version         string
}

func (f *flagable) AliasFlag(alias rune, flagname string) {
	flag, ok := f.flags[flagname]
	if !ok {
		panic(fmt.Sprintf("flag not defined %v", flagname))
	}
	if f.aliases == nil {
		f.aliases = make(map[rune]*Flag)
	}
	f.aliases[alias] = flag
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (f *flagable) DefineBoolFlag(name string, value bool, usage string) *bool {
	p := new(bool)
	f.DefineBoolFlagVar(p, name, value, usage)
	return p
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (f *flagable) DefineBoolFlagVar(p *bool, name string, value bool, usage string) {
	f.DefineFlag(newBoolValue(value, p), name, usage)
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
func (f *flagable) DefineDurationFlag(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DefineDurationFlagVar(p, name, value, usage)
	return p
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
func (f *flagable) DefineDurationFlagVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.DefineFlag(newDurationValue(value, p), name, usage)
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func (f *flagable) DefineFloat64Flag(name string, value float64, usage string) *float64 {
	p := new(float64)
	f.DefineFloat64FlagVar(p, name, value, usage)
	return p
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func (f *flagable) DefineFloat64FlagVar(p *float64, name string, value float64, usage string) {
	f.DefineFlag(newFloat64Value(value, p), name, usage)
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func (f *flagable) DefineInt64Flag(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.DefineInt64FlagVar(p, name, value, usage)
	return p
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func (f *flagable) DefineInt64FlagVar(p *int64, name string, value int64, usage string) {
	f.DefineFlag(newInt64Value(value, p), name, usage)
}

// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func (f *flagable) DefineIntFlag(name string, value int, usage string) *int {
	p := new(int)
	f.DefineIntFlagVar(p, name, value, usage)
	return p
}

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func (f *flagable) DefineIntFlagVar(p *int, name string, value int, usage string) {
	f.DefineFlag(newIntValue(value, p), name, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (f *flagable) DefineStringFlag(name string, value string, usage string) *string {
	p := new(string)
	f.DefineStringFlagVar(p, name, value, usage)
	return p
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (f *flagable) DefineStringFlagVar(p *string, name string, value string, usage string) {
	f.DefineFlag(newStringValue(value, p), name, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func (f *flagable) DefineUint64Flag(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.DefineUint64FlagVar(p, name, value, usage)
	return p
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (f *flagable) DefineUint64FlagVar(p *uint64, name string, value uint64, usage string) {
	f.DefineFlag(newUint64Value(value, p), name, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (f *flagable) DefineUintFlag(name string, value uint, usage string) *uint {
	p := new(uint)
	f.DefineUintFlagVar(p, name, value, usage)
	return p
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *flagable) DefineUintFlagVar(p *uint, name string, value uint, usage string) {
	f.DefineFlag(newUintValue(value, p), name, usage)
}

// DefineFlag defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
func (f *flagable) DefineFlag(value Value, name string, usage string) {
	// Remember the default value as a string; it won't change.
	flag := &Flag{
		Name:     name,
		Usage:    usage,
		value:    value,
		DefValue: value.String(),
	}
	_, alreadythere := f.flags[name]
	if alreadythere {
		f.panicf("flag redefined: %s", name)
	}
	if f.flags == nil {
		f.flags = make(map[string]*Flag)
	}
	f.flags[name] = flag
}

// Flag returns the Value interface to the value of the named flag,
// returning nil if none exists.
func (f *flagable) Flag(name string) Value {
	flag, ok := f.flags[name]
	if !ok {
		panic(fmt.Sprintf("flag not defined %v", name))
	}
	value, ok := f.flagValues[flag]
	if ok {
		return value
	}
	return nil
}

func (f *flagable) Flags() map[string]Value {
	flags := make(map[string]Value)
	for name := range f.flags {
		flags[name] = f.Flag(name)
	}
	return flags
}

// FlagCount returns the number of flags that have been set.
func (f *flagable) FlagCount() int { return len(f.flagValues) }

// Parsed returns if the flags have been parsed
func (f *flagable) Parsed() bool {
	return f.parsed
}

// UsageString returns the flags usage as a string
func (f *flagable) UsageString() string {
	var maxBufferLen int
	flagsUsages := make(map[*Flag]*bytes.Buffer)

	// init the map for each flag
	for _, flag := range f.aliases {
		flagsUsages[flag] = bytes.NewBufferString("")
	}

	// Get each flags aliases
	for r, flag := range f.aliases {
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
	for name, flag := range f.flags {
		buffer := flagsUsages[flag]
		if buffer == nil {
			flagsUsages[flag] = new(bytes.Buffer)
			buffer = flagsUsages[flag]
		}
		var err error
		if buffer.Len() == 0 {
			_, err = buffer.WriteString(fmt.Sprintf("--%s", name))
		} else {
			_, err = buffer.WriteString(fmt.Sprintf(", --%s", name))
		}
		if _, ok := flag.value.(boolFlag); !ok {
			buffer.WriteString(fmt.Sprintf("=\"%s\"", flag.DefValue))
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

func (f *flagable) SetVersion(str string) {
	f.version = str
}

func (f *flagable) Version() string {
	return f.version
}

// defineHelp defines a help function and alias if they are not present
func (f *flagable) defineHelp() {
	if _, ok := f.flags["help"]; !ok {
		f.DefineBoolFlag("help", false, "show help and exit")
		if _, ok := f.aliases['h']; !ok {
			f.AliasFlag('h', "help")
		}
	}
}

// defineVersion defines a version if one has been set
func (f *flagable) defineVersion() {
	if _, ok := f.flags["version"]; !ok && len(f.Version()) > 0 {
		f.DefineBoolFlag("version", false, "show version and exit")
		if _, ok := f.aliases['v']; !ok {
			f.AliasFlag('v', "version")
		}
	}
}

// flagFromArg determines the flags from an argument
func (f *flagable) flagFromArg(arg string) (bool, []*Flag) {
	var flags []*Flag

	// Do nothing if flags terminated
	if f.flagsTerminated {
		return false, flags
	}
	if arg[len(arg)-1] == '=' {
		f.errf("invalid flag format")
	}
	arg = strings.Split(arg, "=")[0]

	// Determine if we need to terminate flags
	isFlag := arg[0] == '-'
	areAliases := isFlag && arg[1] != '-'
	isTerminator := !areAliases && len(arg) == 2

	if !isFlag || isTerminator {
		f.flagsTerminated = true
		return false, flags
	}

	// Determine if name or alias
	if areAliases {
		aliases := arg[1:]
		for _, c := range aliases {
			flag, ok := f.aliases[c]
			if !ok {
				f.errf("invalid alias: %v", string(c))
			}
			flags = append(flags, flag)
		}
	} else {
		name := arg[2:]
		flag, ok := f.flags[name]
		if !ok {
			f.errf("invalid flag")
		}
		flags = append(flags, flag)
	}
	return areAliases, flags
}

// parseFlags flag definitions from the argument list, returns any left over
// arguments after flags have been parsed.
func (f *flagable) parse(args []string) []string {
	f.defineHelp()
	f.defineVersion()
	f.parsed = true
	i := 0
	for i < len(args) {
		isAlias, flags := f.flagFromArg(args[i])
		if f.flagsTerminated {
			break
		}
		if isAlias {
			f.setAliasValues(flags, args[i])
		} else {
			f.setFlagValue(flags[0], args[i])
		}
		i++
	}
	// Set the remaining flags to defaults
	f.setFlagDefaults()
	// return the remaining unused args
	return args[i:]
}

// setAliasValues sets the values of flags from thier aliases
func (f *flagable) setAliasValues(flags []*Flag, arg string) {
	args := strings.Split(arg, "=")
	hasvalue := len(args) > 1
	var lastflag *Flag

	// If a value is provided, set the last flag
	if hasvalue {
		lastflag = flags[len(flags)-1]
		flags = flags[:len(flags)-1]
		f.setFlag(lastflag, args[1])
	}

	for i := 0; i < len(flags); i++ {
		flag := flags[i]
		if fv, ok := flag.value.(boolFlag); ok && fv.IsBoolFlag() {
			f.setFlag(flag, "true")
		} else {
			f.panicf("flag %v missing value", flag.Name)
		}
	}
}

// setFlagDefaults sets the default values of all flags
func (f *flagable) setFlagDefaults() {
	for name, flag := range f.flags {
		if f.Flag(name) == nil {
			f.setFlag(flag, flag.DefValue)
		}
	}
}

// setFlag sets the value of the named flag.
func (f *flagable) setFlag(flag *Flag, value string) error {
	// Verify the flag is a flag for f set
	flag, ok := f.flags[flag.Name]
	if !ok {
		return fmt.Errorf("no such flag -%v", flag.Name)
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

// setFlagValue sets the value of a given flag
func (f *flagable) setFlagValue(flag *Flag, arg string) {
	args := strings.Split(arg, "=")
	hasvalue := len(args) > 1
	if hasvalue {
		value := args[1]
		f.setFlag(flag, value)
	} else {
		if fv, ok := flag.value.(boolFlag); ok && fv.IsBoolFlag() {
			f.setFlag(flag, "true")
		} else {
			f.panicf("flag %v missing value", flag.Name)
		}
	}
}
