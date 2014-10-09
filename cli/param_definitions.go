package cli

// DefineParams sets params names from strings
func (cmd *CLI) DefineParams(names ...string) {
	var params []*Param
	for _, name := range names {
		param := &Param{Name: name}
		params = append(params, param)
	}
	cmd.params = params
}
