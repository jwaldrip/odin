package cli

import "fmt"
import "strconv"

type boolValue bool

func newBoolValue(val bool, p *bool) *boolValue {
  *p = val
  return (*boolValue)(p)
}

func (this *boolValue) Set(s string) error {
  v, err := strconv.ParseBool(s)
  *this = boolValue(v)
  return err
}

func (this *boolValue) Get() interface{} {
  return bool(*this)
}

func (this *boolValue) String() string {
  return fmt.Sprintf("%v", *this)
}

func (this *boolValue) IsBoolFlag() bool {
  return true
}
