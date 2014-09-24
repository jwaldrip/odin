package cli

// A Flag represents the state of a flag.
type Param struct {
	Name     string // name as it appears on command line
	Value    Value  // value as set
	DefValue string // default value (as text); for usage message
}
