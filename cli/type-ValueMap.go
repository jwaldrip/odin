package cli

// ValueMap is a map fo values with strings for keys
type ValueMap map[string]Value

// Keys returns the keys of the value map
func (v ValueMap) Keys() []string {
	var keys []string
	for key := range v {
		keys = append(keys, key)
	}
	return keys
}

// Values returns a ValueList of all the values in the map
func (v ValueMap) Values() ValueList {
	var valueList ValueList
	for _, value := range v {
		valueList = append(valueList, value)
	}
	return valueList
}
