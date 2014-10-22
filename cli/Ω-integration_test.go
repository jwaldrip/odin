package cli_test

import (
	"strings"

	. "github.com/jwaldrip/odin/cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CLI Integration Test", func() {

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

	Describe("complex cli", func() {

		var subDidRun bool
		var subCmd *SubCommand

		BeforeEach(func() {
			cli.DefineParams("host", "path")
			cli.DefineBoolFlag("ssl", false, "do it over ssl")
			cli.AliasFlag('S', "ssl")
			cli.AliasFlag('s', "ssl")
			cli.DefineStringFlag("username", "", "the username")
			cli.AliasFlag('u', "username")
			cli.DefineStringFlag("password", "", "the password")
			cli.AliasFlag('p', "password")
			cli.DefineIntFlag("port", 80, "the port")
			cli.AliasFlag('P', "port")
			cli.DefineBoolFlag("keepopen", false, "keep the connection open")
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

		It("should parse flags that occur after positional params and without sub-command", func() {
			cli.Start(strings.Split("cmd example.com / --ssl", " ")...)
			Expect(cmd.Param("host").Get()).To(Equal("example.com"))
			Expect(cmd.Param("path").Get()).To(Equal("/"))
			Expect(cmd.Flag("ssl").Get()).To(Equal(true))
			Expect(cmd.Args()).To(BeZero())
			Expect(didRun).To(Equal(true))
		})

		It("should parse flags that occur after positional params and with sub-command", func() {
			subCmd.DefineBoolFlag("power", false, "with power")
			cli.Start(strings.Split("cmd example.com / do something --power", " ")...)
			Expect(cmd.Parent().Param("host").Get()).To(Equal("example.com"))
			Expect(cmd.Parent().Param("path").Get()).To(Equal("/"))
			Expect(cmd.Flag("power").Get()).To(Equal(true))
			Expect(cmd.Param("action").Get()).To(Equal("something"))
			Expect(cmd.Args()).To(BeZero())
			Expect(subDidRun).To(Equal(true))
		})

	})

})
