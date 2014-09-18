package main

import "fmt"
import . "github.com/jwaldrip/odin/cli"
import "errors"

var cli = NewCLI(startCmd, "one", "two")

func init(){
  cli.DefineBoolFlag("good", false, "sets if everything is good")
  cli.AliasFlag("good", "g")
}

func main(){
  cli.Start()
}

func startCmd(cmd Command) error {
  fmt.Println( cmd.Param("one"), cmd.Param("two") )
  fmt.Println( cmd.Flags() )
  return errors.New("err")
}

// NewCommand("say", *say, "greeting", "object")

// The Type for a command Function, returns a command interface

// func (cmd *SubCommand) Start(args []string){
//   if cmd.Name != args[0] { log.Panicln("invalid command") }
//   cmd.parse(args[1:])
// }

// TODO: Define a root level CLI
// type CLI Command
//
// func NewCLI( fn *commandFn, args ...string ) *CLI {
//   cmd := NewCommand(".", fn, args...)
//   var cli CLI
//   cli.Name = cmd.Name
//   cli.args = cmd.args
//   cli.flagSet = cmd.flagSet
//   return &cli
// }
//
// func (cli *CLI) Start(args []string){
//   args = append([]string{cli.Name}, args...)
//   cli.Start(args)
// }
