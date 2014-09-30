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

	Describe("Flags() ValueMap", func() {
		It("should return each flags value", func() {
			for k, v := range cmd.Flags() {
				Expect(v).To(Equal(cmd.Flag(k)))
			}
		})

		It("Should be a value map", func() {
			Expect(cmd.Flags().Keys()).To(ContainElement("foo"))
			Expect(cmd.Flags().Keys()).To(ContainElement("bar"))
			Expect(cmd.Flags().Keys()).To(ContainElement("baz"))
			Expect(cmd.Flags().Values().GetAll()).To(ContainElement(false))
			Expect(cmd.Flags().Values().GetAll()).To(ContainElement(""))
			Expect(cmd.Flags().Values().GetAll()).To(ContainElement(true))
			Expect(cmd.Flags().Values().Strings()).To(ContainElement("false"))
			Expect(cmd.Flags().Values().Strings()).To(ContainElement(""))
			Expect(cmd.Flags().Values().Strings()).To(ContainElement("true"))
		})
	})

})
