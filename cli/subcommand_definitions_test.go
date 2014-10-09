package cli_test

import (
	. "github.com/jwaldrip/odin/cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
		//cli.Mute()
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
	})

	Describe("AddSubCommand", func() {
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
