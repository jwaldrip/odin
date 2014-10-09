package cli

import "sort"

// flagMap is a map of flags with the name as a string key
type flagMap map[string]*Flag

// Sort returns a sorted list of flags
func (fm flagMap) Sort() []*Flag {
	list := make(sort.StringSlice, len(fm))
	i := 0
	for _, f := range fm {
		list[i] = f.Name
		i++
	}
	list.Sort()
	result := make([]*Flag, len(list))
	for i, name := range list {
		result[i] = fm[name]
	}
	return result
}

func (fm flagMap) Names() []string {
	var keys []string
	for k := range fm {
		keys = append(keys, k)
	}
	return keys
}
