package odin_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOdin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Odin Suite")
}
