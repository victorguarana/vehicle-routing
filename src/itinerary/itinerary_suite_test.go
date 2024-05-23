package itinerary

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestItinerary(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Itinerary Suite")
}
