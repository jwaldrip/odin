package cli

// ErrorHandling defines how to handle flag parsing errors.
type ErrorHandling int

const (
  ContinueOnError ErrorHandling = iota
  ExitOnError
  PanicOnError
)
