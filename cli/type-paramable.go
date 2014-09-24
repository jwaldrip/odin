package cli

import "fmt"
import "strings"

type paramable struct {
	*writer
	params       paramsList
	paramValues  map[string]Value
	paramsParsed bool
	parsed       bool
}

// Arg returns the i'th argument.  Arg(0) is the first remaining argument
// after flags have been processed.
func (p *paramable) Param(key string) Value {
	value, ok := p.paramValues[key]
	if ok {
		return value
	}
	var emptyString stringValue
	emptyString = ""
	return &emptyString
}

// Args returns the non-flag arguments.
func (p *paramable) Params() map[string]Value {
	return p.paramValues
}

// NArg is the number of arguments remaining after flags have been processed.
func (p *paramable) ParamCount() int {
	return len(p.paramValues)
}

// Parsed returns if the flags have been parsed
func (p *paramable) Parsed() bool {
	return p.parsed
}

// UsageString returns the params usage as a string
func (p *paramable) UsageString() string {
	var formattednames []string
	for i := 0; i < len(p.params); i++ {
		param := p.params[i]
		formattednames = append(formattednames, fmt.Sprintf("<%s>", param.Name))
	}
	return strings.Join(formattednames, " ")
}

// Set Param names from strings
func (p *paramable) setParams(names ...string) {
	var params []*Param
	for i := 0; i < len(names); i++ {
		name := names[i]
		param := &Param{Name: name}
		params = append(params, param)
	}
	p.params = params
}

func (p *paramable) parse(args []string) []string {
	var seenParams paramsList

	if len(p.params) == 0 {
		return args
	}
	i := 0
	for i < len(args) && i < len(p.params) {
		param := p.params[i]
		seenParams = append(seenParams, param)
		str := ""
		if p.paramValues == nil {
			p.paramValues = make(map[string]Value)
		}
		p.paramValues[param.Name] = newStringValue(args[i], &str)
		i++
	}
	missingParams := p.params.Compare(seenParams)
	if len(missingParams) > 0 {
		var msg string
		if len(missingParams) == 1 {
			msg = "missing param"
		} else {
			msg = "missing params"
		}
		p.errf("%s: %s", msg, strings.Join(missingParams.Names(), ", "))
	}

	return args[i:]
}
