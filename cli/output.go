package cli

import (
	"fmt"
	"os"
)
import "io"

// ErrOutput is an alias for StdErr
func (cmd *CLI) ErrOutput() io.Writer {
	return cmd.StdErr()
}

// ErrPrint does a fmt.Print to the std err of the CLI
func (cmd *CLI) ErrPrint(a ...interface{}) {
	fmt.Fprint(cmd.StdErr(), a...)
}

// ErrPrintf does a fmt.Printf to the std err of the CLI
func (cmd *CLI) ErrPrintf(format string, a ...interface{}) {
	fmt.Fprintf(cmd.StdErr(), format, a...)
}

// ErrPrintln does a fmt.Println to the std err of the CLI
func (cmd *CLI) ErrPrintln(a ...interface{}) {
	fmt.Fprintln(cmd.StdErr(), a...)
}

// Mute mutes the output
func (cmd *CLI) Mute() {
	var err error
	cmd.errOutput, err = os.Open(os.DevNull)
	cmd.stdOutput, err = os.Open(os.DevNull)
	exitIfError(err)
}

// Print does a fmt.Print to the std out of the CLI
func (cmd *CLI) Print(a ...interface{}) {
	fmt.Fprint(cmd.StdOut(), a...)
}

// Printf does a fmt.Printf to the std out of the CLI
func (cmd *CLI) Printf(format string, a ...interface{}) {
	fmt.Fprintf(cmd.StdOut(), format, a...)
}

// Println does a fmt.Println to the std out of the CLI
func (cmd *CLI) Println(a ...interface{}) {
	fmt.Fprintln(cmd.StdOut(), a...)
}

// SetErrOutput is an alias for SetStdErr
func (cmd *CLI) SetErrOutput(writer io.Writer) {
	cmd.SetStdErr(writer)
}

// SetStdErr sets the error output for the command
func (cmd *CLI) SetStdErr(writer io.Writer) {
	cmd.errOutput = writer
}

// SetStdOut sets the standard output for the command
func (cmd *CLI) SetStdOut(writer io.Writer) {
	cmd.stdOutput = writer
}

// SetStdOutput is an alias for SetStdOut
func (cmd *CLI) SetStdOutput(writer io.Writer) {
	cmd.SetStdOut(writer)
}

// StdOut is the standard output for the command
func (cmd *CLI) StdOut() io.Writer {
	if cmd.stdOutput == nil {
		cmd.stdOutput = os.Stdout
	}
	return cmd.stdOutput
}

// StdOutput is an alias for StdOut
func (cmd *CLI) StdOutput() io.Writer {
	return cmd.StdOut()
}

// StdErr is the error output for the command
func (cmd *CLI) StdErr() io.Writer {
	if cmd.errOutput == nil {
		cmd.errOutput = os.Stderr
	}
	return cmd.errOutput
}
