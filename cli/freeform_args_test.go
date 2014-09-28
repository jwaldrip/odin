package cli_test

import (
	. "github.com/jwaldrip/odin/cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Freeform Args Parsing", func() {

	var cli *CLI
	var cmd Command
	var didRun bool

	BeforeEach(func() {
		didRun = false
		runFn := func(c Command) {
			cmd = c
			didRun = true
		}
		cli = NewCLI("v1.0.0", "sample description", runFn)
		cli.ErrorHandling = PanicOnError
		cli.Mute()
	})

	It("should return any arguments that have not been specified", func() {
		cli.Start("cmd", "super", "awesome", "dude")
		Expect(cmd.Arg(0).Get()).To(Equal("super"))
		Expect(cmd.Arg(1).Get()).To(Equal("awesome"))
		Expect(cmd.Arg(2).Get()).To(Equal("dude"))
	})

	Context("once flags are terminated", func() {
		It("should return what would usually be flag values", func() {
			cli.DefineBoolFlag("sample", false, "a sample flag")
			cli.Start("cmd", "--", "--sample=true")
			Expect(cmd.Flag("sample").Get()).To(Equal(false))
			Expect(cmd.Arg(0).Get()).To(Equal("--sample=true"))
		})
	})

})
