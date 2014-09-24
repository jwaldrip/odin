package main

import odin "github.com/jwaldrip/odin/cli"

import "fmt"
import "strings"

var cli = odin.NewCLI(func(cmd odin.Command) { cmd.Usage() })

func init() {
	cli.SetVersion("1.0.0")
	saycmd := cli.DefineSubCommand("say", "say a greeting", greet, "greeting", "greetee")

	saycmd.DefineBoolFlag("loudly", false, "say something loudly")
	saycmd.AliasFlag('l', "loudly")

	cli.SetDescription("A command line DSL for go")
}

func main() {
	cli.Start()
}

func greet(c odin.Command) {
	str := fmt.Sprintf("%s %s", c.Param("greeting"), c.Param("greetee"))
	if c.Flag("loudly").Get() == true {
		str = strings.ToUpper(str)
	}
	fmt.Println(str)
}
