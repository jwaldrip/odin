package cli_test

import (
	. "github.com/jwaldrip/odin/cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Param Parsing", func() {

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

		cli.DefineParams("paramA", "paramB")
	})

	It("should set the parameters by position", func() {
		cli.Start("cmd", "foo", "bar")
		Expect(cmd.Param("paramA").Get()).To(Equal("foo"))
		Expect(cmd.Param("paramB").Get()).To(Equal("bar"))
		Expect(cmd.Params()).To(
			Equal(
				map[string]Value{"paramA": cmd.Param("paramA"), "paramB": cmd.Param("paramB")},
			),
		)
	})

	Context("when a paramter is mising", func() {
		It("should raise an error", func() {
			Ω(func() { cli.Start("cmd") }).Should(Panic())
		})
	})

})
