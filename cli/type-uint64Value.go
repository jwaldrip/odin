package cli

import "fmt"
import "strconv"

// -- uint64 Value
type uint64Value uint64

func newUint64Value(val uint64, p *uint64) *uint64Value {
  *p = val
  return (*uint64Value)(p)
}

func (i *uint64Value) Set(s string) error {
  v, err := strconv.ParseUint(s, 0, 64)
  *i = uint64Value(v)
  return err
}

func (i *uint64Value) Get() interface{} { return uint64(*i) }

func (i *uint64Value) String() string { return fmt.Sprintf("%v", *i) }
