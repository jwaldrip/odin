package main

import (
	"os"

	"github.com/jwaldrip/odin/cli"
)
import "fmt"
import "strings"

// Create a new cli application
var app = cli.New("1.0.0", "A simple tool to greet with", greet, "greeting")

func init() {
	// define a flag called --loudly and also alias it to -l
	app.DefineBoolFlag("loudly", false, "say loudly")
	app.AliasFlag('l', "loudly")

	// define a flag called --color and also alias it to -c
	app.DefineStringFlag("color", "blue", "color the output (red, blue, green)")
	app.AliasFlag('c', "color")

	// add a sub command `to` that takes one param called `greetee`
	subcmd := app.DefineSubCommand("to", "greet a person", greetGreetee, "greetee")
	subcmd.SetLongDescription(`
Say a greeting to a specific persion

Example:
  $ greet-with hello to jerry
  hello jerry
	`)

	// tell the subcommand to inherit the flags we just defined
	subcmd.InheritFlags("color", "loudly")
}

func main() {
	// Run the app
	app.Start()
}

// the greet command run by the root command
func greet(c cli.Command) {
	greeting := c.Param("greeting")
	str := fmt.Sprintf("%s", greeting)
	str = styleByFlags(str, c)
	c.Println(str)
}

func greetGreetee(c cli.Command) {
	greeting := c.Parent().Param("greeting")
	greetee := c.Param("greetee")
	str := fmt.Sprintf("%s %s", greeting, greetee)
	str = styleByFlags(str, c)
	c.Println(str, strings.Join(c.Args().Strings(), " "))
}

func styleByFlags(str string, c cli.Command) string {
	if c.Flag("loudly").Get() == true {
		str = louden(str)
	}
	if c.Flag("color").String() != "" {
		str = colorize(str, c.Flag("color").String())
	}
	return str
}

func louden(str string) string {
	return strings.ToUpper(str)
}

func colorize(str string, color string) string {
	switch color {
	case "red":
		str = fmt.Sprintf("\x1b[0;31;49m%s\x1b[0m", str)
	case "blue":
		str = fmt.Sprintf("\x1b[0;34;49m%s\x1b[0m", str)
	case "green":
		str = fmt.Sprintf("\x1b[0;32;49m%s\x1b[0m", str)
	default:
		fmt.Fprintln(os.Stderr, fmt.Sprintf("invalid color: '%s' ... try: 'red', 'blue', or 'green'", color))
		os.Exit(2)
	}
	return str
}
