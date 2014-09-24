package main

import odin "github.com/jwaldrip/odin/cli"

var cli = odin.NewCLI(func(cmd odin.Command) { cmd.Usage() })

func init() {
	cli.SetVersion("1.0.0")
}

func main() {
	cli.Start()
}
