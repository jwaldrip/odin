# odin [![godoc](http://img.shields.io/badge/Go-Doc-blue.svg)](https://godoc.org/github.com/jwaldrip/odin/cli)

![odin](https://github.com/jwaldrip/odin/blob/master/odin.png)]

A go-lang library to help build self documenting command line applications.

**odin currently supports:**

* Flag/Option Parsing
* Flag Aliasing
* Named Command Parameters (currently all required)
* SubCommand DSL for creating complex CLI utils

## Installation

get the package with:

```
go get github.com/jwaldrip/odin/cli
```

## Usage

```go
NewCLI(version string, description string, fn func(), params...)
```

## Example

**the following example can be found in "example/greet-with"**

```go
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
  str := fmt.Sprintf("%s %s", c.Parent().Param("greeting"), c.Param("greetee"))
  if c.Parent().Flag("loudly").Get() == true {
    str = strings.ToUpper(str)
  }
  if c.Parent().Flag("color").String() != "" {
    println("  gg")
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
    fmt.Fprintln(os.Stderr, "invalid color, try: red, blue, or green")
    os.Exit(2)
  }
  return coloredString
}
```

## Self Documentation

Documentation is auto generated for all commands.

**the above example would output the following for its documentation:**

```
$ greet-with -h
Usage:
  greet-with [options...] <greeting> <command> [arg...]

a simple tool to greet with

Options:
  -c, --color="blue"  # color the output (red, blue, green)
  -h, --help          # show help and exit
  -l, --loudly        # say loudly
  -v, --version       # show version and exit

Commands:
  to     greet a person
```

## Contributing

See [CONTRIBUTING](https://github.com/jwaldrip/odin/blob/master/CONTRIBUTING.md) for details on how to contribute.
