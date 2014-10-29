package main

import "github.com/jwaldrip/odin/cli"

var app = cli.New(version, "a command line DSL for go-lang", cli.ShowUsage)

func main() {
	app.Start()
}
