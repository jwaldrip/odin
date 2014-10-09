package cli_test

import (
	. "github.com/jwaldrip/odin/cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CLI Start", func() {

	var cli *CLI
	var sub *SubCommand
	var cmd Command
	var subcmd Command
	var didRun bool
	var didRunSub bool

	BeforeEach(func() {
		didRun = false
		runFn := func(c Command) {
			cmd = c
			didRun = true
		}
		cli = New("v1.0.0", "sample description", runFn)
		cli.ErrorHandling = PanicOnError
		cli.Mute()
		didRunSub = false
		cli.DefineBoolFlag("foo", false, "a foo flag")
		cli.DefineStringFlag("bar", "", "a foo flag")
		sub = cli.DefineSubCommand("razzle", "razzle dazzle me", func(c Command) {
			subcmd = c
			didRunSub = true
		})
	})

	Describe("InheritFlag", func() {
		It("should properly inherit a flag value from its parent", func() {
			sub.InheritFlag("foo")
			cli.Start("cmd", "--foo", "razzle")
			Expect(subcmd.Flag("foo").Get()).To(Equal(true))
		})

		Context("when there is not parent", func() {
			It("should raise an error", func() {
				Expect(cli.Parent()).To(BeNil())
				Expect(func() { cli.InheritFlag("") }).Should(Panic())
			})
		})
	})

	Describe("InheritFlags", func() {
		It("should properly inherit flag values from its parent", func() {
			sub.InheritFlags("foo", "bar")
			cli.Start("cmd", "--foo", "--bar=dive", "razzle")
			Expect(subcmd.Flag("foo").Get()).To(Equal(true))
			Expect(subcmd.Flag("bar").Get()).To(Equal("dive"))
		})
	})

	Describe("SubCommandsInheritFlag", func() {
		It("should propogate its flag to the sub commands", func() {
			cli.SubCommandsInheritFlag("foo")
			cli.Start("cmd", "--foo", "razzle")
			Expect(subcmd.Flag("foo").Get()).To(Equal(true))
		})

		It("should propogate deeply", func() {
			var subsubcmd Command
			sub.DefineSubCommand("baz", "a deeper command", func(c Command) { subsubcmd = c })
			cli.SubCommandsInheritFlag("foo")
			cli.Start("cmd", "--foo", "razzle", "baz")
			Expect(subsubcmd.Flag("foo").Get()).To(Equal(true))
		})
	})

	Describe("SubCommandsInheritFlags", func() {
		It("should propogate its flags to the sub commands", func() {
			cli.SubCommandsInheritFlags("foo", "bar")
			cli.Start("cmd", "--foo", "--bar=dive", "razzle")
			Expect(subcmd.Flag("foo").Get()).To(Equal(true))
			Expect(subcmd.Flag("bar").Get()).To(Equal("dive"))
		})

		It("should propogate deeply", func() {
			var subsubcmd Command
			sub.DefineSubCommand("baz", "a deeper command", func(c Command) { subsubcmd = c })
			cli.SubCommandsInheritFlags("foo", "bar")
			cli.Start("cmd", "--foo", "--bar=dive", "razzle", "baz")
			Expect(subsubcmd.Flag("foo").Get()).To(Equal(true))
			Expect(subsubcmd.Flag("bar").Get()).To(Equal("dive"))
		})
	})
})
