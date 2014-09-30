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

	fn           CommandFn
	name         string
	description  string
	unparsedArgs ValueList
}

// New returns a new cli with the specified name and
// error handling property.
func New(version, desc string, fn CommandFn, paramNames ...string) *CLI {
	nameParts := strings.Split(os.Args[0], "/")
	cli := new(CLI)
	cli.init(nameParts[len(nameParts)-1], desc, fn, paramNames...)
	cli.version = version
	cli.description = desc
	return cli
}

// Alias for New
var NewCLI = New

// Args returns any remaining args that were not parsed as params
func (cmd *CLI) Args() ValueList {
	return cmd.unparsedArgs
}

// Arg takes a position of a remaining arg that was not parsed as a param
func (cmd *CLI) Arg(index int) Value {
	return cmd.Args()[index]
}

// DefineSubCommand return a SubCommand and adds the current CLI as the parent
func (cmd *CLI) DefineSubCommand(name string, desc string, fn CommandFn, paramNames ...string) *SubCommand {
	subcmd := cmd.subCommandable.DefineSubCommand(name, desc, fn, paramNames...)
	subcmd.parent = cmd
	return subcmd
}

// Description returns the command description
func (cmd *CLI) Description() string {
	return cmd.description
}

// Name returns the command name
func (cmd *CLI) Name() string {
	var name string
	if cmd.parent != nil {
		name = strings.Join([]string{cmd.parent.Name(), cmd.name}, " ")
	} else {
		name = cmd.name
	}
	return name
}

// Start starts the command with args, arg[0] is ignored
func (cmd *CLI) Start(args ...string) {
	if args == nil {
		args = os.Args
	}

	if len(args) > 1 {
		args = args[1:]
	} else {
		args = []string{}
	}

	// parse flags and args
	args = cmd.flagable.parse(args)

	// Show a version
	if len(cmd.Version()) > 0 && cmd.Flag("version").Get() == true {
		fmt.Fprintln(cmd.StdOutput(), cmd.Name(), cmd.Version())
		return
	}

	// Show Help
	if cmd.Flag("help").Get() == true {
		cmd.Usage()
		return
	}

	// Parse Params
	args = cmd.paramable.parse(args)

	var subCommandsParsed bool
	if args, subCommandsParsed = cmd.subCommandable.parse(args); subCommandsParsed {
		return
	}

	cmd.assignUnparsedArgs(args)

	// Run the function
	cmd.fn(cmd)
}

// UsageString returns the command usage as a string
func (cmd *CLI) UsageString() string {
	hasSubCommands := len(cmd.subCommands) > 0
	hasParams := len(cmd.params) > 0
	hasDescription := len(cmd.description) > 0

	// Start the Buffer
	var buff bytes.Buffer

	buff.WriteString("Usage:\n")
	buff.WriteString(fmt.Sprintf("  %s [options...]", cmd.Name()))

	// Write Param Syntax
	if hasParams {
		buff.WriteString(fmt.Sprintf(" %s", cmd.paramable.UsageString()))
	}

	// Write Sub Command Syntax
	if hasSubCommands {
		buff.WriteString(" <command> [arg...]")
	}

	if hasDescription {
		buff.WriteString(fmt.Sprintf("\n\n%s", cmd.Description()))
	}

	// Write Flags Syntax
	buff.WriteString("\n\nOptions:\n")
	buff.WriteString(cmd.flagable.UsageString())

	// Write Sub Command List
	if hasSubCommands {
		buff.WriteString("\n\nCommands:\n")
		buff.WriteString(cmd.subCommandable.UsageString())
	}

	// Return buffer as string
	return buff.String()
}

func (cmd *CLI) assignUnparsedArgs(args []string) {
	for i := 0; i < len(args); i++ {
		str := ""
		cmd.unparsedArgs = append(cmd.unparsedArgs, newStringValue(args[i], &str))
	}
}

func (cmd *CLI) init(name, desc string, fn CommandFn, paramNames ...string) {
	writer := &writer{ErrorHandling: ExitOnError}
	cmd.writer = writer
	cmd.flagable = flagable{writer: writer}
	cmd.paramable = paramable{writer: writer}
	cmd.subCommandable = subCommandable{writer: writer}
	cmd.name = name
	cmd.fn = fn
	cmd.description = desc
	cmd.DefineParams(paramNames...)
	cmd.usage = func() { fmt.Fprintln(cmd.StdOutput(), cmd.UsageString()) }
}
