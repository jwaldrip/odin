package cli

import "fmt"
import "strconv"

// -- uint64 Value
type uint64Value uint64

func newUint64Value(val uint64, p *uint64) *uint64Value {
	*p = val
	return (*uint64Value)(p)
}

func (this *uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*this = uint64Value(v)
	return err
}

func (this *uint64Value) Get() interface{} {
	return uint64(*this)
}

func (this *uint64Value) String() string {
	return fmt.Sprintf("%v", *this)
}
