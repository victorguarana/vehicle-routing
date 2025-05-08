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

var _ = Describe("positionDecoder", func() {
	var sut positionDecoder
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

		sut = positionDecoder{}
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

			receivedDecodedChromossomeList := sut.decodeChromossomeList(chromossomeList)

			Expect(receivedDecodedChromossomeList).To(HaveLen(4))
			Expect(receivedDecodedChromossomeList[0].chromossome).To(BeIdenticalTo(car1Chromossome))
			Expect(receivedDecodedChromossomeList[0].car).To(BeIdenticalTo(clonedCar1))
			Expect(receivedDecodedChromossomeList[0].customer).To(Equal(carCustomer1))
			Expect(receivedDecodedChromossomeList[0].drone).To(BeNil())
			Expect(receivedDecodedChromossomeList[1].chromossome).To(BeIdenticalTo(car2Chromossome))
			Expect(receivedDecodedChromossomeList[1].car).To(BeIdenticalTo(clonedCar2))
			Expect(receivedDecodedChromossomeList[1].customer).To(Equal(carCustomer2))
			Expect(receivedDecodedChromossomeList[1].drone).To(BeNil())
			Expect(receivedDecodedChromossomeList[2].chromossome).To(BeIdenticalTo(drone1Chromossome))
			Expect(receivedDecodedChromossomeList[2].car).To(BeIdenticalTo(clonedCar1))
			Expect(receivedDecodedChromossomeList[2].customer).To(Equal(droneCustomer1))
			Expect(receivedDecodedChromossomeList[2].drone).To(BeIdenticalTo(clonedDrone1))
			Expect(receivedDecodedChromossomeList[3].chromossome).To(BeIdenticalTo(drone2Chromossome))
			Expect(receivedDecodedChromossomeList[3].car).To(BeIdenticalTo(clonedCar2))
			Expect(receivedDecodedChromossomeList[3].customer).To(Equal(droneCustomer2))
			Expect(receivedDecodedChromossomeList[3].drone).To(BeIdenticalTo(clonedDrone2))

		})
	})

	Describe("parseChromossomes", func() {
		var decodedChromossomeList []*decodedChromossome

		var carCustomer1 = gps.Point{Latitude: 1, Name: "Customer 1"}
		var droneCustomer1 = gps.Point{Latitude: 2, Name: "Drone customer 1"}
		var droneCustomer2 = gps.Point{Latitude: 2, Name: "Drone customer 2"}

		var mockedCarStop *mockroute.MockIMainStop

		BeforeEach(func() {
			mockedCarStop = mockroute.NewMockIMainStop(mockedCtrl)

			decodedChromossomeList = []*decodedChromossome{
				{
					car:      mockedCar1,
					customer: carCustomer1,
					itn:      mockedItinerary1,
				},
				{
					car:      mockedCar1,
					customer: droneCustomer1,
					itn:      mockedItinerary1,
					drone:    mockedDrone1,
				},
				{
					car:      mockedCar1,
					customer: droneCustomer2,
					itn:      mockedItinerary1,
					drone:    mockedDrone1,
				},
			}
		})

		Context("when drone is flying", func() {
			It("should move chromossome's drone to chromossome's customer", func() {
				// Moving drone when its not flying
				mockedDrone1.EXPECT().IsFlying().Return(false)
				mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
				mockedConstructor1.EXPECT().ActualCarStop().Return(mockedCarStop)
				mockedConstructor1.EXPECT().StartDroneFlight(mockedDrone1, mockedCarStop)
				mockedConstructor1.EXPECT().MoveDrone(mockedDrone1, droneCustomer1)

				// Moving car
				mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
				mockedConstructor1.EXPECT().ActualCarStop().Return(mockedCarStop)
				mockedConstructor1.EXPECT().LandAllDrones(mockedCarStop)
				mockedConstructor1.EXPECT().MoveCar(carCustomer1)

				// Moving drone when its flying
				mockedDrone1.EXPECT().IsFlying().Return(true)
				mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
				mockedConstructor1.EXPECT().MoveDrone(mockedDrone1, droneCustomer2)

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
		var mockedCarStop *mockroute.MockIMainStop
		var dc *decodedChromossome

		BeforeEach(func() {
			mockedCarStop = mockroute.NewMockIMainStop(mockedCtrl)

			dc = &decodedChromossome{
				car:      mockedCar1,
				itn:      mockedItinerary1,
				customer: customer1,
			}
		})

		It("should move chromossome's car to chromossome's customer", func() {
			mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
			mockedConstructor1.EXPECT().ActualCarStop().Return(mockedCarStop)
			mockedConstructor1.EXPECT().LandAllDrones(mockedCarStop)
			mockedConstructor1.EXPECT().MoveCar(customer1)

			sut.parseDecodedCarChromossome(dc)
		})
	})

	Describe("orderDecodedChromossomes", func() {
		var decodedChromossome1 *decodedChromossome
		var decodedChromossome2 *decodedChromossome
		var decodedChromossome3 *decodedChromossome
		var decodedChromossome4 *decodedChromossome

		BeforeEach(func() {
			c1 := brkga.Chromossome(0.1)
			c2 := brkga.Chromossome(0.2)
			c3 := brkga.Chromossome(0.3)
			c4 := brkga.Chromossome(0.4)

			decodedChromossome1 = &decodedChromossome{
				chromossome: &c1,
			}
			decodedChromossome2 = &decodedChromossome{
				chromossome: &c2,
			}
			decodedChromossome3 = &decodedChromossome{
				chromossome: &c3,
			}
			decodedChromossome4 = &decodedChromossome{
				chromossome: &c4,
			}
		})

		It("should return ordered decoded chromossome list", func() {
			decodedChromossomeList := []*decodedChromossome{
				decodedChromossome2, decodedChromossome4, decodedChromossome3, decodedChromossome1,
			}

			expectedDecodedChromossomeList := []*decodedChromossome{
				decodedChromossome1, decodedChromossome2, decodedChromossome3, decodedChromossome4,
			}

			receivedDecodedChromossomeList := sut.orderDecodedChromossomes(decodedChromossomeList)

			Expect(receivedDecodedChromossomeList).To(HaveExactElements(expectedDecodedChromossomeList))
		})
	})
})
