package cli

import "fmt"
import "strconv"

// -- uint64 Value
type uint64Value uint64

func newUint64Value(val uint64, p *uint64) *uint64Value {
	*p = val
	return (*uint64Value)(p)
}

func (u *uint64Value) Get() interface{} {
	return uint64(*u)
}

func (u *uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*u = uint64Value(v)
	return err
}

func (u *uint64Value) String() string {
	return fmt.Sprintf("%v", *u)
}
