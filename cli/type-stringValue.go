package cli

import "fmt"

// -- string Value
type stringValue string

func newStringValue(val string, p *string) *stringValue {
  *p = val
  return (*stringValue)(p)
}

func (this *stringValue) Set(val string) error {
  *this = stringValue(val)
  return nil
}

func (this *stringValue) Get() interface{} {
  return string(*this)
}

func (this *stringValue) String() string {
  return fmt.Sprintf("%s", *this)
}
