package cli

// ValueMap is a map fo values with strings for keys
type ValueMap map[string]Value

// Values returns a ValueList of all the values in the map
func (v *ValueMap) Values() ValueList {
	var valueList ValueList
	for _, value := range *v {
		valueList = append(valueList, value)
	}
	return valueList
}
