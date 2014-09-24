package cli

type paramsList []*Param

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

func (l paramsList) Names() []string {
	var names []string
	for i := 0; i < len(l); i++ {
		names = append(names, l[i].Name)
	}
	return names
}
