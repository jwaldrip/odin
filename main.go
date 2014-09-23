package main

import . "github.com/jwaldrip/odin/cli"
import "os"
//import "fmt"

var cli = NewCLI(startCmd)

func init(){
  cli.Version = "0.1.0"
  cli.ErrorHandling = ExitOnError
  cli.DefineBoolFlag("good", false, "sets if everything is good")
  cli.AliasFlag('g', "good")
  cli.DefineSubCommand("foo", foo)
}

func main(){
  cli.Start()
}

func foo(cmd Command) error {
  return nil
}

func startCmd(cmd Command) error {
  cmd.Usage()
  os.Exit(1)
  return nil
}
