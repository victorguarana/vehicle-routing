package slc_test

import (
	"io"
	"log"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSlc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Slc Suite")
}

var _ = BeforeSuite(func() {
	log.SetOutput(io.Discard)
})

var _ = AfterSuite(func() {
	log.SetOutput(os.Stderr)
})
