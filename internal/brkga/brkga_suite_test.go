package brkga

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBrkga(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Brkga Suite")
}
