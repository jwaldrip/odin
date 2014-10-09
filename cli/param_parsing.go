package cli

import (
	"strings"

	"github.com/jwaldrip/odin/cli/values"
)

func (cmd *CLI) parseParams(args []string) []string {
	var seenParams paramsList

	if len(cmd.params) == 0 {
		return args
	}
	i := 0
	for i < len(args) && i < len(cmd.params) {
		param := cmd.params[i]
		seenParams = append(seenParams, param)
		str := ""
		if cmd.paramValues == nil {
			cmd.paramValues = make(map[*Param]values.Value)
		}
		cmd.paramValues[param] = values.NewString(args[i], &str)
		i++
	}
	missingParams := cmd.params.Compare(seenParams)
	if len(missingParams) > 0 {
		var msg string
		if len(missingParams) == 1 {
			msg = "missing param"
		} else {
			msg = "missing params"
		}
		cmd.errf("%s: %s", msg, strings.Join(missingParams.Names(), ", "))
	}

	return args[i:]
}
