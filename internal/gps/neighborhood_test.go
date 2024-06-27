package gps

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MapNeighborhood", func() {
	Context("when there are no points", func() {
		It("should return empty map", func() {
			points := []Point{}
			maxDistance := 10.0
			Expect(MapNeighborhood(points, maxDistance)).To(BeEmpty())
		})
	})

	Context("when there are points", func() {
		var point0 = Point{Latitude: 0}
		var point1 = Point{Latitude: 10}
		var point2 = Point{Latitude: 20}
		var point3 = Point{Latitude: 30}
		var point4 = Point{Latitude: 40}
		var points = []Point{point0, point1, point2, point3, point4}

		It("should return a map with nearby points", func() {
			maxDistance := 25.0
			expectedMap := Neighborhood{
				point0: {point1, point2},
				point1: {point0, point2, point3},
				point2: {point0, point1, point3, point4},
				point3: {point1, point2, point4},
				point4: {point2, point3},
			}
			receivedMap := MapNeighborhood(points, maxDistance)
			Expect(receivedMap).To(Equal(expectedMap))
		})
	})
})

var _ = Describe("RemovePointFromNearbyMap", func() {
	Context("when point is in the map", func() {
		var point0 = Point{Latitude: 0}
		var point1 = Point{Latitude: 10}
		var point2 = Point{Latitude: 20}
		var point3 = Point{Latitude: 30}
		var point4 = Point{Latitude: 40}
		var nearbyPointsMap = Neighborhood{
			point0: {point1, point2},
			point1: {point0, point2, point3},
			point2: {point0, point1, point3, point4},
			point3: {point1, point2, point4},
			point4: {point2, point3},
		}

		It("should remove the point from the map", func() {
			expectedMap := Neighborhood{
				point0: {point2},
				point2: {point0, point3, point4},
				point3: {point2, point4},
				point4: {point2, point3},
			}
			RemovePointFromNearbyMap(point1, nearbyPointsMap)
			Expect(nearbyPointsMap).To(Equal(expectedMap))
		})
	})

	Context("when point is not in the map", func() {
		var pointToBeRemoved = Point{Latitude: 30}
		var point0 = Point{Latitude: 0}
		var point1 = Point{Latitude: 10}
		var point2 = Point{Latitude: 20}
		var nearbyPointsMap = Neighborhood{
			point0: {point1, point2},
			point1: {point0, point2},
			point2: {point0, point1},
		}
		var expectedMap = Neighborhood{
			point0: {point1, point2},
			point1: {point0, point2},
			point2: {point0, point1},
		}

		It("should not change the map", func() {
			RemovePointFromNearbyMap(pointToBeRemoved, nearbyPointsMap)
			Expect(nearbyPointsMap).To(Equal(expectedMap))
		})
	})
})

var _ = Describe("ClosestPointWithMostNeighbors", func() {
	Context("when there are no points", func() {
		It("should return empty point", func() {
			var initialPoint = Point{}
			nearbyPoints := Neighborhood{}
			receivedPoint := ClosestPointWithMostNeighbors(initialPoint, nearbyPoints)
			Expect(receivedPoint).To(BeZero())
		})
	})

	Context("when there are points", func() {
		var initialPoint = Point{}
		var point1 = Point{Latitude: 10}
		var point2 = Point{Latitude: 20}
		var point3 = Point{Latitude: 30}
		var point4 = Point{Latitude: 40}
		var point5 = Point{Latitude: 50}
		var nearbyPoints = Neighborhood{
			point1: {point2},
			point2: {point1, point3},
			point3: {point2, point4},
			point4: {point3, point5},
			point5: {point4},
		}

		It("should return the closest point with most neighbors", func() {
			receivedPoint := ClosestPointWithMostNeighbors(initialPoint, nearbyPoints)
			Expect(receivedPoint).To(Equal(point2))
		})
	})

	Context("when points does not have any neighbor", func() {
		var initialPoint = Point{}
		var point1 = Point{Latitude: 10}
		var point2 = Point{Latitude: 20}
		var nearbyPoints = Neighborhood{
			point1: {},
			point2: {},
		}

		It("should return the closest point", func() {
			receivedPoint := ClosestPointWithMostNeighbors(initialPoint, nearbyPoints)
			Expect(receivedPoint).To(Equal(point1))
		})
	})
})
