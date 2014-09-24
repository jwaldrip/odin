package cli

import . "log"
import "os"

var stderr = New(os.Stderr, "", 0)
var stdout = New(os.Stdout, "", 0)

func exitIfError(error error) {
	if error != nil {
		stderr.Fatalln(error)
	}
}
