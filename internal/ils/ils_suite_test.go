package ils_test

import (
	"io"
	"log"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIls(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ils Suite")
}

var _ = BeforeSuite(func() {
	log.SetOutput(io.Discard)
})

var _ = AfterSuite(func() {
	log.SetOutput(os.Stderr)
})
