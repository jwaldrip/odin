package main

import odin "github.com/jwaldrip/odin/cli"

// VERSION is the odin version
var VERSION = "1.0.1"

var cli = odin.NewCLI(VERSION, "a command line DSL for go-lang", func(cmd odin.Command) { cmd.Usage() })

func main() {
	cli.Start()
}
