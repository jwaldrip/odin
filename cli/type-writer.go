package cli

import "os"
import "fmt"

type writer struct {
  ErrorHandling ErrorHandling
  usage func()
}

// failf prints to standard error a formatted error and usage message and
// returns the error.
func (this *writer) failf(format string, a ...interface{}) error {
  err := fmt.Errorf(format, a...)
  fmt.Fprintln(os.Stderr, err)
  this.Usage()
  return err
}

func (this *writer) errf(format string, a ...interface{}){
  this.handleErr(this.failf(format, a...))
}

func (this *writer) panicf(format string, a ...interface{}){
  panic(this.failf(format, a...))
}

func (this *writer) handleErr(err error){
  if err != nil {
    switch this.ErrorHandling {
    case ExitOnError:
      os.Exit(2)
    case PanicOnError:
      panic(err)
    }
  }
}

// usage calls the Usage method for the flag set
func (this *writer) Usage() {
  if this.usage != nil {
    this.usage()
  }
}
