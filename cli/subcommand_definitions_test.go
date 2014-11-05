package cli_test

import (
	. "github.com/jwaldrip/odin/cli"

	. "github.com/jwaldrip/odin/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/jwaldrip/odin/Godeps/_workspace/src/github.com/onsi/gomega"
)

var _ = Describe("CLI Start", func() {

	var cli *CLI
	var cmd Command
	var didRun bool

	BeforeEach(func() {
		didRun = false
		runFn := func(c Command) {
			cmd = c
			didRun = true
		}
		cli = New("v1.0.0", "sample description", runFn)
		cli.ErrorHandling = PanicOnError
		cli.Mute()
	})

	Describe("DefineSubCommand", func() {
		It("allow the defined command to run", func() {
			var subDidRun bool
			cli.DefineSubCommand("foo", "sub command", func(c Command) { subDidRun = true })
			cli.Start("cmd", "foo")
			Expect(subDidRun).To(Equal(true))
		})
	})

	Describe("AddSubCommand", func() {
		It("allow the defined command to run", func() {
			var subDidRun bool
			sub := NewSubCommand("foo", "sub command", func(c Command) { subDidRun = true })
			cli.AddSubCommand(sub)
			cli.Start("cmd", "foo")
			Expect(subDidRun).To(Equal(true))
		})

		It("should inherit ErrorHandling and fail when missing params", func() {
			var subDidRun bool
			sub := NewSubCommand("foo", "sub command", func(c Command) { subDidRun = true }, "bar")
			cli.AddSubCommand(sub)
			Expect(func() { cli.Start("cmd", "foo") }).Should(Panic())
			Expect(subDidRun).To(Equal(false))
		})

		It("be able to access its parent", func() {
			var subcmd Command
			sub := NewSubCommand("foo", "sub command", func(c Command) { subcmd = c })
			cli.AddSubCommand(sub)
			cli.Start("cmd", "foo")
			Expect(subcmd.Parent().(*CLI)).To(Equal(cli))
		})

		Context("when the sub command has already been assigned a parent", func() {
			It("should panic", func() {
				sub := NewSubCommand("foo", "sub command", ShowUsage, "bar")
				cli.AddSubCommand(sub)
				Expect(func() { cli.AddSubCommand(sub) }).To(Panic())
			})
		})

		Context("when the sub command is assigned to itself", func() {
			It("should panic", func() {
				sub := NewSubCommand("foo", "sub command", ShowUsage, "bar")
				Expect(func() { sub.AddSubCommand(sub) }).To(Panic())
			})
		})

	})

	Describe("AddSubCommands", func() {
		It("allow the defined command to run", func() {
			var subOneDidRun bool
			var subTwoDidRun bool
			subOne := NewSubCommand("foo", "sub command", func(c Command) { subOneDidRun = true })
			subTwo := NewSubCommand("bar", "sub command", func(c Command) { subTwoDidRun = true })
			cli.AddSubCommands(subOne, subTwo)
			cli.Start("cmd", "foo")
			cli.Start("cmd", "bar")
			Expect(subOneDidRun).To(Equal(true))
			Expect(subTwoDidRun).To(Equal(true))
		})
	})
})
