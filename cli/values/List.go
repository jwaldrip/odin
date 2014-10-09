package values

// List is list of values
type List []Value

// GetAll gets an interface for all values in the list
func (v List) GetAll() []interface{} {
	var interfaces []interface{}
	for _, value := range v {
		interfaces = append(interfaces, value.Get())
	}
	return interfaces
}

// Strings returns all the values as their strings
func (v List) Strings() []string {
	var strings []string
	for _, value := range v {
		strings = append(strings, value.String())
	}
	return strings
}
