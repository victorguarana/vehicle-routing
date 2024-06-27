package csp

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCsp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Csp Suite")
}
