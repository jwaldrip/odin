package values

import "fmt"
import "strconv"

// Uint64 is a uint64 value for the Value interface
type Uint64 uint64

// NewUint64 returns a new uint64 value
func NewUint64(val uint64, p *uint64) *Uint64 {
	*p = val
	return (*Uint64)(p)
}

// Get returns the value interface
func (u *Uint64) Get() interface{} {
	return uint64(*u)
}

// Set sets the value from a string
func (u *Uint64) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*u = Uint64(v)
	return err
}

// String returns the value as a string
func (u *Uint64) String() string {
	return fmt.Sprintf("%v", *u)
}
