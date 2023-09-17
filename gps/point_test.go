package gps

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("DistanceBetweenPoints", func(p1, p2 Point, expectedResponse float64) {
	Expect(DistanceBetweenPoints(p1, p2)).To(Equal(expectedResponse))
},
	Entry("when points are positive",
		Point{Latitude: 10, Longitude: 10},
		Point{Latitude: 20, Longitude: 20},
		20.0,
	),
	Entry("when points are mirrored",
		Point{Latitude: 10, Longitude: 20},
		Point{Latitude: 20, Longitude: 10},
		20.0,
	),
	Entry("when points are oposite",
		Point{Latitude: 10, Longitude: 10},
		Point{Latitude: -10, Longitude: -10},
		40.0,
	),
)
