package route

import (
	"io"
	"log"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRoutes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Routes Suite")
}

var _ = BeforeSuite(func() {
	log.SetOutput(io.Discard)
})

var _ = AfterSuite(func() {
	log.SetOutput(os.Stderr)
})
