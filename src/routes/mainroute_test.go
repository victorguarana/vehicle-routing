package routes

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Append", func() {
	var sut = mainRoute{
		mainStops: []*mainStop{},
	}

	Context("when all params are valid", func() {
		It("appends new car stop", func() {
			ms := NewMainStop(gps.Point{}).(*mainStop)
			sut.Append(ms)
			Expect(sut.mainStops).To(Equal([]*mainStop{ms}))
		})
	})
})

var _ = Describe("AtIndex", func() {
	var sut = mainRoute{
		mainStops: []*mainStop{
			{point: gps.Point{Latitude: 0, Longitude: 0}},
			{point: gps.Point{Latitude: 1, Longitude: 1}},
		},
	}

	Context("when index is valid", func() {
		It("returns car stop at index", func() {
			expectedStop := sut.mainStops[1]
			receivedStop, receivedErr := sut.AtIndex(1)
			Expect(receivedErr).NotTo(HaveOccurred())
			Expect(receivedStop).To(Equal(expectedStop))
		})
	})

	Context("when index is invalid", func() {
		It("returns index out of range error", func() {
			receivedStop, receivedErr := sut.AtIndex(2)
			Expect(receivedErr).To(MatchError(ErrIndexOutOfRange))
			Expect(receivedStop).To(BeNil())
		})
	})
})

var _ = Describe("RemoveMainStop", func() {
	var sut = mainRoute{
		mainStops: []*mainStop{
			{point: gps.Point{Latitude: 0, Longitude: 0}},
			{point: gps.Point{Latitude: 1, Longitude: 1}},
		},
	}

	Context("when index is valid", func() {
		It("removes car stop at index", func() {
			expectedStops := []*mainStop{
				{point: gps.Point{Latitude: 0, Longitude: 0}},
			}
			Expect(sut.RemoveMainStop(1)).To(Succeed())
			Expect(sut.mainStops).To(Equal(expectedStops))
		})
	})

	Context("when index is invalid", func() {
		It("returns index out of range error", func() {
			Expect(sut.RemoveMainStop(10)).To(MatchError(ErrIndexOutOfRange))
		})
	})
})
