package cli

import (
	"fmt"

	"github.com/jwaldrip/odin/cli/values"
)

// Param returns named param
func (cmd *CLI) Param(name string) values.Value {
	value, ok := cmd.Params()[name]
	if !ok {
		panic(fmt.Sprintf("param not defined %v", name))
	}
	return value
}

// Params returns the non-flag arguments.
func (cmd *CLI) Params() values.Map {
	params := make(values.Map)
	for param, value := range cmd.paramValues {
		params[param.Name] = value
	}
	return params
}
