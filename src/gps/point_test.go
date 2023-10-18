package gps

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("DistanceBetweenPoints", func(points []*Point, expectedResponse float64) {
	Expect(DistanceBetweenPoints(points...)).To(Equal(expectedResponse))
},
	Entry("when points are positive",
		[]*Point{
			{Latitude: 10, Longitude: 10},
			{Latitude: 20, Longitude: 20},
			{Latitude: 30, Longitude: 30},
		},
		40.0,
	),
	Entry("when points are mirrored",
		[]*Point{
			{Latitude: 10, Longitude: 20},
			{Latitude: 20, Longitude: 10},
		},
		20.0,
	),
	Entry("when points are oposite",
		[]*Point{
			{Latitude: 10, Longitude: 10},
			{Latitude: -10, Longitude: -10},
		},
		40.0,
	),
	Entry("when points turn back",
		[]*Point{
			{Latitude: 0, Longitude: 0},
			{Latitude: 10, Longitude: 10},
			{Latitude: 0, Longitude: 0},
		},
		40.0,
	),
)
