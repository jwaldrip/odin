package cli_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/jwaldrip/odin/cli"
)

var _ = Describe("Output", func() {

	var cli *CLI
	var cmd Command
	var didRun bool
	var out *bytes.Buffer
	var err *bytes.Buffer
	var runFn = func(c Command) {}

	BeforeEach(func() {
		out = bytes.NewBufferString("")
		err = bytes.NewBufferString("")
		didRun = false
		cli = New("v1.0.0", "sample description", func(c Command) {
			cmd = c
			didRun = true
			runFn(c)
		})
		cli.SetStdErr(err)
		cli.SetStdOut(out)
		cli.ErrorHandling = PanicOnError
	})

	Describe("Custom Outputs", func() {
		It("Should allow for custom standard output", func() {
			customOutput := bytes.NewBufferString("")
			runFn = func(c Command) {
				c.Println("standard")
			}
			cli.SetStdOut(customOutput)
			cli.Start("cmd")
			Expect(customOutput.String()).To(Equal("standard\n"))
		})

		It("Should allow for custom err output", func() {
			customErr := bytes.NewBufferString("")
			runFn = func(c Command) {
				c.ErrPrintln("error")
			}
			cli.SetStdErr(customErr)
			cli.Start("cmd")
			Expect(customErr.String()).To(Equal("error\n"))
		})
	})

	Describe("ErrPrint", func() {
		It("Sould Print to the specified output", func() {
			runFn = func(c Command) {
				c.ErrPrint("Hello", "Foo")
			}
			cli.Start("cmd")
			Expect(err.String()).To(Equal("HelloFoo"))
		})
	})
	Describe("ErrPrintf", func() {
		It("Sould Print to the specified output", func() {
			runFn = func(c Command) {
				c.ErrPrintf("%s-%s", "Hello", "Foo")
			}
			cli.Start("cmd")
			Expect(err.String()).To(Equal("Hello-Foo"))
		})
	})
	Describe("ErrPrintln", func() {
		It("Sould Print to the specified output", func() {
			runFn = func(c Command) {
				c.ErrPrintln("Hello", "Foo")
			}
			cli.Start("cmd")
			Expect(err.String()).To(Equal("Hello Foo\n"))
		})
	})

	Describe("Print", func() {
		It("Sould Print to the specified output", func() {
			runFn = func(c Command) {
				c.Print("Hello", "Foo")
			}
			cli.Start("cmd")
			Expect(out.String()).To(Equal("HelloFoo"))
		})
	})
	Describe("Printf", func() {
		It("Sould Print to the specified output", func() {
			runFn = func(c Command) {
				c.Printf("%s-%s", "Hello", "Foo")
			}
			cli.Start("cmd")
			Expect(out.String()).To(Equal("Hello-Foo"))
		})
	})
	Describe("Println", func() {
		It("Sould Print to the specified output", func() {
			runFn = func(c Command) {
				c.Println("Hello", "Foo")
			}
			cli.Start("cmd")
			Expect(out.String()).To(Equal("Hello Foo\n"))
		})
	})

})
