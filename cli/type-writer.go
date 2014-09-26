package cli

import "os"
import "io"
import "fmt"

type writer struct {
	ErrorHandling ErrorHandling
	usage         func()
	errOutput     io.Writer
	stdOutput     io.Writer
}

// ErrOutput is the error output for the command
func (cmd *writer) ErrOutput() io.Writer {
	if cmd.errOutput == nil {
		cmd.errOutput = os.Stderr
	}
	return cmd.errOutput
}

// Usage calls the Usage method for the flag set
func (cmd *writer) Usage() {
	if cmd.usage != nil {
		cmd.usage()
	}
}

func (cmd *writer) Mute() {
	var err error
	cmd.errOutput, err = os.Open(os.DevNull)
	cmd.stdOutput, err = os.Open(os.DevNull)
	exitIfError(err)
}

// ErrOutput is the error output for the command
func (cmd *writer) StdOutput() io.Writer {
	if cmd.stdOutput == nil {
		cmd.stdOutput = os.Stdout
	}
	return cmd.stdOutput
}

// failf prints to standard error a formatted error and usage message and
// returns the error.
func (cmd *writer) failf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)
	fmt.Fprintln(cmd.ErrOutput(), err)
	fmt.Fprintln(cmd.ErrOutput(), "")
	cmd.Usage()
	return err
}

func (cmd *writer) errf(format string, a ...interface{}) {
	cmd.handleErr(cmd.failf(format, a...))
}

func (cmd *writer) panicf(format string, a ...interface{}) {
	panic(cmd.failf(format, a...))
}

func (cmd *writer) handleErr(err error) {
	if err != nil {
		switch cmd.ErrorHandling {
		case ExitOnError:
			os.Exit(2)
		case PanicOnError:
			panic(err)
		}
	}
}
