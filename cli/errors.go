package cli

import (
	"fmt"
	"os"
)

func exitIfError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
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
