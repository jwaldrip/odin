package main

import (
	"os"

	odin "github.com/jwaldrip/odin/cli"
)
import "fmt"
import "strings"

type colorfulString string

var cli = odin.NewCLI("1.0.0", "a simple tool to greet with", greet, "greeting")

func init() {
	cli.DefineBoolFlag("loudly", false, "say loudly")
	cli.AliasFlag('l', "loudly")
	cli.DefineStringFlag("color", "blue", "color the output (red, blue, green)")
	cli.AliasFlag('c', "color")
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
		str = colorfulString(str).color(c.Flag("color").String())
	}
	fmt.Println(str)
}

func greetGreetee(c odin.Command) {
	greeting := c.Parent().Param("greeting")
	greetee := c.Param("greetee")
	str := fmt.Sprintf("%s %s", greeting, greetee)
	if c.Parent().Flag("loudly").Get() == true {
		str = strings.ToUpper(str)
	}
	if c.Parent().Flag("color").String() != "" {
		str = colorfulString(str).color(c.Parent().Flag("color").String())
	}
	fmt.Println(str)
}

func (s colorfulString) color(color string) string {
	var coloredString string
	switch color {
	case "red":
		coloredString = fmt.Sprintf("\x1b[0;31;49m%s\x1b[0m", s)
	case "blue":
		coloredString = fmt.Sprintf("\x1b[0;34;49m%s\x1b[0m", s)
	case "green":
		coloredString = fmt.Sprintf("\x1b[0;32;49m%s\x1b[0m", s)
	default:
		fmt.Fprintln(os.Stderr, fmt.Sprintf("invalid color: '%s' ... try: 'red', 'blue', or 'green'", color))
		os.Exit(2)
	}
	return coloredString
}
