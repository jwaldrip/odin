package cli_test

import (
	. "github.com/jwaldrip/odin/cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CLI", func() {

	var cli *CLI
	var cmd Command
	var didRun bool

	BeforeEach(func() {
		runFn := func(c Command) {
			cmd = c
			didRun = true
		}
		cli = NewCLI("v1.0.0", "sample description", runFn)
		cli.ErrorHandling = PanicOnError
		cli.Mute()
	})

	Describe("required parameters", func() {

		BeforeEach(func() {
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

	Describe("flags", func() {

		BeforeEach(func() {
			cli.DefineBoolFlag("foo", false, "is foo")
			cli.DefineStringFlag("bar", "", "what bar are you at?")
			cli.DefineBoolFlag("baz", true, "is baz")
		})

		It("should set the flags with set syntax", func() {
			cli.Start("cmd", "--bar=squeaky bean")
			Expect(cmd.Flag("bar").Get()).To(Equal("squeaky bean"))
		})

		It("should set the flags with positional syntax", func() {
			cli.Start("cmd", "--bar", "squeaky bean")
			Expect(cmd.Flag("bar").Get()).To(Equal("squeaky bean"))
		})

		Context("boolean flags", func() {

			It("should set boolean flags as true if set", func() {
				cli.Start("cmd", "--foo", "--baz")
				Expect(cmd.Flag("foo").Get()).To(Equal(true))
				Expect(cmd.Flag("baz").Get()).To(Equal(true))
			})

			It("should set as the default value true if not set", func() {
				cli.Start("cmd")
				Expect(cmd.Flag("foo").Get()).To(Equal(false))
				Expect(cmd.Flag("baz").Get()).To(Equal(true))
			})

		})

		Context("when an invalid flag was passed", func() {
			It("should raise an error", func() {

			})
		})

		Context("when a non-boolflag was not provided a value", func() {
			It("should raise an error", func() {

			})
		})

		Context("with aliases", func() {

			It("should set the last flag with set syntax", func() {

			})

			It("should set the last flag with set syntax", func() {

			})

			Context("when an invalid alias was passed", func() {
				It("should raise an error", func() {

				})
			})

			Context("when a non-boolflag was not provided a value", func() {
				It("should raise an error", func() {

				})
			})

		})

	})

	Describe("remaining arguments", func() {
		It("should return any arguments that have not been specified", func() {

		})
	})

	Describe("subcommands", func() {
		Context("when the subcommand is valid", func() {
			It("should start a subcommand", func() {

			})
		})

		Context("when the subcommand is not valid", func() {
			It("should raise an error", func() {

			})
		})
	})

})
