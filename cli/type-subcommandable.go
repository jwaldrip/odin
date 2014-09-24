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

func (s *subCommandable) DefineSubCommand(name string, desc string, fn commandFn, paramNames ...string) *SubCommand {
	if s.subCommands == nil {
		s.subCommands = make(map[string]*SubCommand)
	}
	subcommand := newSubCommand(name, desc, fn, paramNames...)
	s.subCommands[name] = subcommand
	return subcommand
}

func (s *subCommandable) Parent() Command {
	return s.parent
}

func (s *subCommandable) Parsed() bool {
	return s.parsed
}

func (s *subCommandable) UsageString() string {
	var maxBufferLen int
	for _, cmd := range s.subCommands {
		buffLen := len(cmd.name)
		if buffLen > maxBufferLen {
			maxBufferLen = buffLen
		}
	}

	var outputLines []string
	for _, cmd := range s.subCommands {
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

func (s *subCommandable) parseSubCommands(args []string) bool {
	if len(args) == 0 {
		return false
	}
	s.parsed = true
	name := args[0]
	cmd, ok := s.subCommands[name]
	if !ok {
		s.errf("invalid command: %s", name)
	}
	cmd.Start(args...)

	return true
}
