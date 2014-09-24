package cli

import "fmt"
import "strconv"

// -- int Value
type intValue int

func newIntValue(val int, p *int) *intValue {
	*p = val
	return (*intValue)(p)
}

func (this *intValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*this = intValue(v)
	return err
}

func (this *intValue) Get() interface{} {
	return int(*this)
}

func (this *intValue) String() string {
	return fmt.Sprintf("%v", *this)
}
