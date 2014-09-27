package cli

import "fmt"
import "strings"

type paramable struct {
	*writer
	params       paramsList
	paramValues  map[string]Value
	paramsParsed bool
}

// Param returns named param
func (cmd *paramable) Param(key string) Value {
	value, ok := cmd.paramValues[key]
	if !ok {
		cmd.errf("invalid param: %s", key)
	}
	return value
}

// Args returns the non-flag arguments.
func (cmd *paramable) Params() map[string]Value {
	return cmd.paramValues
}

// UsageString returns the params usage as a string
func (cmd *paramable) UsageString() string {
	var formattednames []string
	for i := 0; i < len(cmd.params); i++ {
		param := cmd.params[i]
		formattednames = append(formattednames, fmt.Sprintf("<%s>", param.Name))
	}
	return strings.Join(formattednames, " ")
}

// Set Param names from strings
func (cmd *paramable) DefineParams(names ...string) {
	var params []*Param
	for i := 0; i < len(names); i++ {
		name := names[i]
		param := &Param{Name: name}
		params = append(params, param)
	}
	cmd.params = params
}

func (cmd *paramable) parse(args []string) []string {
	var seenParams paramsList

	if len(cmd.params) == 0 {
		return args
	}
	i := 0
	for i < len(args) && i < len(cmd.params) {
		param := cmd.params[i]
		seenParams = append(seenParams, param)
		str := ""
		if cmd.paramValues == nil {
			cmd.paramValues = make(map[string]Value)
		}
		cmd.paramValues[param.Name] = newStringValue(args[i], &str)
		i++
	}
	missingParams := cmd.params.Compare(seenParams)
	if len(missingParams) > 0 {
		var msg string
		if len(missingParams) == 1 {
			msg = "missing param"
		} else {
			msg = "missing params"
		}
		cmd.errf("%s: %s", msg, strings.Join(missingParams.Names(), ", "))
	}

	return args[i:]
}
