package main

import odin "github.com/jwaldrip/odin/cli"

// VERSION is the odin version
var VERSION = "0.9.0"

var cli = odin.NewCLI(VERSION, "a command line DSL for go-lang", func(cmd odin.Command) { cmd.Usage() })

func main() {
	cli.Start()
}
