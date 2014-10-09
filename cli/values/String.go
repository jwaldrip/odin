package values

import "fmt"

// String is a string value for the Value interface
type String string

// NewString returns a new string value
func NewString(val string, p *string) *String {
	*p = val
	return (*String)(p)
}

// Get returns the value interface
func (s *String) Get() interface{} {
	return string(*s)
}

// Set sets the value from a string
func (s *String) Set(val string) error {
	*s = String(val)
	return nil
}

// String returns the value as a string
func (s *String) String() string {
	return fmt.Sprintf("%s", *s)
}
