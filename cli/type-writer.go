package cli

import "os"
import "fmt"

type writer struct {
	ErrorHandling ErrorHandling
	usage         func()
}

// Usage calls the Usage method for the flag set
func (w *writer) Usage() {
	if w.usage != nil {
		w.usage()
	}
}

// failf prints to standard error a formatted error and usage message and
// returns the error.
func (w *writer) failf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)
	fmt.Fprintln(os.Stderr, err)
	fmt.Fprintln(os.Stderr, "")
	w.Usage()
	return err
}

func (w *writer) errf(format string, a ...interface{}) {
	w.handleErr(w.failf(format, a...))
}

func (w *writer) panicf(format string, a ...interface{}) {
	panic(w.failf(format, a...))
}

func (w *writer) handleErr(err error) {
	if err != nil {
		switch w.ErrorHandling {
		case ExitOnError:
			os.Exit(2)
		case PanicOnError:
			panic(err)
		}
	}
}
