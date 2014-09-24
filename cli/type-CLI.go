package cli

import "os"
import "fmt"
import "bytes"
import "strings"

// CLI represents a set of defined flags.  The zero value of a FlagSet
// has no name and has ContinueOnError error handling.
type CLI struct {
	flagable
	paramable
	subCommandable
	*writer

	fn          commandFn
	name        string
	description string
	parsed      bool
}

// NewCLI returns a new cli with the specified name and
// error handling property.
func NewCLI(fn commandFn, paramNames ...string) *CLI {
	cli := new(CLI)
	cli.init(os.Args[0], fn, paramNames...)
	cli.SetVersion("v0.0.1")
	return cli
}

// DefineSubCommand return a SubCommand and adds the current CLI as the parent
func (c *CLI) DefineSubCommand(name string, desc string, fn commandFn, paramNames ...string) *SubCommand {
	cmd := c.subCommandable.DefineSubCommand(name, desc, fn, paramNames...)
	cmd.parent = c
	return cmd
}

// Description returns the command description
func (c *CLI) Description() string {
	return c.description
}

// Name returns the command name
func (c *CLI) Name() string {
	var name string
	if c.parent != nil {
		name = strings.Join([]string{c.parent.Name(), c.name}, " ")
	} else {
		name = c.name
	}
	return name
}

// Parsed reports whether f.Parse has been called.
func (c *CLI) Parsed() bool {
	c.parsed = c.flagable.Parsed() && c.paramable.Parsed() && c.subCommandable.Parsed()
	return c.parsed
}

// SetDescription sets the command description
func (c *CLI) SetDescription(desc string) {
	c.description = desc
}

// Start starts the command with args, arg[0] is ignored
func (c *CLI) Start(args ...string) {
	if args == nil {
		args = os.Args
	}

	if len(args) > 1 {
		args = args[1:]
		args = c.parseFlags(args)
		args = c.parseParams(args)

		// Show a version
		if len(c.Version()) > 0 && c.Flag("version").Get() == true {
			fmt.Println(c.Name(), c.Version())
			return
		}

		// Show Help
		if c.Flag("help").Get() == true {
			c.Usage()
			return
		}

		// run subcommands
		ransubcommand := c.parseSubCommands(args)

		if ransubcommand {
			return
		}
	}

	// Run the function
	c.fn(c)
}

// UsageString returns the command usage as a string
func (c *CLI) UsageString() string {
	hasSubCommands := len(c.subCommands) > 0
	hasParams := len(c.Params()) > 0
	hasDescription := len(c.description) > 0

	// Start the Buffer
	var buff bytes.Buffer

	buff.WriteString("Usage:\n")
	buff.WriteString(fmt.Sprintf("  %s [options...]", c.Name()))

	// Write Param Syntax
	if hasParams {
		buff.WriteString(fmt.Sprintf(" %s", c.paramable.UsageString()))
	}

	// Write Sub Command Syntax
	if hasSubCommands {
		buff.WriteString(" <command> [arg...]")
	}

	if hasDescription {
		buff.WriteString(fmt.Sprintf("\n\n%s", c.Description()))
	}

	// Write Flags Syntax
	buff.WriteString("\n\nOptions:\n")
	buff.WriteString(c.flagable.UsageString())

	// Write Sub Command List
	if hasSubCommands {
		buff.WriteString("\n\nCommands:\n")
		buff.WriteString(c.subCommandable.UsageString())
	}

	// Return buffer as string
	return buff.String()
}

func (c *CLI) init(name string, fn commandFn, paramNames ...string) {
	writer := &writer{ErrorHandling: ExitOnError}
	c.writer = writer
	c.flagable = flagable{writer: writer}
	c.paramable = paramable{writer: writer}
	c.subCommandable = subCommandable{writer: writer}
	c.name = name
	c.fn = fn
	c.setParams(paramNames...)
	c.usage = func() { fmt.Println(c.UsageString()) }
}
