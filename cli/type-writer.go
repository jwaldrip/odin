package cli

import "io"
import "os"
import "fmt"

type writer struct {
  output io.Writer
}

func (f *writer) out() io.Writer {
  if f.output == nil {
    return os.Stderr
  }
  return f.output
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, os.Stderr is used.
func (f *writer) SetOutput(output io.Writer) {
  f.output = output
}

// failf prints to standard error a formatted error and usage message and
// returns the error.
func (f *writer) failf(format string, a ...interface{}) error {
  err := fmt.Errorf(format, a...)
  fmt.Fprintln(f.out(), err)
  f.usage()
  return err
}

func (f *writer) usage(){
  panic("not implemented")
}
