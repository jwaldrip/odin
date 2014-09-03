odin
====

a go library to help build command line applications

Example
====

```go
import "fmt"
import "github.com/jwaldrip/odin"

func init(){
  odin.MainCommand.NewBoolOption("verbose", run).Alias("v")
}

func main(){
  
}

func run(command *odin.Command){
  if command.GetOption("verbose") {
    fmt.Println("hello world verbosely")
  } else {
    fmt.Println("hello world")
  }
}

```

