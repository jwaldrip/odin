# odin [![godoc](http://img.shields.io/badge/Go-Doc-blue.svg)](https://godoc.org/github.com/jwaldrip/odin/cli) [![Build Status](https://travis-ci.org/jwaldrip/odin.svg?branch=master)](https://travis-ci.org/jwaldrip/odin) [![Coverage Status](https://img.shields.io/coveralls/jwaldrip/odin.svg)](https://coveralls.io/r/jwaldrip/odin?branch=master)

> _**A note to readers:**  
> I did my best to make this readme as comprehensive as possible but
> I strongly suggest reading the [godocs](https://godoc.org/github.com/jwaldrip/odin/cli) detailed
> information on the internals of odin._

---

> [Installation](#installation) | [Usage](#usage) | [Examples](https://github.com/jwaldrip/odin/tree/master/examples) | [Contributing](https://github.com/jwaldrip/odin/blob/master/CONTRIBUTING.md)

![odin](https://raw.githubusercontent.com/jwaldrip/odin/master/odin.png)

A go-lang library to help build self documenting command line applications.

**odin currently supports:**

* Required Parameters
* Typed Flag/Option Parsing
* Flag Aliasing
* SubCommand DSL for creating complex CLI utils

## Installation

get the package with:

```
go get github.com/jwaldrip/odin/cli
```

## Usage

> [Creating a new CLI](#creating-a-new-cli) | [Flags](#flags) | [Required Parameters](#required-parameters) | [Freeform Parameters](#freeform-parameters) | [Sub Commands](#sub-commands) | [Self Documentation](#self-documentation)

### Creating a new CLI

```go
NewCLI(version string, description string, fn func(), params...)
```

### Flags

Flags are optional parameters that can be specifed when running the command.

```
$ cmd --flag=value -g
```

#### Defining

Flags support a number of different types and can be defined by using a basic definition or pointer definition.

*Basic Definitions are in the format of:*
```go
DefineTypeFlag(name string, deafultValue typedValue, usage string)
```
*Pointer Definitions are in the format of:*
```go
var ptr TYPE
DefineTypeFlag(ptr, name string, defaultValue typedValue, usage string)
```

#### Supported Value Types
| Value Type    | Basic Definition Method | Pointer Defintion Method
|:---           |:---                     |:---
| bool          | `DefineBoolFlag`        | `DefineBoolFlagVar`
| float64       | `DefineFloat64Flag`     | `DefineFloat64FlagVar`
| int           | `DefineIntFlag`         | `DefineIntFlagVar`
| int64         | `DefineInt64Flag`       | `DefineInt64FlagVar`
| string        | `DefineStringFlag`      | `DefineStringFlagVar`
| time.Duration | `DefineDurationFlag`    | `DefineDurationFlagVar`
| uint          | `DefineUintFlag`        | `DefineUintFlagVar`  
| uint64        | `DefineUint64Flag`      | `DefineUint64FlagVar`

*Flags also support aliases:*
aliases are always defined as a `rune` to limit them to one character.

```go
FlagAlias(alias rune, flagName string)
```

#### Example

```go
package main

import (
	"fmt"

	"github.com/jwaldrip/odin/cli"
)

// CLI is the odin CLI
var CLI = NewCLI("0.0.1", "my cli", func(c cli.Command){
	if c.Flag("gopher").Get() == true {
		fmt.Println("IT IS JUST GOPHERTASTIC!!!")
	} else {
		fmt.Println("It is just fine")
	}
})

func init(){
	CLI.DefineBoolFlag("gopher", false, "is it gophertastic?")
	CLI.FlagAlias('g', "gopher")
}

func main(){
	CLI.Start()
}
```

```
$ mycli
It is just fine

$ mycli --gopher
IT IS JUST GOPHERTASTIC!!!

$ mycli -g
IT IS JUST GOPHERTASTIC!!!
```

### Required Parameters

#### Defining

Commands can specify parameters that they require in the order they are passed. All parameter arguments are treated as string values.

*They can be defined at CLI creation...*

```go
NewCLI(version string, description string, fn commandFn, params ...string)
```

*or at a later time...*

```go
cli.DefineParams(params ...string)
```

#### Accessing

*Parameters can be accessed invidually...*

```go
	cmd.Param(name string) Value
```

*or as a `map[string]Value` where the name is the string key of the map.*

```go
	cmd.Params() map[string]Value
```

### Freeform Parameters

Free form parameters are any parameters that remain after all parsing has completed.

*To access them you can call them by their order...*

```go
cmd.Arg(0) Value
```

*or just by getting them all as a `slice`*

```go
cmd.Args() []Value
```

### Sub Commands

Sub commands can be defined in order to create chainable, complex command line utlilities.

#### Defining

Sub commands have all the same defintion methods as the root level command, with one caveat; they have a `Parent()` method that can be used to fetch parameters and flags from further up the command chain.

```go
mycli.DefineSubCommand(name string, description string, fn commandFn, params ...string)
```

#### Accessing Params + Flags

You would access the sub-command's flags and params just as you would a normal root level CLI.

```go
cmd.Param("name")
cmd.Flag("name")
```

*To access a parent commands just call:*

```go
cmd.Parent().Param("name")
cmd.Parent().Flag("name")
```

### Self Documentation

#### Usage

Documentation is auto generated for all commands (main and sub). By Default a help flag `--help` and alias `-h` are defined on each command (unless overridden) to display usage to the end user.

**Auto generated documentation for [./examples/greet-with](https://github.com/jwaldrip/odin/tree/master/examples/greet-with):**

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

#### Version

By Default a version flag `--version` and alias `-v` are defined on the main command (unless overridden), which will display the version specified when [creating a new CLI](#creating-a-new-cli).

**Version for [./examples/greet-with](https://github.com/jwaldrip/odin/tree/master/examples/greet-with):**

```
$ greet-with -v
greet-with 1.0.0
```

## Examples

*Example CLIs can be found in [./examples](https://github.com/jwaldrip/odin/tree/master/examples)*

## Contributing

See [CONTRIBUTING](https://github.com/jwaldrip/odin/blob/master/CONTRIBUTING.md) for details on how to contribute.
