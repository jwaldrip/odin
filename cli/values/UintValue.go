package values

import "fmt"
import "strconv"

// Uint is a unit value for the Value interface
type Uint uint

// NewUint returns a new uint value
func NewUint(val uint, p *uint) *Uint {
	*p = val
	return (*Uint)(p)
}

// Get returns the value interface
func (u *Uint) Get() interface{} {
	return uint(*u)
}

// Set sets the value from a string
func (u *Uint) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*u = Uint(v)
	return err
}

// String returns the value as a string
func (u *Uint) String() string {
	return fmt.Sprintf("%v", *u)
}
