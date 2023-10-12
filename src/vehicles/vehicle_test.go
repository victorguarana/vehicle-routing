package vehicles

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/victorguarana/go-vehicle-route/src/gps"
)

var _ = Describe("ActualPosition", func() {
	Context("when car can be created", func() {
		It("create with correct params", func() {
			expectedPoint := &gps.Point{
				Latitude:  10,
				Longitude: 10,
			}
			sut := vehicle{
				actualPosition: expectedPoint,
			}

			Expect(sut.ActualPosition()).To(Equal(expectedPoint))
		})
	})
})
