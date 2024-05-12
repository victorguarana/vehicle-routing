package routes

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewMainRoute", func() {
	It("should return correct struct", func() {
		actualMainStop := &mainStop{}
		expectedMainRoute := mainRoute{
			mainStops: []*mainStop{actualMainStop},
		}
		receivedMainRoute := NewMainRoute(actualMainStop)
		Expect(receivedMainRoute).To(Equal(&expectedMainRoute))
	})
})

var _ = Describe("mainRoute{}", func() {
	Describe("Append", func() {
		var sut = mainRoute{
			mainStops: []*mainStop{},
		}

		Context("when all params are valid", func() {
			It("should append new car stop", func() {
				ms := NewMainStop(gps.Point{}).(*mainStop)
				sut.Append(ms)
				Expect(sut.mainStops).To(Equal([]*mainStop{ms}))
			})
		})
	})

	Describe("AtIndex", func() {
		var sut = mainRoute{
			mainStops: []*mainStop{
				{point: gps.Point{Latitude: 0, Longitude: 0}},
				{point: gps.Point{Latitude: 1, Longitude: 1}},
			},
		}

		Context("when index is valid", func() {
			It("should return car stop at index", func() {
				expectedStop := sut.mainStops[1]
				receivedStop := sut.AtIndex(1)
				Expect(receivedStop).To(Equal(expectedStop))
			})
		})

		Context("when index is invalid", func() {
			It("should return nil", func() {
				receivedStop := sut.AtIndex(2)
				Expect(receivedStop).To(BeNil())
			})
		})
	})

	Describe("First", func() {
		var firstMainStop = &mainStop{point: gps.Point{Latitude: 0}}
		var secondMainStop = &mainStop{point: gps.Point{Latitude: 1}}
		var sut = mainRoute{
			mainStops: []*mainStop{firstMainStop, secondMainStop},
		}

		It("should return first car stop", func() {
			receivedStop := sut.First()
			Expect(receivedStop).To(Equal(firstMainStop))
		})
	})

	Describe("Last", func() {
		var firstMainStop = &mainStop{point: gps.Point{Latitude: 0}}
		var secondMainStop = &mainStop{point: gps.Point{Latitude: 1}}
		var sut = mainRoute{
			mainStops: []*mainStop{firstMainStop, secondMainStop},
		}

		It("should return last car stop", func() {
			receivedStop := sut.Last()
			Expect(receivedStop).To(Equal(secondMainStop))
		})
	})

	Describe("RemoveMainStop", func() {
		var sut = mainRoute{
			mainStops: []*mainStop{
				{point: gps.Point{Latitude: 0, Longitude: 0}},
				{point: gps.Point{Latitude: 1, Longitude: 1}},
			},
		}

		Context("when index is valid", func() {
			It("should remove stop at index", func() {
				expectedStops := []*mainStop{
					{point: gps.Point{Latitude: 0, Longitude: 0}},
				}
				sut.RemoveMainStop(1)
				Expect(sut.mainStops).To(Equal(expectedStops))
			})
		})

		Context("when index is invalid", func() {
			It("should not change stops", func() {
				expectedStops := []*mainStop{
					{point: gps.Point{Latitude: 0, Longitude: 0}},
					{point: gps.Point{Latitude: 1, Longitude: 1}},
				}
				sut.RemoveMainStop(10)
				Expect(sut.mainStops).To(Equal(expectedStops))
			})
		})
	})
})
