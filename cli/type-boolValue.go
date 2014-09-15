package cli

import "fmt"
import "strconv"

type boolValue bool

func newBoolValue(val bool, p *bool) *boolValue {
  *p = val
  return (*boolValue)(p)
}

func (b *boolValue) Set(s string) error {
  v, err := strconv.ParseBool(s)
  *b = boolValue(v)
  return err
}

func (b *boolValue) Get() interface{} { return bool(*b) }

func (b *boolValue) String() string { return fmt.Sprintf("%v", *b) }

func (b *boolValue) IsBoolFlag() bool { return true }
