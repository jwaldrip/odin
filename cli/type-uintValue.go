package cli

import "fmt"
import "strconv"

// -- uint Value
type uintValue uint

func newUintValue(val uint, p *uint) *uintValue {
	*p = val
	return (*uintValue)(p)
}

func (this *uintValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*this = uintValue(v)
	return err
}

func (this *uintValue) Get() interface{} {
	return uint(*this)
}

func (this *uintValue) String() string {
	return fmt.Sprintf("%v", *this)
}
