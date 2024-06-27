package route

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/slc"

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

	Describe("InserAt", func() {
		var sut mainRoute
		var newStop = &mainStop{point: gps.Point{}}
		var stop1 = &mainStop{point: gps.Point{Latitude: 1}}
		var stop2 = &mainStop{point: gps.Point{Latitude: 2}}
		var stop3 = &mainStop{point: gps.Point{Latitude: 3}}

		BeforeEach(func() {
			sut = mainRoute{
				mainStops: []*mainStop{stop1, stop2, stop3},
			}
		})

		It("should insert car stop at index when it is inside length", func() {
			sut.InserAt(1, newStop)
			Expect(sut.mainStops).To(Equal([]*mainStop{stop1, newStop, stop2, stop3}))
		})

		It("should append car stop at index when it is equal length", func() {
			sut.InserAt(len(sut.mainStops), newStop)
			Expect(sut.mainStops).To(Equal([]*mainStop{stop1, stop2, stop3, newStop}))
		})

		It("should not insert car stop at index out of range", func() {
			sut.InserAt(4, newStop)
			Expect(sut.mainStops).To(Equal([]*mainStop{stop1, stop2, stop3}))
		})

		It("should not insert car stop at negative index", func() {
			sut.InserAt(-1, newStop)
			Expect(sut.mainStops).To(Equal([]*mainStop{stop1, stop2, stop3}))
		})
	})

	var _ = Describe("Iterator", func() {
		var mainStop1 = mainStop{point: gps.Point{Latitude: 1}}
		var mainStop2 = mainStop{point: gps.Point{Latitude: 2}}
		var iMainStop1 IMainStop = &mainStop1
		var iMainStop2 IMainStop = &mainStop2
		var iMainStops = []IMainStop{iMainStop1, iMainStop2}
		var sut = mainRoute{
			mainStops: []*mainStop{&mainStop1, &mainStop2},
		}

		It("should return iterator", func() {
			expectedIterator := slc.NewIterator(iMainStops)
			receivedIterator := sut.Iterator()
			Expect(receivedIterator).To(Equal(expectedIterator))
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

	Describe("Length", func() {
		var sut = mainRoute{
			mainStops: []*mainStop{
				{point: gps.Point{Latitude: 1}},
				{point: gps.Point{Latitude: 2}},
			},
		}

		It("should return length of car stops", func() {
			Expect(sut.Length()).To(Equal(2))
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
				}
				sut.RemoveMainStop(10)
				Expect(sut.mainStops).To(Equal(expectedStops))
			})
		})
	})
})
