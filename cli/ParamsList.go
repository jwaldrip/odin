package cli

// paramsList a list of params
type paramsList []*Param

// Compare compares two lists and returns the difference
func (l paramsList) Compare(Y paramsList) paramsList {
	m := make(map[*Param]int)

	for _, y := range Y {
		m[y]++
	}

	var ret paramsList
	for _, x := range l {
		if m[x] > 0 {
			m[x]--
			continue
		}
		ret = append(ret, x)
	}

	return ret
}

// Names returns the list of parameters names as a slice of strings
func (l paramsList) Names() []string {
	var names []string
	for _, item := range l {
		names = append(names, item.Name)
	}
	return names
}
