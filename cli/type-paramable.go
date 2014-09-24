package cli

import "fmt"
import "strings"

type paramable struct {
	*writer
	params       []*Param
	paramValues  map[string]Getter
	paramsParsed bool
	parsed       bool
}

// Arg returns the i'th argument.  Arg(0) is the first remaining argument
// after flags have been processed.
func (this *paramable) Param(key string) Getter {
	value, ok := this.paramValues[key]
	if ok {
		return value
	} else {
		var emptyString stringValue
		emptyString = ""
		return &emptyString
	}
}

// Args returns the non-flag arguments.
func (this *paramable) Params() map[string]Getter {
	return this.paramValues
}

// NArg is the number of arguments remaining after flags have been processed.
func (this *paramable) ParamCount() int {
	return len(this.paramValues)
}

// Parsed returns if the flags have been parsed
func (this *paramable) Parsed() bool {
	return this.parsed
}

// UsageString returns the params usage as a string
func (this *paramable) UsageString() string {
	var formattednames []string
	for i := 0; i < len(this.params); i++ {
		param := this.params[i]
		formattednames = append(formattednames, fmt.Sprintf("<%s>", param.Name))
	}
	return strings.Join(formattednames, " ")
}

// Set Param names from strings
func (this *paramable) setParams(names ...string) {
	var params []*Param
	for i := 0; i < len(names); i++ {
		name := names[i]
		param := &Param{Name: name}
		params = append(params, param)
	}
	this.params = params
}

func (this *paramable) parseParams(args []string) []string {
	return args
}
