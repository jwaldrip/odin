package gbytes_test

import (
	. "github.com/jwaldrip/odin/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/jwaldrip/odin/Godeps/_workspace/src/github.com/onsi/gomega"

	"testing"
)

func TestGbytes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gbytes Suite")
}
