package cli

import "github.com/jwaldrip/odin/cli/values"

func (cmd *CLI) assignUnparsedArgs(args []string) {
	for i := 0; i < len(args); i++ {
		str := ""
		cmd.unparsedArgs = append(cmd.unparsedArgs, values.NewString(args[i], &str))
	}
}
