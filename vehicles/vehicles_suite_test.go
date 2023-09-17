package vehicles_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestVehicles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vehicles Suite")
}
