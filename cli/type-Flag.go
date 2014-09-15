package cli

// A Flag represents the state of a flag.
type Flag struct {
  Name     string // name as it appears on command line
  Usage    string // help message
  Value    FlagValue  // value as set
  DefValue string // default value (as text); for usage message
}
