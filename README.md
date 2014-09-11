odin
====

a go library to help build command line applications

Basic Example
====

```go
package say

import "fmt"
import "github.com/jwaldrip/odin"
import "os"

var sayCommand = odin.NewCommand(".", start, "greeting", "object")

func main(){
  sayCommand.StartDefault(os.Args[1:])
}

func say(cmd *odin.Command){
  line := fmt.Sprintf(
    "I would like to say... %[1]v, %[2]v",
    cmd.params["greeting"],
    cmd.params["object"]
  )
  fmt.Println(line)
}
```

```sh
$ say hello world
I would like to say... hello, world
```

Basic Example W/ Flags
====

```go
package say

import "fmt"
import "github.com/jwaldrip/odin"
import "os"
import "strings"

var sayCommand = odin.NewCommand(".", start, "greeting", "object")

func init(){
  sayCommand.BoolFlag("loud", false, "make it loud")
  sayCommand.AliasFlag("loud", "l")
}

func main(){
  sayCommand.StartDefault(os.Args[1:])
}

func say(cmd *odin.Command){
  line := fmt.Sprintf(
    "I would like to say... %[1]v, %[2]v",
    cmd.params["greeting"],
    cmd.params["object"]
  )
  if odin.flags["loud"] {
    line = string.ToUpper(line)
  }
  fmt.Println(line)
}
```

```sh
$ say hello world
I would like to say... hello, world

$ say --loud hello world
I WOULD LIKE TO SAY... HELLO, WORLD

$ say --l hello world
I WOULD LIKE TO SAY... HELLO, WORLD
```
