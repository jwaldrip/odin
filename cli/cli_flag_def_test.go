package cli_test

import (
	"time"

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

})
