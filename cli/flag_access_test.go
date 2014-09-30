package cli_test

import (
	. "github.com/jwaldrip/odin/cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flag Access", func() {

	var cli *CLI
	var cmd Command

	BeforeEach(func() {
		runFn := func(c Command) {
			cmd = c
		}
		cli = New("v1.0.0", "sample description", runFn)
		cli.ErrorHandling = PanicOnError
		cli.Mute()

		cli.DefineBoolFlag("foo", false, "is foo")
		cli.AliasFlag('o', "foo")
		cli.DefineStringFlag("bar", "", "what bar are you at?")
		cli.AliasFlag('r', "bar")
		cli.DefineBoolFlag("baz", true, "is baz")
		cli.AliasFlag('z', "baz")
		cli.Start("cmd")
	})

	Describe("Flag(string) Value", func() {
		It("should return a flag's value", func() {
			Expect(cmd.Flag("foo").Get()).To(Equal(false))
		})

		Context("when a flag is not defined", func() {
			It("should panic", func() {
				Expect(func() { cmd.Flag("undefined") }).Should(Panic())
			})
		})
	})

	Describe("Flags() map[string]value", func() {
		It("should return each flags value", func() {
			for k, v := range cmd.Flags() {
				Expect(v).To(Equal(cmd.Flag(k)))
			}
		})
	})

})
