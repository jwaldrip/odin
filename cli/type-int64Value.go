package cli

import "fmt"
import "strconv"

// -- int64 Value
type int64Value int64

func newInt64Value(val int64, p *int64) *int64Value {
	*p = val
	return (*int64Value)(p)
}

func (this *int64Value) Get() interface{} {
	return int64(*this)
}

func (this *int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*this = int64Value(v)
	return err
}

func (this *int64Value) String() string {
	return fmt.Sprintf("%v", *this)
}
