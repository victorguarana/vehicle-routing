package routes

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewSubRoute", Ordered, func() {
	var startingPoint = &mainStop{point: gps.Point{}}
	var returningPoint = &mainStop{point: gps.Point{}}

	It("takes off drone and returns correct struct", func() {
		expectedSubRoute := &subRoute{
			startingPoint:  startingPoint,
			returningPoint: returningPoint,
			stops:          []*subStop{},
		}
		receivedSubRoute := NewSubRoute(startingPoint, returningPoint)
		Expect(receivedSubRoute).To(Equal(expectedSubRoute))
		Expect(startingPoint.subRoutes).To(ContainElement(receivedSubRoute))
		Expect(returningPoint.subRoutes).To(ContainElement(receivedSubRoute))
	})
})

var _ = Describe("Append", Ordered, func() {
	var validPoint = gps.Point{}
	var sut = subRoute{
		startingPoint:  &mainStop{point: validPoint},
		returningPoint: &mainStop{point: validPoint},
	}

	It("appends substop to subRoute", func() {
		appendedPoint := &subStop{point: validPoint}
		sut.Append(appendedPoint)
		Expect(sut.stops).To(Equal([]*subStop{appendedPoint}))
	})
})

var _ = Describe("Land", Ordered, func() {
	var sut = subRoute{
		startingPoint:  &mainStop{point: gps.Point{}},
		returningPoint: &mainStop{point: gps.Point{}},
	}

	It("lands drone and sets landing point", func() {
		returningPoint := &mainStop{point: gps.Point{}}
		sut.Return(returningPoint)
		Expect(sut.returningPoint).To(Equal(returningPoint))
	})
})
