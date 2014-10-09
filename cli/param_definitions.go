package cli

// DefineParams sets params names from strings
func (cmd *CLI) DefineParams(names ...string) {
	var params []*Param
	for i := 0; i < len(names); i++ {
		name := names[i]
		param := &Param{Name: name}
		params = append(params, param)
	}
	cmd.params = params
}
