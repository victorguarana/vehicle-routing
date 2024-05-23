package gps

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGps(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gps Suite")
}
