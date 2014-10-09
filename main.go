package main

import odin "github.com/jwaldrip/odin/cli"

// VERSION is the odin version
var VERSION = "1.3.0"

var cli = odin.New(VERSION, "a command line DSL for go-lang", func(cmd odin.Command) { cmd.Usage() })

func main() {
	cli.Start()
}
