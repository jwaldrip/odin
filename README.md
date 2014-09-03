odin
====

a go library to help build command line applications

Example
====

```go
package say

import "fmt"
import "github.com/jwaldrip/odin"
import "os"

var sayCommand = odin.NewCommand

func init(){
  sayCommand.NewBoolOption("verbose", say)
  sayCommand.Alias("v")
}

func main(){
  sayCommand.Start(os.Args)
}

func say(command *odin.Command){
  args := command.Args
  if command.GetOption("verbose") {
    fmt.Println("I will be doing thing verbosely")
  }
  worldCommand := odin.NewCommand("world", world)
  worldCommand.Start(args)
}

func world(command *odin.Command){
  args := command.Args
  
}

```

