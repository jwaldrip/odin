package cli

import "github.com/jwaldrip/odin/cli/values"

// A Param represents the state of a flag.
type Param struct {
	Name     string       // name as it appears on command line
	Value    values.Value // value as set
	DefValue string       // default value (as text); for usage message
}
