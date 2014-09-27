package cli_test

import (
	. "github.com/jwaldrip/odin/cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"
)

var _ = Describe("CLI", func() {

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

	Describe("flag definitions", func() {
		Describe("BoolFlag", func() {

			It("should set and parse", func() {
				cli.DefineBoolFlag("boolflag", false, "desc")
				cli.Start("cmd", "--boolflag")
				Expect(cmd.Flag("boolflag").Get()).To(Equal(true))
			})

			It("should set and parse from var", func() {
				var val bool
				cli.DefineBoolFlagVar(&val, "boolflag", false, "desc")
				cli.DefineBoolFlagVar(&val, "boolflag2", false, "desc")
				cli.Start("cmd", "--boolflag")
				Expect(cmd.Flag("boolflag2").Get()).To(Equal(true))
			})
		})

		Describe("DurationFlag", func() {

			It("should set and parse", func() {
				cli.DefineDurationFlag("durflag", time.Duration(60)*time.Second, "desc")
				cli.Start("cmd", "--durflag=30s")
				Expect(cmd.Flag("durflag").Get()).To(Equal(time.Duration(30) * time.Second))
			})

			It("should set and parse from var", func() {
				var val time.Duration
				cli.DefineDurationFlagVar(&val, "durflag", time.Duration(60)*time.Second, "desc")
				cli.DefineDurationFlagVar(&val, "durflag2", time.Duration(45)*time.Second, "desc")
				cli.Start("cmd", "--durflag=100s")
				Expect(cmd.Flag("durflag2").Get()).To(Equal(time.Duration(100) * time.Second))
			})
		})

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

	})

	Describe("remaining arguments", func() {
		It("should return any arguments that have not been specified", func() {
			cli.Start("cmd", "super", "awesome", "dude")
			Expect(cmd.Args()).To(Equal([]string{"super", "awesome", "dude"}))
			Expect(cmd.Arg(0)).To(Equal("super"))
			Expect(cmd.Arg(1)).To(Equal("awesome"))
			Expect(cmd.Arg(2)).To(Equal("dude"))
		})
	})

	Describe("subcommands", func() {

		var didRunSub bool

		BeforeEach(func() {
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
				Ω(func() { cli.Start("cmd", "bad") }).Should(Panic())
			})
		})
	})

})
