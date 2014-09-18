package cli

type params struct {
  writer
  paramNames      []*Param
  paramValues     map[string]Value
}

// NArg is the number of arguments remaining after flags have been processed.
func (f *params) ParamCount() int { return len(f.paramValues) }

// Args returns the non-flag arguments.
func (f *params) Params() map[string]Value { return f.paramValues }

// Arg returns the i'th argument.  Arg(0) is the first remaining argument
// after flags have been processed.
func (f *CLI) Param(key string) Value {
  value, ok := f.paramValues[key]
  if ok {
    return value
  } else {
    var emptyString stringValue
    emptyString = ""
    return &emptyString
  }
}

// Set Param names from strings
func (f *params) setParams(names ...string) {
  var paramNames []*Param
  for i := 0 ; i < len(names) ; i++ {
    name  := names[i]
    param := &Param{Name: name}
    paramNames = append(paramNames, param)
  }
  f.paramNames = paramNames
}

func (f *params) parseParams(args []string) error {
  var err error
  return err
}
