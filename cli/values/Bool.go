package values

import "fmt"
import "strconv"

// Bool is a boolean value for the Value interface
type Bool bool

// NewBool returns a new bool value
func NewBool(val bool, p *bool) *Bool {
	*p = val
	return (*Bool)(p)
}

// Get returns the value interface
func (b *Bool) Get() interface{} {
	return bool(*b)
}

// IsBoolValue returns true since it is a BoolValue
func (b *Bool) IsBoolValue() bool {
	return true
}

// Set sets the value from a string
func (b *Bool) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = Bool(v)
	return err
}

// String returns the value as a string
func (b *Bool) String() string {
	return fmt.Sprintf("%v", *b)
}
