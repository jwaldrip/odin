package cli

import "fmt"
import "strconv"

// -- uint Value
type uintValue uint

func newUintValue(val uint, p *uint) *uintValue {
	*p = val
	return (*uintValue)(p)
}

func (u *uintValue) Get() interface{} {
	return uint(*u)
}

func (u *uintValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*u = uintValue(v)
	return err
}

func (u *uintValue) String() string {
	return fmt.Sprintf("%v", *u)
}
