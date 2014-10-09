package values

// Map is a map fo values with strings for keys
type Map map[string]Value

// Keys returns the keys of the value map
func (v Map) Keys() []string {
	var keys []string
	for key := range v {
		keys = append(keys, key)
	}
	return keys
}

// Values returns a ValueList of all the values in the map
func (v Map) Values() List {
	var valueList List
	for _, value := range v {
		valueList = append(valueList, value)
	}
	return valueList
}
