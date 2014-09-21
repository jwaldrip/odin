package cli

type paramable struct {
  writer
  paramNames      []*Param
  paramValues     map[string]Value
  paramsParsed    bool
}

// NArg is the number of arguments remaining after flags have been processed.
func (this *paramable) ParamCount() int {
  return len(this.paramValues)
}

// Args returns the non-flag arguments.
func (this *paramable) Params() map[string]Value {
  return this.paramValues
}

// Arg returns the i'th argument.  Arg(0) is the first remaining argument
// after flags have been processed.
func (this *paramable) Param(key string) Value {
  value, ok := this.paramValues[key]
  if ok {
    return value
  } else {
    var emptyString stringValue
    emptyString = ""
    return &emptyString
  }
}

// Set Param names from strings
func (this *paramable) setParams(names ...string) {
  var paramNames []*Param
  for i := 0 ; i < len(names) ; i++ {
    name  := names[i]
    param := &Param{Name: name}
    paramNames = append(paramNames, param)
  }
  this.paramNames = paramNames
}

func (this *paramable) parseParams(args []string) ([]string, error) {
  var err error
  return args, err
}
