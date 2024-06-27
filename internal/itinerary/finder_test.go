package itinerary

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/route"
	routemock "github.com/victorguarana/vehicle-routing/internal/route/mock"
	"github.com/victorguarana/vehicle-routing/internal/slc"

	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("finder{}", func() {
	Describe("FindWorstDroneStop", func() {
		Context("when there are two sub itineraries", func() {
			var mockedCtrl *gomock.Controller
			var mockedFlight1 *routemock.MockISubRoute
			var mockedFlight2 *routemock.MockISubRoute
			var mockedDroneStop1 *routemock.MockIMainStop
			var mockedDroneStop2 *routemock.MockIMainStop
			var mockedStartingPoint *routemock.MockIMainStop
			var mockedReturningPoint *routemock.MockIMainStop
			var sut Finder

			BeforeEach(func() {
				mockedCtrl = gomock.NewController(GinkgoT())
				mockedFlight1 = routemock.NewMockISubRoute(mockedCtrl)
				mockedFlight2 = routemock.NewMockISubRoute(mockedCtrl)
				mockedDroneStop1 = routemock.NewMockIMainStop(mockedCtrl)
				mockedDroneStop2 = routemock.NewMockIMainStop(mockedCtrl)
				mockedStartingPoint = routemock.NewMockIMainStop(mockedCtrl)
				mockedReturningPoint = routemock.NewMockIMainStop(mockedCtrl)

				sut = &finder{&info{
					&itinerary{
						completedSubItineraryList: []SubItinerary{
							{Drone: nil, Flight: mockedFlight1},
							{Drone: nil, Flight: mockedFlight2},
						},
					},
				}}
			})

			It("should return the worst drone stop", func() {
				mockedFlight1.EXPECT().First().Return(mockedDroneStop1).Times(3)
				mockedFlight1.EXPECT().Last().Return(mockedDroneStop1)
				mockedFlight1.EXPECT().StartingStop().Return(mockedStartingPoint)
				mockedFlight1.EXPECT().ReturningStop().Return(mockedReturningPoint)
				mockedDroneStop1.EXPECT().Point().Return(gps.Point{Latitude: 2, Longitude: 10})

				mockedFlight2.EXPECT().First().Return(mockedDroneStop2).Times(3)
				mockedFlight2.EXPECT().Last().Return(mockedDroneStop2)
				mockedFlight2.EXPECT().StartingStop().Return(mockedStartingPoint)
				mockedFlight2.EXPECT().ReturningStop().Return(mockedReturningPoint)
				mockedDroneStop2.EXPECT().Point().Return(gps.Point{Latitude: 2, Longitude: 15})

				mockedStartingPoint.EXPECT().Point().Return(gps.Point{Latitude: 1}).Times(2)
				mockedReturningPoint.EXPECT().Point().Return(gps.Point{Latitude: 3}).Times(2)

				expectedDroneStopCost := DroneStopCost{
					Stop:   mockedDroneStop2,
					Flight: mockedFlight2,
					Index:  0,
					cost:   30,
				}

				receivedDroneStopCost := sut.FindWorstDroneStop()
				Expect(receivedDroneStopCost).To(Equal(expectedDroneStopCost))
			})
		})
	})

	Describe("FindWorstSwappableCarStopsOrdered", func() {
		Context("when there are car stops", func() {
			var mockedCtrl *gomock.Controller
			var mockedCarStop1 *routemock.MockIMainStop
			var mockedCarStop2 *routemock.MockIMainStop
			var mockedCarStop3 *routemock.MockIMainStop
			var mockedCarStop4 *routemock.MockIMainStop
			var mockedRoute *routemock.MockIMainRoute
			var sut Finder

			BeforeEach(func() {
				mockedCtrl = gomock.NewController(GinkgoT())
				mockedCarStop1 = routemock.NewMockIMainStop(mockedCtrl)
				mockedCarStop2 = routemock.NewMockIMainStop(mockedCtrl)
				mockedCarStop3 = routemock.NewMockIMainStop(mockedCtrl)
				mockedCarStop4 = routemock.NewMockIMainStop(mockedCtrl)
				mockedRoute = routemock.NewMockIMainRoute(mockedCtrl)

				sut = finder{&info{
					itinerary: &itinerary{
						route: mockedRoute,
					},
				}}
			})

			Context("when only one stop is swappable", func() {
				It("should return a list with the swappable car stops cost", func() {
					mockedRoute.EXPECT().Iterator().Return(slc.NewIterator([]route.IMainStop{mockedCarStop1, mockedCarStop2, mockedCarStop3, mockedCarStop4}))
					mockedCarStop1.EXPECT().Point().Return(gps.Point{Latitude: 1})
					mockedCarStop2.EXPECT().Point().Return(gps.Point{Latitude: 2, Longitude: 10})
					mockedCarStop2.EXPECT().IsWarehouse().Return(false)
					mockedCarStop2.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})
					mockedCarStop2.EXPECT().ReturningSubRoutes().Return([]route.ISubRoute{})
					mockedCarStop3.EXPECT().Point().Return(gps.Point{Latitude: 3, Longitude: -10})
					mockedCarStop3.EXPECT().IsWarehouse().Return(false)
					mockedCarStop3.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{nil})
					mockedCarStop4.EXPECT().IsWarehouse().Return(true)

					expectedCarStopCosts := []CarStopCost{
						{
							Stop:  mockedCarStop2,
							Index: 1,
							cost:  20,
						},
					}

					receivedCarStopCosts := sut.FindWorstSwappableCarStopsOrdered()
					Expect(receivedCarStopCosts).To(Equal(expectedCarStopCosts))
				})
			})

			Context("when both stops are swappable", func() {
				It("should return a list with the car stops costs ordered", func() {
					mockedRoute.EXPECT().Iterator().Return(slc.NewIterator([]route.IMainStop{mockedCarStop1, mockedCarStop2, mockedCarStop3, mockedCarStop4}))
					mockedCarStop1.EXPECT().Point().Return(gps.Point{Latitude: 1})
					mockedCarStop2.EXPECT().Point().Return(gps.Point{Latitude: 2, Longitude: 10}).Times(2)
					mockedCarStop2.EXPECT().IsWarehouse().Return(false)
					mockedCarStop2.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})
					mockedCarStop2.EXPECT().ReturningSubRoutes().Return([]route.ISubRoute{})
					mockedCarStop3.EXPECT().Point().Return(gps.Point{Latitude: 3, Longitude: -10}).Times(2)
					mockedCarStop3.EXPECT().IsWarehouse().Return(false)
					mockedCarStop3.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})
					mockedCarStop3.EXPECT().ReturningSubRoutes().Return([]route.ISubRoute{})
					mockedCarStop4.EXPECT().Point().Return(gps.Point{Latitude: 4, Longitude: -5})
					mockedCarStop4.EXPECT().IsWarehouse().Return(true)

					expectedCarStopCosts := []CarStopCost{
						{
							Stop:  mockedCarStop3,
							Index: 2,
							cost:  10,
						},
						{
							Stop:  mockedCarStop2,
							Index: 1,
							cost:  20,
						},
					}

					receivedCarStopCosts := sut.FindWorstSwappableCarStopsOrdered()
					Expect(receivedCarStopCosts).To(Equal(expectedCarStopCosts))
				})
			})
		})
	})
})

var _ = Describe("findWorstDroneStopInFlight", func() {
	Context("when there are drone stops", func() {
		var mockedCtrl *gomock.Controller
		var mockedFlight *routemock.MockISubRoute
		var mockedStartingPoint *routemock.MockIMainStop
		var mockedReturningPoint *routemock.MockIMainStop
		var mockedDroneStop1 *routemock.MockIMainStop
		var mockedDroneStop2 *routemock.MockIMainStop

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedFlight = routemock.NewMockISubRoute(mockedCtrl)
			mockedStartingPoint = routemock.NewMockIMainStop(mockedCtrl)
			mockedReturningPoint = routemock.NewMockIMainStop(mockedCtrl)
			mockedDroneStop1 = routemock.NewMockIMainStop(mockedCtrl)
			mockedDroneStop2 = routemock.NewMockIMainStop(mockedCtrl)
		})

		Context("when there is only one stop", func() {
			It("should return the stop", func() {
				mockedFlight.EXPECT().First().Return(mockedDroneStop1).Times(3)
				mockedFlight.EXPECT().Last().Return(mockedDroneStop1)
				mockedFlight.EXPECT().StartingStop().Return(mockedStartingPoint)
				mockedFlight.EXPECT().ReturningStop().Return(mockedReturningPoint)
				mockedStartingPoint.EXPECT().Point().Return(gps.Point{Latitude: 1})
				mockedDroneStop1.EXPECT().Point().Return(gps.Point{Latitude: 2, Longitude: 10})
				mockedReturningPoint.EXPECT().Point().Return(gps.Point{Latitude: 3})

				expectedDroneStopCost := DroneStopCost{
					Index:  0,
					Stop:   mockedDroneStop1,
					Flight: mockedFlight,
					cost:   20,
				}

				receivedDroneStopCost := findWorstDroneStopInFlight(mockedFlight)
				Expect(receivedDroneStopCost).To(Equal(expectedDroneStopCost))
			})
		})

		Context("when first stop is the worst", func() {
			It("should return first stop cost", func() {
				mockedFlight.EXPECT().First().Return(mockedDroneStop1)
				mockedFlight.EXPECT().Last().Return(mockedDroneStop2)
				mockedFlight.EXPECT().StartingStop().Return(mockedStartingPoint)
				mockedFlight.EXPECT().ReturningStop().Return(mockedReturningPoint)
				mockedFlight.EXPECT().Iterator().Return(slc.NewIterator([]route.ISubStop{mockedDroneStop1, mockedDroneStop2}))
				mockedStartingPoint.EXPECT().Point().Return(gps.Point{Latitude: 1})
				mockedDroneStop1.EXPECT().Point().Return(gps.Point{Latitude: 2, Longitude: 10})
				mockedDroneStop2.EXPECT().Point().Return(gps.Point{Latitude: 3}).Times(2)
				mockedReturningPoint.EXPECT().Point().Return(gps.Point{Latitude: 4})

				expectedDroneStopCost := DroneStopCost{
					Index:  0,
					Stop:   mockedDroneStop1,
					Flight: mockedFlight,
					cost:   20,
				}

				receivedDroneStopCost := findWorstDroneStopInFlight(mockedFlight)
				Expect(receivedDroneStopCost).To(Equal(expectedDroneStopCost))
			})
		})

		Context("when last stop is the worst", func() {
			It("should return last stop cost", func() {
				mockedFlight.EXPECT().First().Return(mockedDroneStop1)
				mockedFlight.EXPECT().Last().Return(mockedDroneStop2)
				mockedFlight.EXPECT().StartingStop().Return(mockedStartingPoint)
				mockedFlight.EXPECT().ReturningStop().Return(mockedReturningPoint)
				mockedFlight.EXPECT().Iterator().Return(slc.NewIterator([]route.ISubStop{mockedDroneStop1, mockedDroneStop2}))
				mockedStartingPoint.EXPECT().Point().Return(gps.Point{Latitude: 1})
				mockedDroneStop1.EXPECT().Point().Return(gps.Point{Latitude: 2})
				mockedDroneStop2.EXPECT().Point().Return(gps.Point{Latitude: 3, Longitude: 10}).Times(2)
				mockedReturningPoint.EXPECT().Point().Return(gps.Point{Latitude: 4})

				expectedDroneStopCost := DroneStopCost{
					Index:  1,
					Stop:   mockedDroneStop2,
					Flight: mockedFlight,
					cost:   20,
				}

				receivedDroneStopCost := findWorstDroneStopInFlight(mockedFlight)
				Expect(receivedDroneStopCost).To(Equal(expectedDroneStopCost))
			})
		})
	})
})

var _ = Describe("insertCarStopCostOrdered", func() {
	Context("when there are no car stops", func() {
		It("should return a list with the received car stop cost", func() {
			carStopCost := CarStopCost{
				Index: 1,
				cost:  1,
			}

			expectedCarStopCosts := []CarStopCost{carStopCost}
			receivedCarStopCosts := insertCarStopCostOrdered([]CarStopCost{}, carStopCost)
			Expect(receivedCarStopCosts).To(Equal(expectedCarStopCosts))
		})
	})

	Context("when there are car stops", func() {
		var carStopCost1 CarStopCost
		var carStopCost2 CarStopCost
		var carStopCost3 CarStopCost

		BeforeEach(func() {
			carStopCost1 = CarStopCost{
				Index: 1,
				cost:  1,
			}

			carStopCost2 = CarStopCost{
				Index: 2,
				cost:  2,
			}

			carStopCost3 = CarStopCost{
				Index: 3,
				cost:  3,
			}
		})

		Context("when the new car stop cost is the cheaper", func() {
			It("should return a list with the received car stop cost at the start", func() {
				actualCarStopCosts := []CarStopCost{carStopCost2, carStopCost3}
				expectedCarStopCosts := []CarStopCost{carStopCost1, carStopCost2, carStopCost3}
				receivedCarStopCosts := insertCarStopCostOrdered(actualCarStopCosts, carStopCost1)
				Expect(receivedCarStopCosts).To(Equal(expectedCarStopCosts))
			})
		})

		Context("when the new car stop cost is the most expensive", func() {
			It("should return a list with the received car stop cost at the end", func() {
				actualCarStopCosts := []CarStopCost{carStopCost1, carStopCost2}
				expectedCarStopCosts := []CarStopCost{carStopCost1, carStopCost2, carStopCost3}
				receivedCarStopCosts := insertCarStopCostOrdered(actualCarStopCosts, carStopCost3)
				Expect(receivedCarStopCosts).To(Equal(expectedCarStopCosts))
			})
		})

		Context("when the new car stop cost is in the middle", func() {
			It("should return a list with the received car stop cost in the middle", func() {
				actualCarStopCosts := []CarStopCost{carStopCost1, carStopCost3}
				expectedCarStopCosts := []CarStopCost{carStopCost1, carStopCost2, carStopCost3}
				receivedCarStopCosts := insertCarStopCostOrdered(actualCarStopCosts, carStopCost2)
				Expect(receivedCarStopCosts).To(Equal(expectedCarStopCosts))
			})
		})
	})
})
