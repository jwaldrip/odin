package cli_test

import (
	. "github.com/jwaldrip/odin/cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flag parsing", func() {

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

		cli.DefineBoolFlag("foo", false, "is foo")
		cli.AliasFlag('o', "foo")
		cli.DefineStringFlag("bar", "", "what bar are you at?")
		cli.AliasFlag('r', "bar")
		cli.DefineBoolFlag("baz", true, "is baz")
		cli.AliasFlag('z', "baz")
	})

	It("should set the flags with set syntax", func() {
		cli.Start("cmd", "--bar=squeaky bean")
		Expect(cmd.Flag("bar").Get()).To(Equal("squeaky bean"))
	})

	It("should set the flags with positional syntax", func() {
		cli.Start("cmd", "--bar", "squeaky bean")
		Expect(cmd.Flag("bar").Get()).To(Equal("squeaky bean"))
	})

	Context("invalid flags", func() {
		It("undefined should panic", func() {
			Ω(func() { cli.Start("cmd", "--undefined") }).Should(Panic())
		})

		It("malformed should panic", func() {
			Ω(func() { cli.Start("cmd", "-=") }).Should(Panic())
		})

		It("improper value should panic", func() {
			cli.DefineBoolFlag("bool", false, "")
			Ω(func() { cli.Start("cmd", "--bool=funny") }).Should(Panic())
		})
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

		It("should not support positional setting", func() {
			cli.Start("cmd", "--foo", "false")
			Expect(cmd.Flag("foo").Get()).To(Equal(true))
			Expect(cmd.Arg(0).Get()).To(Equal("false"))
		})

	})

	Context("when an invalid flag was passed", func() {
		It("should raise an error", func() {
			Ω(func() { cli.Start("cmd", "--bad") }).Should(Panic())
		})
	})

	Context("when a non-boolflag was not provided a value", func() {
		It("should raise an error", func() {
			Ω(func() { cli.Start("cmd", "--bar") }).Should(Panic())
		})
	})

	Context("with aliases", func() {

		It("should set the last flag with set syntax", func() {
			cli.Start("cmd", "-or=dive bar")
			Expect(cmd.Flag("foo").Get()).To(Equal(true))
			Expect(cmd.Flag("bar").Get()).To(Equal("dive bar"))
		})

		It("should set the last flag with positional syntax", func() {
			cli.Start("cmd", "-or", "dive bar")
			Expect(cmd.Flag("foo").Get()).To(Equal(true))
			Expect(cmd.Flag("bar").Get()).To(Equal("dive bar"))
		})

		Context("when an invalid alias was passed", func() {
			It("should raise an error", func() {
				Ω(func() { cli.Start("cmd", "-op") }).Should(Panic())
			})
		})

		Context("when a non-boolflag was not provided a value", func() {
			It("should raise an error", func() {
				Ω(func() { cli.Start("cmd", "-or") }).Should(Panic())
			})
		})

	})

	Context("once flags are terminated", func() {
		Context("with --", func() {
			It("should not parse additional flags", func() {
				cli.DefineBoolFlag("sample", false, "a sample flag")
				cli.Start("cmd", "--", "--sample=true")
				Expect(cmd.Flag("sample").Get()).To(Equal(false))
				Expect(cmd.Arg(0).Get()).To(Equal("--sample=true"))
			})
		})

		Context("with non flag", func() {
			It("should not parse additional flags", func() {
				cli.DefineBoolFlag("sample", false, "a sample flag")
				cli.Start("cmd", "foo", "--sample=true")
				Expect(cmd.Flag("sample").Get()).To(Equal(false))
				Expect(cmd.Arg(0).Get()).To(Equal("foo"))
				Expect(cmd.Arg(1).Get()).To(Equal("--sample=true"))
			})
		})
	})

})
