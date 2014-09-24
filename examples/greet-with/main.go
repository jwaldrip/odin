package main

import odin "github.com/jwaldrip/odin/cli"
import "github.com/mgutz/ansi"
import "fmt"
import "strings"

var cli = odin.NewCLI(greet, "greeting")

func init() {
	cli.SetVersion("1.0.0")
	cli.SetDescription("a simple tool to greet with")
	cli.DefineBoolFlag("loudly", false, "say loudly")
	cli.AliasFlag('l', "loudly")
	cli.DefineStringFlag("color", "blue", "color the output")
	cli.DefineSubCommand("to", "greet a person", greetGreetee, "greetee")
}

func main() {
	cli.Start()
}

func greet(c odin.Command) {
	str := fmt.Sprintf("%s", c.Param("greeting"))
	if c.Flag("loudly").Get() == true {
		str = strings.ToUpper(str)
	}
	if c.Flag("color").String() != "" {
		str = ansi.Color(str, c.Flag("color").String())
	}
	fmt.Println(str)
}

func greetGreetee(c odin.Command) {
	str := fmt.Sprintf("%s %s", c.Parent().Param("greeting"), c.Param("greetee"))
	if c.Parent().Flag("loudly").Get() == true {
		str = strings.ToUpper(str)
	}
	if c.Parent().Flag("color").String() != "" {
		println("  gg")
		str = ansi.Color(str, c.Parent().Flag("color").String())
	}
	fmt.Println(str)
}
