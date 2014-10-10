package cli

import "sort"

// flagMap is a map of flags with the name as a string key
type flagMap map[string]*Flag

func (fm flagMap) Merge(fm2 flagMap) flagMap {
	mergedMap := make(flagMap)
	if fm != nil {
		for k, v := range fm {
			mergedMap[k] = v
		}
	}
	if fm2 != nil {
		for k, v := range fm2 {
			mergedMap[k] = v
		}
	}
	return mergedMap
}

func (fm flagMap) Names() []string {
	var keys []string
	for k := range fm {
		keys = append(keys, k)
	}
	return keys
}

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

func (fm flagMap) Without(fm2 flagMap) flagMap {
	diffedMap := make(flagMap)
	if fm == nil {
		return diffedMap
	}
	if fm2 == nil {
		return fm
	}
	for k, v := range fm {
		if _, exist := fm2[k]; !exist {
			diffedMap[k] = v
		}
	}
	return diffedMap
}
