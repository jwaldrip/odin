package values

import "fmt"
import "strconv"

// Float64 is a float64 value for the Value interface
type Float64 float64

// NewFloat64 returns a new float64 value
func NewFloat64(val float64, p *float64) *Float64 {
	*p = val
	return (*Float64)(p)
}

// Get returns the value interface
func (f *Float64) Get() interface{} {
	return float64(*f)
}

// Set sets the value from a string
func (f *Float64) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*f = Float64(v)
	return err
}

// String returns the value as a string
func (f *Float64) String() string {
	return fmt.Sprintf("%v", *f)
}
