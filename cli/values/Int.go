package values

import "fmt"
import "strconv"

// Int is an integer value for the Value interface
type Int int

// NewInt returns a new integer value
func NewInt(val int, p *int) *Int {
	*p = val
	return (*Int)(p)
}

// Get returns the value interface
func (i *Int) Get() interface{} {
	return int(*i)
}

// Set sets the value from a string
func (i *Int) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = Int(v)
	return err
}

// String returns the value as a string
func (i *Int) String() string {
	return fmt.Sprintf("%v", *i)
}
