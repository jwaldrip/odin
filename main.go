package main

import . "github.com/jwaldrip/odin/cli"
import "os"
import "fmt"

var cli = NewCLI(startCmd)

func init(){
  cli.SetVersion("0.1.0")
  cli.DefineBoolFlag("good", false, "sets if everything is good")
  cli.AliasFlag('g', "good")
  cli.DefineSubCommand("hello", "say hello world", helloCmd)
  cli.SetDescription("A command line DSL for go")
}

func main(){
  cli.Start()
}

func helloCmd(cmd Command) error {
  fmt.Println("hello world")
  return nil
}

func startCmd(cmd Command) error {
  // cmd.Usage()
  os.Exit(1)
  return nil
}
