package cli_test

import (
	. "github.com/jwaldrip/odin/cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Param Definitions", func() {

	var cli *CLI
	var cmd Command
	var didRun bool

	BeforeEach(func() {
		didRun = false
		runFn := func(c Command) {
			cmd = c
			didRun = true
		}
		cli = NewCLI("v1.0.0", "sample description", runFn, "a", "b")
		cli.ErrorHandling = PanicOnError
		cli.Mute()
	})

	Describe("DefineParams", func() {
		It("should define and parse a param list", func() {
			cli.DefineParams("c", "d")
			cli.Start("cmd", "foo", "bar")
			Expect(cmd.Param("c").Get()).To(Equal("foo"))
			Expect(cmd.Param("d").Get()).To(Equal("bar"))
		})
	})

	Describe("Via Start CLI", func() {
		It("should define and parse a param list", func() {
			cli.Start("cmd", "foo", "bar")
			Expect(cmd.Param("a").Get()).To(Equal("foo"))
			Expect(cmd.Param("b").Get()).To(Equal("bar"))
		})
	})

})
