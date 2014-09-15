package cli

import "fmt"
import "strconv"

// -- uint Value
type uintValue uint

func newUintValue(val uint, p *uint) *uintValue {
  *p = val
  return (*uintValue)(p)
}

func (i *uintValue) Set(s string) error {
  v, err := strconv.ParseUint(s, 0, 64)
  *i = uintValue(v)
  return err
}

func (i *uintValue) Get() interface{} { return uint(*i) }

func (i *uintValue) String() string { return fmt.Sprintf("%v", *i) }
