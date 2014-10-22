package cli

import "fmt"

// parse flag and param definitions from the argument list, returns any left
// over arguments after flags have been parsed.
func (cmd *CLI) parse(args []string) []string {
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
	var paramArgs []string

	// Set each flag by its set value
	for {
		// Break if no arguments remain
		if len(args) == 0 {
			cmd.flagsTerminated = true
			break
		}
		arg := args[0]

		if arg[0] != '-' { // Parse Param

			if paramIndex == len(cmd.params) {
				break
			}
			paramArgs = append(paramArgs, arg)
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

	// Show a version
	if len(cmd.Version()) > 0 && cmd.Flag("version").Get() == true {
		fmt.Fprintln(cmd.StdOutput(), cmd.Name(), cmd.Version())
		return []string{}
	}

	// Show Help
	if cmd.Flag("help").Get() == true {
		cmd.Usage()
		return []string{}
	}

	// pass the paramArgs to parseParams()
	cmd.parseParams(paramArgs)

	// return the remaining unused args
	return args
}
