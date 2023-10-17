package greedy_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGreedy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Greedy Suite")
}
