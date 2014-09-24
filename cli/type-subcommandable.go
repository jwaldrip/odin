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

func (this *subCommandable) DefineSubCommand(name string, desc string, fn commandFn, paramNames ...string) *SubCommand {
	if this.subCommands == nil {
		this.subCommands = make(map[string]*SubCommand)
	}
	subcommand := newSubCommand(name, desc, fn, paramNames...)
	this.subCommands[name] = subcommand
	return subcommand
}

func (this *subCommandable) Parsed() bool {
	return this.parsed
}

func (this *subCommandable) UsageString() string {
	var maxBufferLen int
	for _, cmd := range this.subCommands {
		buffLen := len(cmd.Name())
		if buffLen > maxBufferLen {
			maxBufferLen = buffLen
		}
	}

	var outputLines []string
	for _, cmd := range this.subCommands {
		var whitespace bytes.Buffer
		for {
			buffLen := len(cmd.Name()) + len(whitespace.String())
			if buffLen > maxBufferLen+5 {
				break
			}
			whitespace.WriteString(" ")
		}
		outputLines = append(outputLines, fmt.Sprintf("  %s %s %s", cmd.Name(), whitespace.String(), cmd.Description()))
	}

	return strings.Join(outputLines, "\n")
}

func (this *subCommandable) parseSubCommands(args []string) bool {
	this.parsed = true
	name := args[0]
	cmd, ok := this.subCommands[name]
	if !ok {
		this.errf("invalid command: %s", name)
	}
	cmd.Start(args...)

	return len(this.subCommands) > 0
}
