package cli

import "fmt"
import "strconv"

// -- int64 Value
type int64Value int64

func newInt64Value(val int64, p *int64) *int64Value {
  *p = val
  return (*int64Value)(p)
}

func (i *int64Value) Set(s string) error {
  v, err := strconv.ParseInt(s, 0, 64)
  *i = int64Value(v)
  return err
}

func (i *int64Value) Get() interface{} { return int64(*i) }

func (i *int64Value) String() string { return fmt.Sprintf("%v", *i) }
