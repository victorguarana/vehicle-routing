package measure_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMeasure(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Measure Suite")
}
