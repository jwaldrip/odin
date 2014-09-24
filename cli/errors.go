package cli

import "log"
import "os"

var stderr = log.New(os.Stderr, "", 0)
var stdout = log.New(os.Stdout, "", 0)

func exitIfError(err error) {
	if err != nil {
		stderr.Fatalln(err)
	}
}
