package decoder

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	mockitinerary "github.com/victorguarana/vehicle-routing/internal/itinerary/mock"
	mockroute "github.com/victorguarana/vehicle-routing/internal/route/mock"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
	mockvehicle "github.com/victorguarana/vehicle-routing/internal/vehicle/mock"
	"go.uber.org/mock/gomock"
)

var _ = Describe("timeWindowDecoder", func() {
	var sut timeWindowDecoder
	var mockedCtrl *gomock.Controller
	var mockedCar1 *mockvehicle.MockICar
	var mockedCar2 *mockvehicle.MockICar
	var mockedDrone1 *mockvehicle.MockIDrone
	var mockedItinerary1 *mockitinerary.MockItinerary
	var mockedConstructor1 *mockitinerary.MockConstructor
	var mockedStrategy *Mockstrategy

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedCar1 = mockvehicle.NewMockICar(mockedCtrl)
		mockedCar2 = mockvehicle.NewMockICar(mockedCtrl)
		mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedItinerary1 = mockitinerary.NewMockItinerary(mockedCtrl)
		mockedConstructor1 = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedStrategy = NewMockstrategy(mockedCtrl)

		sut = timeWindowDecoder{}
	})

	Describe("decodeChromossomeList", func() {
		var initialPoint = gps.Point{Name: "Initial Point"}
		var carCustomer1 = gps.Point{Name: "Customer 1"}
		var carCustomer2 = gps.Point{Name: "Customer 2"}
		var droneCustomer1 = gps.Point{Name: "Drone customer 1"}
		var droneCustomer2 = gps.Point{Name: "Drone customer 2"}

		var clonedCar1 *mockvehicle.MockICar
		var clonedCar2 *mockvehicle.MockICar
		var clonedDrone1 *mockvehicle.MockIDrone
		var clonedDrone2 *mockvehicle.MockIDrone

		var chromossomeList []*brkga.Chromossome
		var carList []vehicle.ICar
		var car1Chromossome *brkga.Chromossome
		var car2Chromossome *brkga.Chromossome
		var drone1Chromossome *brkga.Chromossome
		var drone2Chromossome *brkga.Chromossome

		BeforeEach(func() {
			c1 := brkga.Chromossome(0.1)
			c2 := brkga.Chromossome(0.2)
			d3 := brkga.Chromossome(0.4)
			d4 := brkga.Chromossome(0.4)
			car1Chromossome = &c1
			car2Chromossome = &c2
			drone1Chromossome = &d3
			drone2Chromossome = &d4

			clonedCar1 = mockvehicle.NewMockICar(mockedCtrl)
			clonedCar2 = mockvehicle.NewMockICar(mockedCtrl)
			clonedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			clonedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)

			carList = []vehicle.ICar{clonedCar1, clonedCar2}

			chromossomeList = []*brkga.Chromossome{
				car1Chromossome,
				car2Chromossome,
				drone1Chromossome,
				drone2Chromossome,
			}

			sut.strategy = mockedStrategy
			sut.masterCarList = []vehicle.ICar{mockedCar1, mockedCar2}
			sut.gpsMap = gps.Map{
				Clients:    []gps.Point{carCustomer1, carCustomer2, droneCustomer1, droneCustomer2},
				Warehouses: []gps.Point{initialPoint},
			}
		})

		It("should return decoded chromossome list", func() {
			mockedCar1.EXPECT().Clone().Return(clonedCar1)
			mockedCar2.EXPECT().Clone().Return(clonedCar2)
			clonedCar1.EXPECT().ActualPoint().Return(initialPoint)
			clonedCar2.EXPECT().ActualPoint().Return(initialPoint)

			mockedStrategy.EXPECT().DefineVehicle(carList, car1Chromossome).Return(clonedCar1, nil)
			mockedStrategy.EXPECT().DefineVehicle(carList, car2Chromossome).Return(clonedCar2, nil)
			mockedStrategy.EXPECT().DefineVehicle(carList, drone1Chromossome).Return(clonedCar1, clonedDrone1)
			mockedStrategy.EXPECT().DefineVehicle(carList, drone2Chromossome).Return(clonedCar2, clonedDrone2)
			mockedStrategy.EXPECT().DefineWindowTime(carList, car1Chromossome).Return(1)
			mockedStrategy.EXPECT().DefineWindowTime(carList, car2Chromossome).Return(2)
			mockedStrategy.EXPECT().DefineWindowTime(carList, drone1Chromossome).Return(3)
			mockedStrategy.EXPECT().DefineWindowTime(carList, drone2Chromossome).Return(4)

			receivedDecodedChromossomeList := sut.decodeChromossomeList(chromossomeList)

			Expect(receivedDecodedChromossomeList).To(HaveLen(4))
			Expect(receivedDecodedChromossomeList[0].chromossome).To(BeIdenticalTo(car1Chromossome))
			Expect(receivedDecodedChromossomeList[0].car).To(BeIdenticalTo(clonedCar1))
			Expect(receivedDecodedChromossomeList[0].customer).To(Equal(carCustomer1))
			Expect(receivedDecodedChromossomeList[0].drone).To(BeNil())
			Expect(receivedDecodedChromossomeList[0].timeWindowIndex).To(Equal(1))
			Expect(receivedDecodedChromossomeList[1].chromossome).To(BeIdenticalTo(car2Chromossome))
			Expect(receivedDecodedChromossomeList[1].car).To(BeIdenticalTo(clonedCar2))
			Expect(receivedDecodedChromossomeList[1].customer).To(Equal(carCustomer2))
			Expect(receivedDecodedChromossomeList[1].drone).To(BeNil())
			Expect(receivedDecodedChromossomeList[1].timeWindowIndex).To(Equal(2))
			Expect(receivedDecodedChromossomeList[2].chromossome).To(BeIdenticalTo(drone1Chromossome))
			Expect(receivedDecodedChromossomeList[2].car).To(BeIdenticalTo(clonedCar1))
			Expect(receivedDecodedChromossomeList[2].customer).To(Equal(droneCustomer1))
			Expect(receivedDecodedChromossomeList[2].drone).To(BeIdenticalTo(clonedDrone1))
			Expect(receivedDecodedChromossomeList[2].timeWindowIndex).To(Equal(3))
			Expect(receivedDecodedChromossomeList[3].chromossome).To(BeIdenticalTo(drone2Chromossome))
			Expect(receivedDecodedChromossomeList[3].car).To(BeIdenticalTo(clonedCar2))
			Expect(receivedDecodedChromossomeList[3].customer).To(Equal(droneCustomer2))
			Expect(receivedDecodedChromossomeList[3].drone).To(BeIdenticalTo(clonedDrone2))
			Expect(receivedDecodedChromossomeList[3].timeWindowIndex).To(Equal(4))

		})
	})

	Describe("parseChromossomes", func() {
		var carCustomer1 = gps.Point{Name: "Customer 1"}
		var carCustomer2 = gps.Point{Name: "Customer 2"}
		var droneCustomer1 = gps.Point{Name: "Drone customer 1"}
		var droneCustomer2 = gps.Point{Name: "Drone customer 2"}
		var chromossome1 *brkga.Chromossome
		var chromossome2 *brkga.Chromossome
		var chromossome3 *brkga.Chromossome
		var chromossome4 *brkga.Chromossome

		var mockedCarStop *mockroute.MockIMainStop

		BeforeEach(func() {
			mockedCarStop = mockroute.NewMockIMainStop(mockedCtrl)
			c1 := brkga.Chromossome(0.1)
			c2 := brkga.Chromossome(0.2)
			c3 := brkga.Chromossome(0.4)
			c4 := brkga.Chromossome(0.4)
			chromossome1 = &c1
			chromossome2 = &c2
			chromossome3 = &c3
			chromossome4 = &c4

		})

		Context("when there are simultaneous chromossomes", func() {
			It("should move drone before car", func() {
				decodedChromossomeList := []*decodedChromossome{
					{
						car:             mockedCar1,
						customer:        droneCustomer2,
						itn:             mockedItinerary1,
						drone:           mockedDrone1,
						chromossome:     chromossome3,
						timeWindowIndex: 1,
					},
					{
						car:             mockedCar1,
						customer:        carCustomer1,
						itn:             mockedItinerary1,
						chromossome:     chromossome1,
						timeWindowIndex: 1,
					},
					{
						car:             mockedCar1,
						customer:        droneCustomer1,
						itn:             mockedItinerary1,
						drone:           mockedDrone1,
						chromossome:     chromossome2,
						timeWindowIndex: 1,
					},
					{
						car:             mockedCar1,
						customer:        carCustomer2,
						itn:             mockedItinerary1,
						chromossome:     chromossome4,
						timeWindowIndex: 1,
					},
				}

				// Moving drone when its not flying
				mockedDrone1.EXPECT().IsFlying().Return(false)
				mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
				mockedConstructor1.EXPECT().ActualCarStop().Return(mockedCarStop)
				mockedConstructor1.EXPECT().StartDroneFlight(mockedDrone1, mockedCarStop)
				mockedConstructor1.EXPECT().MoveDrone(mockedDrone1, droneCustomer1)

				// Moving drone when its flying
				mockedDrone1.EXPECT().IsFlying().Return(true)
				mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
				mockedConstructor1.EXPECT().MoveDrone(mockedDrone1, droneCustomer2)

				// Moving car
				mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
				mockedConstructor1.EXPECT().MoveCar(carCustomer1)

				// Moving car
				mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
				mockedConstructor1.EXPECT().MoveCar(carCustomer2)

				// Land drones at the end of the time window
				mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
				mockedConstructor1.EXPECT().ActualCarStop().Return(mockedCarStop)
				mockedConstructor1.EXPECT().LandAllDrones(mockedCarStop)

				sut.parseChromossomes(decodedChromossomeList)
			})
		})
	})

	Describe("parseDecodedDroneChromossome", func() {
		var customer1 = gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3}
		var mockedCarStop *mockroute.MockIMainStop
		var dc *decodedChromossome

		BeforeEach(func() {
			mockedCarStop = mockroute.NewMockIMainStop(mockedCtrl)

			dc = &decodedChromossome{
				car:      mockedCar1,
				drone:    mockedDrone1,
				itn:      mockedItinerary1,
				customer: customer1,
			}
		})

		Context("when drone is flying", func() {
			It("should move chromossome's drone to chromossome's customer", func() {
				mockedDrone1.EXPECT().IsFlying().Return(true)
				mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
				mockedConstructor1.EXPECT().MoveDrone(mockedDrone1, customer1)

				sut.parseDecodedDroneChromossome(dc)
			})
		})

		Context("when drone is not flying", func() {
			It("should start flight and move chromossome's drone to chromossome's customer", func() {
				mockedDrone1.EXPECT().IsFlying().Return(false)
				mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
				mockedConstructor1.EXPECT().ActualCarStop().Return(mockedCarStop)
				mockedConstructor1.EXPECT().StartDroneFlight(mockedDrone1, mockedCarStop)
				mockedConstructor1.EXPECT().MoveDrone(mockedDrone1, customer1)

				sut.parseDecodedDroneChromossome(dc)
			})
		})

	})

	Describe("parseDecodedCarChromossome", func() {
		var customer1 = gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3}
		var dc *decodedChromossome

		BeforeEach(func() {
			dc = &decodedChromossome{
				car:      mockedCar1,
				itn:      mockedItinerary1,
				customer: customer1,
			}
		})

		It("should move chromossome's car to chromossome's customer", func() {
			mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
			mockedConstructor1.EXPECT().MoveCar(customer1)

			sut.parseDecodedCarChromossome(dc)
		})
	})

	Describe("orderTimeWindow", func() {
		It("should return the ordered time windows without duplicates", func() {
			timeWindowList := []int{3, 1, 2}
			orderedTimeWindow := sut.orderTimeWindow(timeWindowList)
			Expect(orderedTimeWindow).To(HaveLen(3))
			Expect(orderedTimeWindow[0]).To(Equal(1))
			Expect(orderedTimeWindow[1]).To(Equal(2))
			Expect(orderedTimeWindow[2]).To(Equal(3))
		})

		It("should return the ordered time windows with duplicates", func() {
			timeWindowList := []int{2, 3, 1, 2, 1}
			orderedTimeWindow := sut.orderTimeWindow(timeWindowList)
			Expect(orderedTimeWindow).To(HaveLen(5))
			Expect(orderedTimeWindow[0]).To(Equal(1))
			Expect(orderedTimeWindow[1]).To(Equal(1))
			Expect(orderedTimeWindow[2]).To(Equal(2))
			Expect(orderedTimeWindow[3]).To(Equal(2))
			Expect(orderedTimeWindow[4]).To(Equal(3))
		})
	})

	Describe("mapDecodedChromossomeByTimeWindow", func() {
		It("should return a map of decoded chromossomes by time window", func() {
			decodedChromossomome1_0 := &decodedChromossome{timeWindowIndex: 1}
			decodedChromossomome1_1 := &decodedChromossome{timeWindowIndex: 1}
			decodedChromossomome2_0 := &decodedChromossome{timeWindowIndex: 2}
			decodedChromossomome3_0 := &decodedChromossome{timeWindowIndex: 3}
			decodedChromossomome3_1 := &decodedChromossome{timeWindowIndex: 3}
			decodedChromossomeList := []*decodedChromossome{
				decodedChromossomome1_0,
				decodedChromossomome3_0,
				decodedChromossomome2_0,
				decodedChromossomome1_1,
				decodedChromossomome3_1,
			}

			decodedChromossomesByTimeWindow := sut.mapDecodedChromossomeByTimeWindow(decodedChromossomeList)

			Expect(decodedChromossomesByTimeWindow).To(HaveLen(3))
			Expect(decodedChromossomesByTimeWindow[1]).To(HaveLen(2))
			Expect(decodedChromossomesByTimeWindow[2]).To(HaveLen(1))
			Expect(decodedChromossomesByTimeWindow[3]).To(HaveLen(2))
			Expect(decodedChromossomesByTimeWindow[1][0]).To(BeIdenticalTo(decodedChromossomome1_0))
			Expect(decodedChromossomesByTimeWindow[1][1]).To(BeIdenticalTo(decodedChromossomome1_1))
			Expect(decodedChromossomesByTimeWindow[2][0]).To(BeIdenticalTo(decodedChromossomome2_0))
			Expect(decodedChromossomesByTimeWindow[3][0]).To(BeIdenticalTo(decodedChromossomome3_0))
			Expect(decodedChromossomesByTimeWindow[3][1]).To(BeIdenticalTo(decodedChromossomome3_1))
		})
	})
})
