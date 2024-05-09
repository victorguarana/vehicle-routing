package vehicles

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/victorguarana/go-vehicle-route/src/gps"
)

var _ = Describe("Land", func() {
	It("land drone", func() {
		sut := drone{
			isFlying:       true,
			totalStorage:   10,
			totalRange:     100,
			actualPosition: gps.Point{},
		}

		landingPoint := gps.Point{
			Latitude:    10,
			Longitude:   10,
			PackageSize: 1,
		}

		sut.Land(landingPoint)

		Expect(sut.isFlying).To(BeFalse())
		Expect(sut.actualPosition).To(Equal(landingPoint))
		Expect(sut.remaningRange).To(Equal(sut.totalRange))
		Expect(sut.remaningStorage).To(Equal(defaultDroneStorage))
	})
})

var _ = Describe("Move", func() {
	Context("when drone can move to next position", func() {
		It("move drone", func() {
			initialRange := 100.0
			p := gps.Point{
				Latitude:  10,
				Longitude: 10,
			}
			sut := drone{
				remaningRange: initialRange,
				actualPosition: gps.Point{
					Latitude:  0,
					Longitude: 0,
				},
			}
			distance := gps.DistanceBetweenPoints(p, sut.actualPosition)

			sut.Move(p)
			Expect(sut.actualPosition).To(Equal(p))
			Expect(sut.remaningRange).To(Equal(initialRange - distance))
		})
	})
})

var _ = Describe("Support", func() {
	Describe("single destination cases", func() {
		Context("when drone can support point with plenty of range and storage", func() {
			It("returns true", func() {
				destination := gps.Point{
					Latitude:    10,
					PackageSize: 1,
				}

				sut := drone{
					remaningStorage: 10,
					remaningRange:   100,
					actualPosition: gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				}
				Expect(sut.Support(destination)).To(BeTrue())
			})
		})

		Context("when drone can support point without plenty of range and storage", func() {
			It("returns true", func() {
				destination := gps.Point{
					Latitude:    10,
					PackageSize: 10,
				}

				sut := drone{
					remaningRange:   10,
					remaningStorage: 10,
					actualPosition: gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				}
				Expect(sut.Support(destination)).To(BeTrue())
			})
		})

		Context("when drone can not support point because of range", func() {
			It("returns false", func() {
				destination := gps.Point{
					Latitude:    1,
					PackageSize: 1,
				}

				sut := drone{
					remaningRange:   0,
					remaningStorage: 10,
					actualPosition: gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				}
				Expect(sut.Support(destination)).To(BeFalse())
			})
		})

		Context("when drone can not support point because of storage", func() {
			It("returns false", func() {
				destination := gps.Point{
					Latitude:    1,
					PackageSize: 1,
				}

				sut := drone{
					remaningRange:   10,
					remaningStorage: 0,
					actualPosition: gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				}
				Expect(sut.Support(destination)).To(BeFalse())
			})
		})
	})

	Describe("multi destinations cases", func() {
		Context("when drone can reach point with plenty", func() {
			It("returns true", func() {
				destination1 := gps.Point{Latitude: 10}
				destination2 := gps.Point{Latitude: 15}
				sut := drone{
					remaningRange: 100,
					actualPosition: gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				}
				Expect(sut.Support(destination1, destination2)).To(BeTrue())
			})
		})

		Context("when drone can reach point without plenty", func() {
			It("returns true", func() {
				destination1 := gps.Point{Latitude: 5}
				destination2 := gps.Point{Latitude: 10}
				sut := drone{
					remaningRange: 10,
					actualPosition: gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				}
				Expect(sut.Support(destination1, destination2)).To(BeTrue())
			})
		})

		Context("when drone can not reach first point", func() {
			It("returns false", func() {
				destination1 := gps.Point{Latitude: 5}
				destination2 := gps.Point{Latitude: 10}
				sut := drone{
					remaningRange: 0,
					actualPosition: gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				}
				Expect(sut.Support(destination1, destination2)).To(BeFalse())
			})
		})

		Context("when drone can not reach second point", func() {
			It("returns false", func() {
				destination1 := gps.Point{Latitude: 5}
				destination2 := gps.Point{Latitude: 10}
				sut := drone{
					remaningRange: 8,
					actualPosition: gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				}
				Expect(sut.Support(destination1, destination2)).To(BeFalse())
			})
		})
	})
})

var _ = Describe("ActualPosition", func() {
	Context("when drone has position", func() {
		It("returns drone position", func() {
			p := gps.Point{
				Latitude:  10,
				Longitude: 10,
			}
			sut := drone{
				actualPosition: p,
			}
			Expect(sut.ActualPosition()).To(Equal(p))
		})
	})
})
