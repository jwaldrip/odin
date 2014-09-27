package cli_test

import (
	"strings"

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
		cli = NewCLI("v1.0.0", "sample description", runFn)
		cli.ErrorHandling = PanicOnError
		cli.Mute()
	})

	Describe("complex cli", func() {

		var subDidRun bool
		var subCmd *SubCommand

		BeforeEach(func() {
			cli.DefineParams("host", "path")
			cli.DefineBoolFlag("ssl", false, "do it over ssl")
			cli.AliasFlag('S', "ssl")
			cli.DefineStringFlag("username", "", "the username")
			cli.AliasFlag('u', "username")
			cli.DefineStringFlag("password", "", "the password")
			cli.AliasFlag('p', "password")
			cli.DefineIntFlag("port", 80, "the port")
			cli.AliasFlag('P', "port")
			subCmd = cli.DefineSubCommand("do", "what action to do", func(c Command) { cmd = c; subDidRun = true }, "action")
		})

		It("should parse the main command properly", func() {
			cli.Start(strings.Split("cmd -Su=wally -p App1etw0 --port 3001 example.com /", " ")...)
			Expect(cmd.Param("host").Get()).To(Equal("example.com"))
			Expect(cmd.Param("path").Get()).To(Equal("/"))
			Expect(cmd.Flag("port").Get()).To(Equal(3001))
			Expect(cmd.Flag("username").Get()).To(Equal("wally"))
			Expect(cmd.Flag("password").Get()).To(Equal("App1etw0"))
			Expect(cmd.Flag("ssl").Get()).To(Equal(true))
			Expect(didRun).To(Equal(true))
		})

		It("should parse the sub command properly", func() {
			subCmd.DefineBoolFlag("power", false, "with power")
			cli.Start(strings.Split("cmd -Su=wally -p App1etw0 --port 3001 example.com / do --power something", " ")...)
			Expect(cmd.Parent().Param("host").Get()).To(Equal("example.com"))
			Expect(cmd.Parent().Param("path").Get()).To(Equal("/"))
			Expect(cmd.Parent().Flag("port").Get()).To(Equal(3001))
			Expect(cmd.Parent().Flag("username").Get()).To(Equal("wally"))
			Expect(cmd.Parent().Flag("password").Get()).To(Equal("App1etw0"))
			Expect(cmd.Parent().Flag("ssl").Get()).To(Equal(true))
			Expect(cmd.Flag("power").Get()).To(Equal(true))
			Expect(cmd.Param("action").Get()).To(Equal("something"))
			Expect(subDidRun).To(Equal(true))
		})

		Describe("version", func() {
			It("should not panic", func() {
				cli.Start("cmd", "--version")
			})
		})

		Describe("help", func() {
			It("should not panic", func() {
				cli.Start(strings.Split("cmd --help", " ")...)
				cli.Start(strings.Split("cmd host path do --help", " ")...)
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

		Context("invalid flags", func() {
			It("should panic", func() {
				Ω(func() { cli.Start("cmd", "--invalid") }).Should(Panic())
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

	Describe("remaining arguments", func() {
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
