package cli

import "bytes"
import "fmt"
import "strings"

type subCommandable struct {
	*writer

	parent      Command
	subCommands map[string]*SubCommand
	parsed      bool
}

func (cmd *subCommandable) DefineSubCommand(name string, desc string, fn commandFn, paramNames ...string) *SubCommand {
	if cmd.subCommands == nil {
		cmd.subCommands = make(map[string]*SubCommand)
	}
	subcommand := newSubCommand(name, desc, fn, paramNames...)
	cmd.subCommands[name] = subcommand
	return subcommand
}

func (cmd *subCommandable) Parent() Command {
	return cmd.parent
}

func (cmd *subCommandable) Parsed() bool {
	return cmd.parsed
}

func (cmd *subCommandable) UsageString() string {
	var maxBufferLen int
	for _, cmd := range cmd.subCommands {
		buffLen := len(cmd.name)
		if buffLen > maxBufferLen {
			maxBufferLen = buffLen
		}
	}

	var outputLines []string
	for _, cmd := range cmd.subCommands {
		var whitespace bytes.Buffer
		for {
			buffLen := len(cmd.name) + len(whitespace.String())
			if buffLen == maxBufferLen+5 {
				break
			}
			whitespace.WriteString(" ")
		}
		outputLines = append(outputLines, fmt.Sprintf("  %s%s%s", cmd.name, whitespace.String(), cmd.Description()))
	}

	return strings.Join(outputLines, "\n")
}

func (cmd *subCommandable) parse(args []string) ([]string, bool) {
	if len(args) == 0 || len(cmd.subCommands) == 0 {
		return args, false
	}
	cmd.parsed = true
	name := args[0]
	subcmd, ok := cmd.subCommands[name]
	if !ok {
		subcmd.errf("invalid command: %s", name)
	}
	subcmd.Start(args...)

	return []string{}, true
}
