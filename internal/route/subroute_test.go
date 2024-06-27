package route

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/slc"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewSubRoute", func() {
	var startingStop = &mainStop{point: gps.Point{}}

	It("takes off drone and returns correct struct", func() {
		expectedSubRoute := &subRoute{
			startingStop: startingStop,
			stops:        []*subStop{},
		}
		receivedSubRoute := NewSubRoute(startingStop)
		Expect(receivedSubRoute).To(Equal(expectedSubRoute))
		Expect(startingStop.startingSubRoutes).To(ContainElement(receivedSubRoute))
	})
})

var _ = Describe("subRoute{}", func() {
	var _ = Describe("Append", func() {
		var validPoint = gps.Point{}
		var sut = subRoute{
			startingStop:  &mainStop{point: validPoint},
			returningStop: &mainStop{point: validPoint},
		}

		It("should append sub stop to sub route", func() {
			appendedPoint := &subStop{point: validPoint}
			sut.Append(appendedPoint)
			Expect(sut.stops).To(Equal([]*subStop{appendedPoint}))
		})
	})

	var _ = Describe("First", func() {
		var firstSubStop = &subStop{point: gps.Point{Latitude: 1}}
		var secondSubStop = &subStop{point: gps.Point{Latitude: 2}}
		var sut = subRoute{
			stops: []*subStop{firstSubStop, secondSubStop},
		}

		It("should return first stop", func() {
			Expect(sut.First()).To(Equal(firstSubStop))
		})
	})

	var _ = Describe("InsertAt", func() {
		var sut subRoute
		var newStop = &subStop{gps.Point{}}
		var stop1 = &subStop{point: gps.Point{Latitude: 1}}
		var stop2 = &subStop{point: gps.Point{Latitude: 2}}
		var stop3 = &subStop{point: gps.Point{Latitude: 3}}
		BeforeEach(func() {
			sut = subRoute{
				stops: []*subStop{stop1, stop2, stop3},
			}
		})

		It("should insert sub stop at index when it is inside length", func() {
			expectedStops := []*subStop{stop1, newStop, stop2, stop3}
			sut.InsertAt(1, newStop)
			Expect(sut.stops).To(Equal(expectedStops))
		})

		It("should append sub stop when index is equal length", func() {
			sut.InsertAt(len(sut.stops), newStop)
			Expect(sut.stops).To(Equal([]*subStop{stop1, stop2, stop3, newStop}))
		})

		It("should not insert sub stop at index out of range", func() {
			sut.InsertAt(4, newStop)
			Expect(sut.stops).To(Equal([]*subStop{stop1, stop2, stop3}))
		})

		It("should not insert sub stop at negative index", func() {
			sut.InsertAt(-1, newStop)
			Expect(sut.stops).To(Equal([]*subStop{stop1, stop2, stop3}))
		})
	})

	var _ = Describe("Iterator", func() {
		var subStop1 = subStop{point: gps.Point{Latitude: 1}}
		var subStop2 = subStop{point: gps.Point{Latitude: 2}}
		var iSubStop1 ISubStop = &subStop1
		var iSubStop2 ISubStop = &subStop2
		var iSubStops = []ISubStop{iSubStop1, iSubStop2}
		var sut = subRoute{
			stops: []*subStop{&subStop1, &subStop2},
		}

		It("should return iterator", func() {
			expectedIterator := slc.NewIterator(iSubStops)
			receivedIterator := sut.Iterator()
			Expect(receivedIterator).To(Equal(expectedIterator))
		})
	})

	var _ = Describe("Last", func() {
		var firstSubStop = &subStop{point: gps.Point{Latitude: 1}}
		var secondSubStop = &subStop{point: gps.Point{Latitude: 2}}
		var sut = subRoute{
			stops: []*subStop{firstSubStop, secondSubStop},
		}

		It("should return last stop", func() {
			Expect(sut.Last()).To(Equal(secondSubStop))
		})
	})

	var _ = Describe("Length", func() {
		var sut = subRoute{
			stops: []*subStop{
				{point: gps.Point{Latitude: 1}},
				{point: gps.Point{Latitude: 2}},
			},
		}

		It("should return length of stops", func() {
			Expect(sut.Length()).To(Equal(2))
		})
	})

	var _ = Describe("Return", func() {
		var sut = subRoute{
			startingStop:  &mainStop{point: gps.Point{}},
			returningStop: &mainStop{point: gps.Point{}},
		}

		It("should set returning point", func() {
			returningStop := &mainStop{point: gps.Point{}}
			sut.Return(returningStop)
			Expect(sut.returningStop).To(Equal(returningStop))
		})
	})

	var _ = Describe("ReturningStop", func() {
		var returningStop = &mainStop{point: gps.Point{}}
		var sut = subRoute{
			returningStop: returningStop,
		}

		It("should return returning point", func() {
			Expect(sut.ReturningStop()).To(Equal(returningStop))
		})
	})

	var _ = Describe("StartingStop", func() {
		var startingStop = &mainStop{point: gps.Point{}}
		var sut = subRoute{
			startingStop: startingStop,
		}

		It("should return starting point", func() {
			Expect(sut.StartingStop()).To(Equal(startingStop))
		})
	})

	var _ = Describe("RemoveSubStop", func() {
		var sut subRoute
		var stop1 = &subStop{point: gps.Point{Latitude: 1}}
		var stop2 = &subStop{point: gps.Point{Latitude: 2}}
		var stop3 = &subStop{point: gps.Point{Latitude: 3}}
		BeforeEach(func() {
			sut = subRoute{
				stops: []*subStop{stop1, stop2, stop3},
			}
		})

		It("should remove sub stop at index when it is inside length", func() {
			expectedStops := []*subStop{stop1, stop3}
			sut.RemoveSubStop(1)
			Expect(sut.stops).To(Equal(expectedStops))
		})

		It("should not remove sub stop at index out of range", func() {
			sut.RemoveSubStop(3)
			Expect(sut.stops).To(Equal([]*subStop{stop1, stop2, stop3}))
		})

		It("should not remove sub stop at negative index", func() {
			sut.RemoveSubStop(-1)
			Expect(sut.stops).To(Equal([]*subStop{stop1, stop2, stop3}))
		})
	})
})
