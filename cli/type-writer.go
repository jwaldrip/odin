package cli

import "os"
import "fmt"

type writer struct {
	ErrorHandling ErrorHandling
	usage         func()
}

// Usage calls the Usage method for the flag set
func (cmd *writer) Usage() {
	if cmd.usage != nil {
		cmd.usage()
	}
}

// failf prints to standard error a formatted error and usage message and
// returns the error.
func (cmd *writer) failf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)
	fmt.Fprintln(os.Stderr, err)
	fmt.Fprintln(os.Stderr, "")
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
