package gps

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClosestPoint", func() {
	Context("when cadidates latitudes are equal", func() {
		var originPoint = Point{Latitude: 0, Longitude: 0}
		var candidatePoints = []Point{
			{Latitude: 1, Longitude: 2},
			{Latitude: 1, Longitude: 1},
			{Latitude: 1, Longitude: 3},
		}

		It("should return closest point", func() {
			expectedPoint := candidatePoints[1]
			receivedPoint := ClosestPoint(originPoint, candidatePoints)
			Expect(receivedPoint).To(Equal(expectedPoint))
		})
	})

	Context("when candidates longitudes are equal", func() {
		var originPoint = Point{Latitude: 0, Longitude: 0}
		var candidatePoints = []Point{
			{Latitude: 2, Longitude: 1},
			{Latitude: 1, Longitude: 1},
			{Latitude: 3, Longitude: 1},
		}

		It("should return closest point", func() {
			expectedPoint := candidatePoints[1]
			receivedPoint := ClosestPoint(originPoint, candidatePoints)
			Expect(receivedPoint).To(Equal(expectedPoint))
		})
	})

	Context("when there are no candidate points", func() {
		It("should return empty", func() {
			originPoint := Point{Latitude: 10, Longitude: 10}
			receivedPoint := ClosestPoint(originPoint, []Point{})
			Expect(receivedPoint).To(Equal(Point{}))
		})
	})
})

var _ = Describe("DistanceBetweenPoints", func() {
	Context("when points are positive", func() {
		var points = []Point{
			{Latitude: 10, Longitude: 10},
			{Latitude: 20, Longitude: 20},
			{Latitude: 30, Longitude: 30},
		}
		It("should return the sum of the distances between points", func() {
			Expect(DistanceBetweenPoints(points...)).To(Equal(40.0))
		})
	})
	Context("when points are mirrored", func() {
		var points = []Point{
			{Latitude: 10, Longitude: 20},
			{Latitude: 20, Longitude: 10},
		}
		It("should return the sum of the distances between points", func() {
			Expect(DistanceBetweenPoints(points...)).To(Equal(20.0))
		})
	})
	Context("when points are oposite", func() {
		var points = []Point{
			{Latitude: 10, Longitude: 10},
			{Latitude: -10, Longitude: -10},
		}
		It("should return the sum of the distances between points", func() {
			Expect(DistanceBetweenPoints(points...)).To(Equal(40.0))
		})
	})
	Context("when points turn back", func() {
		var points = []Point{
			{Latitude: 0, Longitude: 0},
			{Latitude: 10, Longitude: 10},
			{Latitude: 0, Longitude: 0},
		}
		It("should return the sum of the distances between points", func() {
			Expect(DistanceBetweenPoints(points...)).To(Equal(40.0))
		})
	})
})
