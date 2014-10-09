package cli

import "os"
import "io"
import "fmt"

// ErrOutput is the error output for the command
func (cmd *CLI) ErrOutput() io.Writer {
	if cmd.errOutput == nil {
		cmd.errOutput = os.Stderr
	}
	return cmd.errOutput
}

// Mute mutes the output
func (cmd *CLI) Mute() {
	var err error
	cmd.errOutput, err = os.Open(os.DevNull)
	cmd.stdOutput, err = os.Open(os.DevNull)
	exitIfError(err)
}

// StdOutput is the error output for the command
func (cmd *CLI) StdOutput() io.Writer {
	if cmd.stdOutput == nil {
		cmd.stdOutput = os.Stdout
	}
	return cmd.stdOutput
}

// failf prints to standard error a formatted error and usage message and
// returns the error.
func (cmd *CLI) failf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)
	fmt.Fprintln(cmd.ErrOutput(), err)
	fmt.Fprintln(cmd.ErrOutput(), "")
	cmd.Usage()
	return err
}

func (cmd *CLI) errf(format string, a ...interface{}) {
	cmd.handleErr(cmd.failf(format, a...))
}

func (cmd *CLI) panicf(format string, a ...interface{}) {
	panic(cmd.failf(format, a...))
}

func (cmd *CLI) handleErr(err error) {
	if err != nil {
		switch cmd.ErrorHandling {
		case ExitOnError:
			os.Exit(2)
		case PanicOnError:
			panic(err)
		}
	}
}
