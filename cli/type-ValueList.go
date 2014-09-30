package cli

// ValueList is list of values
type ValueList []Value

// GetAll gets an interface for all values in the list
func (v ValueList) GetAll() []interface{} {
	var interfaces []interface{}
	for _, value := range v {
		interfaces = append(interfaces, value.Get())
	}
	return interfaces
}

// Strings returns all the values as their strings
func (v ValueList) Strings() []string {
	var strings []string
	for _, value := range v {
		strings = append(strings, value.String())
	}
	return strings
}
