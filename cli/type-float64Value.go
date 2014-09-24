package cli

import "fmt"
import "strconv"

// -- float64 Value
type float64Value float64

func newFloat64Value(val float64, p *float64) *float64Value {
	*p = val
	return (*float64Value)(p)
}

func (this *float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*this = float64Value(v)
	return err
}

func (this *float64Value) Get() interface{} {
	return float64(*this)
}

func (this *float64Value) String() string {
	return fmt.Sprintf("%v", *this)
}
