package decoder

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDecoder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Decoder Suite")
}
