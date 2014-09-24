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
func (p *paramable) Param(key string) Getter {
	value, ok := p.paramValues[key]
	if ok {
		return value
	}
	var emptyString stringValue
	emptyString = ""
	return &emptyString
}

// Args returns the non-flag arguments.
func (p *paramable) Params() map[string]Getter {
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
	if len(p.params) == 0 {
		return args
	}
	if len(args) < len(p.params) {
		p.errf("missing param")
	}
	i := 0
	for i < len(args) && i < len(p.params) {
		param := p.params[i]
		str := ""
		if p.paramValues == nil {
			p.paramValues = make(map[string]Getter)
		}
		p.paramValues[param.Name] = newStringValue(args[i], &str)
		i++
	}

	return args[i:]
}
