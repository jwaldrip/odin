package cli

import "sort"

type flagMap map[string]*Flag

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
