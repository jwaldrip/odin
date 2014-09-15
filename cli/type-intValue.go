package cli

import "fmt"
import "strconv"

// -- int Value
type intValue int

func newIntValue(val int, p *int) *intValue {
  *p = val
  return (*intValue)(p)
}

func (i *intValue) Set(s string) error {
  v, err := strconv.ParseInt(s, 0, 64)
  *i = intValue(v)
  return err
}

func (i *intValue) Get() interface{} { return int(*i) }

func (i *intValue) String() string { return fmt.Sprintf("%v", *i) }
