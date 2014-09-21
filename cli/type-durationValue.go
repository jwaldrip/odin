package cli

import "time"

// -- time.Duration Value
type durationValue time.Duration

func newDurationValue(val time.Duration, p *time.Duration) *durationValue {
  *p = val
  return (*durationValue)(p)
}

func (this *durationValue) Set(s string) error {
  v, err := time.ParseDuration(s)
  *this = durationValue(v)
  return err
}

func (this *durationValue) Get() interface{} {
  return time.Duration(*this)
}

func (this *durationValue) String() string {
  return (*time.Duration)(this).String()
}
