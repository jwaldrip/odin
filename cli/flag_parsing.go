package cli

import (
	"strings"

	"github.com/jwaldrip/odin/cli/values"
)

// defineHelp defines a help function and alias if they are not present
func (cmd *CLI) defineHelp() {
	if _, ok := cmd.flags["help"]; !ok {
		cmd.DefineBoolFlag("help", false, "show help and exit")
		cmd.flagHelp = cmd.flags["help"]
		if _, ok := cmd.aliases['h']; !ok {
			cmd.AliasFlag('h', "help")
		}
	}
}

// defineVersion defines a version if one has been set
func (cmd *CLI) defineVersion() {
	if _, ok := cmd.flags["version"]; !ok && len(cmd.Version()) > 0 {
		cmd.DefineBoolFlag("version", false, "show version and exit")
		cmd.flagVersion = cmd.flags["version"]
		if _, ok := cmd.aliases['v']; !ok {
			cmd.AliasFlag('v', "version")
		}
	}
}

// flagFromArg determines the flags from an argument
func (cmd *CLI) flagFromArg(arg string) (bool, []*Flag) {
	var flags []*Flag

	// Do nothing if flags terminated
	if cmd.flagsTerminated {
		return false, flags
	}
	if arg[len(arg)-1] == '=' {
		cmd.errf("invalid flag format")
	}
	arg = strings.Split(arg, "=")[0]

	// Determine if we need to terminate flags
	isFlag := arg[0] == '-'
	areAliases := isFlag && arg[1] != '-'
	isTerminator := !areAliases && len(arg) == 2

	if isTerminator {
		cmd.flagsTerminated = true
		return false, flags
	}

	// Determine if name or alias
	if areAliases {
		aliases := arg[1:]
		for _, c := range aliases {
			flag, ok := cmd.aliases[c]
			if !ok {
				cmd.errf("invalid alias: %v", string(c))
			}
			flags = append(flags, flag)
		}
	} else {
		name := arg[2:]
		flag, ok := cmd.flags[name]
		if !ok {
			cmd.errf("invalid flag")
		}
		flags = append(flags, flag)
	}
	return areAliases, flags
}

func (cmd *CLI) initFlagValues() {
	if cmd.flagValues == nil {
		cmd.flagValues = make(map[*Flag]values.Value)
	}
}

// parse flag and param definitions from the argument list, returns any left
// over arguments after flags have been parsed.
func (cmd *CLI) parseFlags(args []string) []string {
	cmd.defineHelp()
	cmd.defineVersion()
	cmd.initFlagValues()

	// Set all the flags to defaults before setting
	cmd.setFlagDefaults()

	// copy propogating flags
	cmd.copyPropogatingFlags()

	// Set inherited values
	cmd.setFlagValuesFromParent()

	var paramIndex int
	var nonFlags []string

	// Set each flag by its set value
	for {
		// Break if no arguments remain
		if len(args) == 0 {
			cmd.flagsTerminated = true
			break
		}
		arg := args[0]

		if arg[0] != '-' { // Parse Param

			if paramIndex == len(cmd.params) && len(cmd.subCommands) > 0 {
				break
			}
			nonFlags = append(nonFlags, arg)
			paramIndex++
			args = args[1:]

		} else { // Parse Flags

			isAlias, flags := cmd.flagFromArg(arg)

			// Break if the flags have been terminated
			if cmd.flagsTerminated {
				// Remove the flag terminator if it exists
				if arg == "--" {
					args = args[1:]
				}
				break
			}

			// Set flag values
			if isAlias {
				args = cmd.setAliasValues(flags, args)
			} else {
				args = cmd.setFlagValue(flags[0], args)
			}
		}
	}

	// reposition the Args so that they may still be parsed
	args = append(nonFlags, args...)

	// return the remaining unused args
	return args
}

// setAliasValues sets the values of flags from thier aliases
func (cmd *CLI) setAliasValues(flags []*Flag, args []string) []string {
	for i, flag := range flags {
		isLastFlag := i == len(flags)-1
		if isLastFlag {
			args = cmd.setFlagValue(flag, args)
		} else {
			cmd.setFlagValue(flag, []string{})
		}
	}
	return args
}

// setFlagDefaults sets the default values of all flags
func (cmd *CLI) setFlagDefaults() {
	for _, flag := range cmd.flags {
		cmd.setFlag(flag, flag.DefValue)
	}
}

// setFlag sets the value of the named flag.
func (cmd *CLI) setFlag(flag *Flag, value string) error {
	_ = cmd.flags[flag.Name] // Verify the flag is a flag for f set
	err := flag.value.Set(value)
	if err != nil {
		return err
	}
	cmd.flagValues[flag] = flag.value
	return nil
}

// setFlagValue sets the value of a given flag
func (cmd *CLI) setFlagValue(flag *Flag, args []string) []string {
	if flag == nil {
		flag = noFlag // Fix for when we continue on error
	}
	splitArgs := []string{}
	hasSetValue := false
	hasPosValue := false
	isBoolFlag := false
	if fv, ok := flag.value.(boolFlag); ok && fv.IsBoolValue() {
		isBoolFlag = true
	}

	if len(args) > 0 {
		splitArgs = strings.Split(args[0], "=")
		hasSetValue = len(splitArgs) >= 2
		hasPosValue = len(args) >= 2 && args[1][0] != '-'
	}

	cutLen := 0

	var err error

	if hasSetValue {
		err = cmd.setFlag(flag, splitArgs[1])
		cutLen = 1
	} else if isBoolFlag {
		cmd.setFlag(flag, "true")
		cutLen = 1
	} else if hasPosValue {
		err = cmd.setFlag(flag, args[1])
		cutLen = 2
	} else {
		cmd.errf("flag \"--%v\" is missing a value", flag.Name)
	}

	cmd.handleErr(err)

	if len(args) > cutLen {
		return args[cutLen:]
	}
	return []string{}
}
