package values

import "time"

// Duration is a Time.Duration value for the Value interface
type Duration time.Duration

// NewDuration returns a new time.Duration value
func NewDuration(val time.Duration, p *time.Duration) *Duration {
	*p = val
	return (*Duration)(p)
}

// Get returns the value interface
func (d *Duration) Get() interface{} {
	return time.Duration(*d)
}

// Set sets the value from a string
func (d *Duration) Set(s string) error {
	v, err := time.ParseDuration(s)
	*d = Duration(v)
	return err
}

// String returns the value as a string
func (d *Duration) String() string {
	return (*time.Duration)(d).String()
}
