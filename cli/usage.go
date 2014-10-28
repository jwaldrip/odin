package cli

import (
	"bytes"
	"fmt"
	"strings"
)

var sep = "     : "
var columnWidth []int

// DefaultUsage is the Default Usage used by odin
func (cmd *CLI) DefaultUsage() {
	fmt.Fprintln(cmd.StdOutput(), cmd.UsageString())
}

// FlagsUsageString returns the flags usage as a string
func (cmd *CLI) FlagsUsageString(title string) string {
	flagStrings := make(map[*Flag][]string)

	// init the usage strings slice for each flag
	for _, flag := range cmd.flags {
		flagStrings[flag] = []string{}
	}

	// alias keys
	for alias, flag := range cmd.aliases {
		flagStrings[flag] = append(flagStrings[flag], "-"+string(alias))
	}

	// flag keys and values
	for _, flag := range cmd.flags {
		if _, boolflag := flag.value.(boolFlag); boolflag {
			flagStrings[flag] = append(flagStrings[flag], "--"+flag.Name)
		} else {
			flagStrings[flag] = append(flagStrings[flag], "--"+flag.Name+"="+flag.DefValue)
		}
	}

	// build the table
	tbl := NewSharedShellTable(&sep, &columnWidth)
	for flag, usages := range flagStrings {
		row := tbl.Row()
		row.Column(" ", strings.Join(usages, ", "))
		row.Column(flag.Usage)
	}

	var usage string

	if cmd.parent != nil && cmd.parent.(*CLI).hasFlags() {
		parent := cmd.parent.(*CLI)
		parentUsage := parent.FlagsUsageString(fmt.Sprintf("Options for `%s`", parent.Name()))
		usage = fmt.Sprintf("%s%s", tbl.String(), parentUsage)
	} else {
		usage = tbl.String()
	}

	return fmt.Sprintf("\n\n%s:\n%s", title, usage)
}

// ParamsUsageString returns the params usage as a string
func (cmd *CLI) ParamsUsageString() string {
	var formattednames []string
	for i := 0; i < len(cmd.params); i++ {
		param := cmd.params[i]
		formattednames = append(formattednames, fmt.Sprintf("<%s>", param.Name))
	}
	return strings.Join(formattednames, " ")
}

// SubCommandsUsageString is the usage string for sub commands
func (cmd *CLI) SubCommandsUsageString(title string) string {
	tbl := NewSharedShellTable(&sep, &columnWidth)
	for _, subcmd := range cmd.subCommands {
		row := tbl.Row()
		row.Column(" ", subcmd.Name())
		row.Column(subcmd.Description())
	}
	return fmt.Sprintf("\n\n%s:\n%s", title, tbl.String())
}

// Usage calls the Usage method for the flag set
func (cmd *CLI) Usage() {
	if cmd.usage == nil {
		cmd.usage = cmd.DefaultUsage
	}
	cmd.usage()
}

// CommandUsageString returns the command and its accepted options and params
func (cmd *CLI) CommandUsageString() string {
	hasParams := len(cmd.params) > 0
	hasParent := cmd.parent != nil
	var buff bytes.Buffer

	// Write the parent string
	if hasParent {
		buff.WriteString(cmd.parent.(*CLI).CommandUsageString())
	} else {
		buff.WriteString(" ")
	}

	// Write the name with options
	buff.WriteString(fmt.Sprintf(" %s", cmd.Name()))

	if cmd.hasFlags() {
		buff.WriteString(" [options...]")
	}

	// Write Param Syntax
	if hasParams {
		buff.WriteString(fmt.Sprintf(" %s", cmd.ParamsUsageString()))
	}

	return buff.String()
}

// UsageString returns the command usage as a string
func (cmd *CLI) UsageString() string {
	hasSubCommands := len(cmd.subCommands) > 0
	hasDescription := len(cmd.description) > 0
	hasLongDescription := len(cmd.longDescription) > 0

	// Prefetch table to calculate the widths
	_ = cmd.SubCommandsUsageString("")
	_ = cmd.FlagsUsageString("")

	// Start the Buffer
	var buff bytes.Buffer

	buff.WriteString("Usage:\n")
	buff.WriteString(cmd.CommandUsageString())

	// Write Sub Command Syntax
	if hasSubCommands {
		buff.WriteString(" <command> [arg...]")
	}

	if hasLongDescription {
		buff.WriteString(fmt.Sprintf("\n\n%s", cmd.LongDescription()))
	} else if hasDescription {
		buff.WriteString(fmt.Sprintf("\n\n%s", cmd.Description()))
	}

	// Write Options Syntax
	buff.WriteString(cmd.FlagsUsageString("Options"))

	// Write Sub Command List
	if hasSubCommands {
		buff.WriteString(cmd.SubCommandsUsageString("Commands"))
	}

	// Return buffer as string
	return buff.String()
}
