package cli_test

import (
	"time"

	. "github.com/jwaldrip/odin/cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flag definitions", func() {

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

	Describe("aliasing", func() {
		BeforeEach(func() {
			cli.DefineBoolFlag("valid", false, "valid flag")
		})

		Context("with a valid flag", func() {
			It("should not panic", func() {
				Expect(func() { cli.AliasFlag('a', "valid") }).ShouldNot(Panic())
			})
		})

		Context("without a valid flag", func() {
			It("should panic", func() {
				Expect(func() { cli.AliasFlag('a', "notvalid") }).Should(Panic())
			})
		})
	})

	Describe("type definitions", func() {

		It("should panic if a flag is redefined", func() {
			cli.DefineBoolFlag("foo", false, "")
			Expect(func() { cli.DefineBoolFlag("foo", false, "") }).Should(Panic())
		})

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

		Describe("Float64Flag", func() {

			It("should set and parse", func() {
				cli.DefineFloat64Flag("float64flag", 10.1, "desc")
				cli.Start("cmd", "--float64flag=9.9")
				Expect(cmd.Flag("float64flag").Get()).To(Equal(9.9))
			})

			It("should set and parse from var", func() {
				var val float64
				cli.DefineFloat64FlagVar(&val, "float64flag", 10.10, "desc")
				cli.DefineFloat64FlagVar(&val, "float64flag2", 9.9, "desc")
				cli.Start("cmd", "--float64flag=8.8")
				Expect(cmd.Flag("float64flag2").Get()).To(Equal(8.8))
			})
		})

		Describe("Int64Flag", func() {

			It("should set and parse", func() {
				cli.DefineInt64Flag("int64flag", 10, "desc")
				cli.Start("cmd", "--int64flag=9")
				Expect(cmd.Flag("int64flag").Get()).To(Equal(int64(9)))
			})

			It("should set and parse from var", func() {
				var val int64
				cli.DefineInt64FlagVar(&val, "int64flag", 10, "desc")
				cli.DefineInt64FlagVar(&val, "int64flag2", 9, "desc")
				cli.Start("cmd", "--int64flag=8")
				Expect(cmd.Flag("int64flag2").Get()).To(Equal(int64(8)))
			})
		})

		Describe("IntFlag", func() {

			It("should set and parse", func() {
				cli.DefineIntFlag("intflag", 10, "desc")
				cli.Start("cmd", "--intflag=9")
				Expect(cmd.Flag("intflag").Get()).To(Equal(int(9)))
			})

			It("should set and parse from var", func() {
				var val int
				cli.DefineIntFlagVar(&val, "intflag", 10, "desc")
				cli.DefineIntFlagVar(&val, "intflag2", 9, "desc")
				cli.Start("cmd", "--intflag=8")
				Expect(cmd.Flag("intflag2").Get()).To(Equal(int(8)))
			})
		})

		Describe("StringFlag", func() {

			It("should set and parse", func() {
				cli.DefineStringFlag("stringflag", "foo", "desc")
				cli.Start("cmd", "--stringflag=bar")
				Expect(cmd.Flag("stringflag").Get()).To(Equal("bar"))
			})

			It("should set and parse from var", func() {
				var val string
				cli.DefineStringFlagVar(&val, "stringflag", "foo", "desc")
				cli.DefineStringFlagVar(&val, "stringflag2", "bar", "desc")
				cli.Start("cmd", "--stringflag=baz")
				Expect(cmd.Flag("stringflag2").Get()).To(Equal("baz"))
			})
		})

		Describe("Uint64Flag", func() {

			It("should set and parse", func() {
				cli.DefineUint64Flag("uint64flag", uint64(10), "desc")
				cli.Start("cmd", "--uint64flag=9")
				Expect(cmd.Flag("uint64flag").Get()).To(Equal(uint64(9)))
			})

			It("should set and parse from var", func() {
				var val uint64
				cli.DefineUint64FlagVar(&val, "uint64flag", uint64(10), "desc")
				cli.DefineUint64FlagVar(&val, "uint64flag2", uint64(9), "desc")
				cli.Start("cmd", "--uint64flag=8")
				Expect(cmd.Flag("uint64flag2").Get()).To(Equal(uint64(8)))
			})
		})

		Describe("UintFlag", func() {

			It("should set and parse", func() {
				cli.DefineUintFlag("uintflag", uint(10), "desc")
				cli.Start("cmd", "--uintflag=9")
				Expect(cmd.Flag("uintflag").Get()).To(Equal(uint(9)))
			})

			It("should set and parse from var", func() {
				var val uint
				cli.DefineUintFlagVar(&val, "uintflag", uint(10), "desc")
				cli.DefineUintFlagVar(&val, "uintflag2", uint(9), "desc")
				cli.Start("cmd", "--uintflag=8")
				Expect(cmd.Flag("uintflag2").Get()).To(Equal(uint(8)))
			})
		})

	})

})
