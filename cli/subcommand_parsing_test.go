package cli_test

import (
	"bytes"

	. "github.com/jwaldrip/odin/cli"

	. "github.com/jwaldrip/odin/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/jwaldrip/odin/Godeps/_workspace/src/github.com/onsi/gomega"
)

var _ = Describe("CLI Start", func() {

	var cli *CLI
	var cmd Command
	var didRun bool
	var didRunSub bool
	var errout *bytes.Buffer

	BeforeEach(func() {
		errout = bytes.NewBufferString("")
		didRun = false
		runFn := func(c Command) {
			cmd = c
			didRun = true
		}
		cli = New("v1.0.0", "sample description", runFn)
		cli.ErrorHandling = PanicOnError
		cli.Mute()
		cli.SetErrOutput(errout)
		didRunSub = false
		cli.DefineSubCommand("razzle", "razzle dazzle me", func(c Command) {
			didRunSub = true
		})
	})

	Context("when the subcommand is valid", func() {
		It("should start a subcommand", func() {
			cli.Start("cmd", "razzle")
			Expect(didRunSub).To(Equal(true))
		})
	})

	Context("when the subcommand is not valid", func() {
		It("should raise an error", func() {
			Expect(func() { cli.Start("cmd", "bad") }).Should(Panic())
			Expect(errout.String()).To(ContainSubstring("invalid command: bad"))
		})
	})

})
