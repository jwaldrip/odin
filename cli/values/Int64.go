package values

import "fmt"
import "strconv"

// Int64 is an int64 value for the Value interface
type Int64 int64

// NewInt64 returns a new int64 value
func NewInt64(val int64, p *int64) *Int64 {
	*p = val
	return (*Int64)(p)
}

// Get returns the value interface
func (i *Int64) Get() interface{} {
	return int64(*i)
}

// Set sets the value from a string
func (i *Int64) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = Int64(v)
	return err
}

// String returns the value as a string
func (i *Int64) String() string {
	return fmt.Sprintf("%v", *i)
}
