package cli

import "io"
import "os"
import "fmt"

type writer struct {
  output io.Writer
}

func (this *writer) out() io.Writer {
  if this.output == nil {
    return os.Stderr
  }
  return this.output
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, os.Stderr is used.
func (this *writer) SetOutput(output io.Writer) {
  this.output = output
}

// failf prints to standard error a formatted error and usage message and
// returns the error.
func (this *writer) failf(format string, a ...interface{}) error {
  err := fmt.Errorf(format, a...)
  fmt.Fprintln(this.out(), err)
  this.usage()
  return err
}

func (this *writer) usage(){
  panic("not implemented")
}
