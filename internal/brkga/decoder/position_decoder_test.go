package decoder

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
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
	var mockedDrone2 *mockvehicle.MockIDrone
	var mockedItinerary1 *mockitinerary.MockItinerary
	var mockedConstructor1 *mockitinerary.MockConstructor

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedCar1 = mockvehicle.NewMockICar(mockedCtrl)
		mockedCar2 = mockvehicle.NewMockICar(mockedCtrl)
		mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedItinerary1 = mockitinerary.NewMockItinerary(mockedCtrl)
		mockedConstructor1 = mockitinerary.NewMockConstructor(mockedCtrl)

		sut = positionDecoder{}
	})

	Describe("initializeDecoding", func() {
		var individual *brkga.Individual
		var chromossome1 *brkga.Chromossome
		var chromossome2 *brkga.Chromossome
		var chromossome3 *brkga.Chromossome
		var chromossome4 *brkga.Chromossome
		var warehouse1 gps.Point
		var warehouse2 gps.Point
		var customer1 gps.Point
		var customer2 gps.Point
		var customer3 gps.Point
		var customer4 gps.Point

		var clonedCar1 *mockvehicle.MockICar
		var clonedCar2 *mockvehicle.MockICar

		BeforeEach(func() {
			clonedCar1 = mockvehicle.NewMockICar(mockedCtrl)
			clonedCar2 = mockvehicle.NewMockICar(mockedCtrl)

			c1 := brkga.Chromossome(0.13)
			chromossome1 = &c1
			c2 := brkga.Chromossome(0.14)
			chromossome2 = &c2
			c3 := brkga.Chromossome(0.17)
			chromossome3 = &c3
			c4 := brkga.Chromossome(0.19)
			chromossome4 = &c4

			individual = &brkga.Individual{
				Chromosomes: []*brkga.Chromossome{
					chromossome1, chromossome2, chromossome3, chromossome4},
			}

			warehouse1 = gps.Point{Name: "Warehouse 1"}
			warehouse2 = gps.Point{Name: "Warehouse 2"}
			customer1 = gps.Point{Name: "Customer 1"}
			customer2 = gps.Point{Name: "Customer 2"}
			customer3 = gps.Point{Name: "Customer 3"}
			customer4 = gps.Point{Name: "Customer 4"}

			sut.masterCarList = []vehicle.ICar{mockedCar1, mockedCar2}
			sut.gpsMap.Clients = []gps.Point{customer1, customer2, customer3, customer4}
		})

		It("should initialize instance properties", func() {
			expectedCustomerByChromossomes := map[*brkga.Chromossome]gps.Point{
				chromossome1: customer1,
				chromossome2: customer2,
				chromossome3: customer3,
				chromossome4: customer4,
			}

			expectedOrderedChromossomes := []*brkga.Chromossome{
				chromossome1, chromossome2, chromossome3, chromossome4,
			}

			expectedCarMap := map[*brkga.Chromossome]vehicle.ICar{
				chromossome1: mockedCar1,
				chromossome2: mockedCar1,
				chromossome3: mockedCar2,
				chromossome4: mockedCar2,
			}

			expectedDroneMap := map[*brkga.Chromossome]vehicle.IDrone{
				chromossome3: mockedDrone1,
				chromossome4: mockedDrone2,
			}

			mockedCar1.EXPECT().Clone().Return(clonedCar1)
			mockedCar2.EXPECT().Clone().Return(clonedCar2)
			clonedCar1.EXPECT().ActualPoint().Return(warehouse1)
			clonedCar2.EXPECT().ActualPoint().Return(warehouse2)

			clonedCar1.EXPECT().Storage().Return(2.0).AnyTimes()
			clonedCar2.EXPECT().Storage().Return(2.0).AnyTimes()
			clonedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1}).AnyTimes()
			clonedCar2.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone2}).AnyTimes()
			mockedDrone1.EXPECT().Storage().Return(1.0).AnyTimes()
			mockedDrone2.EXPECT().Storage().Return(1.0).AnyTimes()

			sut.initializeDecoding(individual)

			// Should have set individual
			Expect(sut.individual).To(BeIdenticalTo(individual))

			// Should have cloned cars
			Expect(sut.carList).To(HaveExactElements(BeIdenticalTo(clonedCar1), BeIdenticalTo(clonedCar2)))

			// Should have mapped customers by their chromossomes
			Expect(sut.customerByChromossome).To(Equal(expectedCustomerByChromossomes))

			// Should have ordered chromossomes
			Expect(sut.orderedChromossomes).To(Equal(expectedOrderedChromossomes))

			// Should have created and mapped Itineraries
			Expect(sut.itineraryByCar).To(HaveKey(clonedCar1))
			Expect(sut.itineraryByCar).To(HaveKey(clonedCar2))
			Expect(sut.itineraryByDrone).To(HaveKey(mockedDrone1))
			Expect(sut.itineraryByDrone).To(HaveKey(mockedDrone2))
			Expect(sut.itineraryByCar[clonedCar1]).To(BeIdenticalTo(sut.itineraryByDrone[mockedDrone1]))
			Expect(sut.itineraryByCar[clonedCar2]).To(BeIdenticalTo(sut.itineraryByDrone[mockedDrone2]))

			// Should have mapped vehicles by their chromossomes
			Expect(sut.carByChromossome).To(Equal(expectedCarMap))
			Expect(sut.droneByChromossome).To(Equal(expectedDroneMap))
		})

		BeforeEach(func() {

			individual = &brkga.Individual{
				Chromosomes: []*brkga.Chromossome{
					chromossome1, chromossome2, chromossome3, chromossome4},
			}

			sut.cachedGeneAmplifier = 60
			sut.cachedGeneModule = 6
			sut.carList = []vehicle.ICar{mockedCar1, mockedCar2}
			sut.individual = individual
		})
	})

	Describe("processChromossomes", func() {
		var carChromossome1 *brkga.Chromossome
		var droneChromossome1 *brkga.Chromossome
		var droneChromossome2 *brkga.Chromossome

		var carCustomer1 = gps.Point{Latitude: 1, Name: "Customer 1"}
		var droneCustomer1 = gps.Point{Latitude: 2, Name: "Drone customer 1"}
		var droneCustomer2 = gps.Point{Latitude: 2, Name: "Drone customer 2"}

		var mockedCarStop *mockroute.MockIMainStop

		BeforeEach(func() {
			mockedCarStop = mockroute.NewMockIMainStop(mockedCtrl)

			c1 := brkga.Chromossome(0.1)
			c2 := brkga.Chromossome(0.2)
			c3 := brkga.Chromossome(0.3)
			carChromossome1 = &c2
			droneChromossome1 = &c1
			droneChromossome2 = &c3

			sut.customerByChromossome = map[*brkga.Chromossome]gps.Point{
				carChromossome1: carCustomer1, droneChromossome1: droneCustomer1, droneChromossome2: droneCustomer2,
			}

			sut.droneByChromossome = map[*brkga.Chromossome]vehicle.IDrone{
				droneChromossome1: mockedDrone1, droneChromossome2: mockedDrone1}
			sut.carByChromossome = map[*brkga.Chromossome]vehicle.ICar{carChromossome1: mockedCar1}

			sut.itineraryByDrone = map[vehicle.IDrone]itinerary.Itinerary{mockedDrone1: mockedItinerary1}
			sut.itineraryByCar = map[vehicle.ICar]itinerary.Itinerary{mockedCar1: mockedItinerary1}

			sut.orderedChromossomes = []*brkga.Chromossome{
				droneChromossome1, carChromossome1, droneChromossome2,
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

				sut.processChromossomes()
			})
		})
	})

	Describe("finalizeItineraries", func() {
		var warehouse1 gps.Point
		var warehouse2 gps.Point
		var carPoint1 gps.Point
		var carPoint2 gps.Point

		var mockedItinerary1 *mockitinerary.MockItinerary
		var mockedItinerary2 *mockitinerary.MockItinerary

		var mockedConstructor1 *mockitinerary.MockConstructor
		var mockedConstructor2 *mockitinerary.MockConstructor

		var mockedCarStop1 *mockroute.MockIMainStop
		var mockedCarStop2 *mockroute.MockIMainStop

		BeforeEach(func() {
			warehouse1 = gps.Point{Latitude: 1, Longitude: 1}
			warehouse2 = gps.Point{Latitude: 10, Longitude: 10}
			carPoint1 = gps.Point{Latitude: 0, Longitude: 0}
			carPoint2 = gps.Point{Latitude: 15, Longitude: 15}

			sut.gpsMap.Warehouses = []gps.Point{warehouse1, warehouse2}

			mockedItinerary1 = mockitinerary.NewMockItinerary(mockedCtrl)
			mockedItinerary2 = mockitinerary.NewMockItinerary(mockedCtrl)
			mockedConstructor1 = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedConstructor2 = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedCarStop1 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedCarStop2 = mockroute.NewMockIMainStop(mockedCtrl)

			sut.itineraryByCar = map[vehicle.ICar]itinerary.Itinerary{
				mockedCar1: mockedItinerary1,
				mockedCar2: mockedItinerary2,
			}
		})

		It("should move cars to closest warehouses", func() {
			mockedCar1.EXPECT().ActualPoint().Return(carPoint1)
			mockedCar2.EXPECT().ActualPoint().Return(carPoint2)

			mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1).AnyTimes()
			mockedItinerary2.EXPECT().Constructor().Return(mockedConstructor2).AnyTimes()

			mockedConstructor1.EXPECT().MoveCar(warehouse1)
			mockedConstructor1.EXPECT().ActualCarStop().Return(mockedCarStop1)
			mockedConstructor1.EXPECT().LandAllDrones(mockedCarStop1)

			mockedConstructor2.EXPECT().MoveCar(warehouse2)
			mockedConstructor2.EXPECT().ActualCarStop().Return(mockedCarStop2)
			mockedConstructor2.EXPECT().LandAllDrones(mockedCarStop2)

			sut.finalizeItineraries()
		})
	})

	Describe("cloneCars", func() {
		var clonedCar1 *mockvehicle.MockICar
		var clonedCar2 *mockvehicle.MockICar

		BeforeEach(func() {
			sut.masterCarList = []vehicle.ICar{mockedCar1, mockedCar2}

			clonedCar1 = mockvehicle.NewMockICar(mockedCtrl)
			clonedCar2 = mockvehicle.NewMockICar(mockedCtrl)
		})

		It("should add cloned cars to cars list", func() {
			mockedCar1.EXPECT().Clone().Return(clonedCar1)
			mockedCar2.EXPECT().Clone().Return(clonedCar2)

			sut.cloneCars()

			Expect(sut.carList).To(HaveExactElements(BeIdenticalTo(clonedCar1), BeIdenticalTo(clonedCar2)))
		})
	})

	Describe("mapItineraryByVehicles", func() {
		var point1 gps.Point
		var point2 gps.Point

		BeforeEach(func() {
			point1 = gps.Point{Name: "Point1"}
			point2 = gps.Point{Name: "Point2"}

			sut.carList = []vehicle.ICar{mockedCar1, mockedCar2}
		})

		It("should create itineraries for all cars and use then for theis drones", func() {
			mockedCar1.EXPECT().ActualPoint().Return(point1)
			mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
			mockedCar2.EXPECT().ActualPoint().Return(point2)
			mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{})

			sut.mapItineraryByVehicles()

			Expect(sut.itineraryByCar).To(HaveKey(mockedCar1))
			Expect(sut.itineraryByCar).To(HaveKey(mockedCar2))
			Expect(sut.itineraryByDrone).To(HaveKey(mockedDrone1))
			Expect(sut.itineraryByDrone).To(HaveKey(mockedDrone2))
			Expect(sut.itineraryByCar[mockedCar1]).To(BeIdenticalTo(sut.itineraryByDrone[mockedDrone1]))
			Expect(sut.itineraryByDrone[mockedDrone1]).To(BeIdenticalTo(sut.itineraryByDrone[mockedDrone2]))
		})

	})

	Describe("mapCustomerByChromossome", func() {
		var individual *brkga.Individual
		var chromossome1 *brkga.Chromossome
		var chromossome2 *brkga.Chromossome
		var chromossome3 *brkga.Chromossome
		var chromossome4 *brkga.Chromossome
		var customer1 gps.Point
		var customer2 gps.Point
		var customer3 gps.Point
		var customer4 gps.Point
		var gpsMap gps.Map

		BeforeEach(func() {
			c1 := brkga.Chromossome(0.2)
			chromossome1 = &c1
			c2 := brkga.Chromossome(0.4)
			chromossome2 = &c2
			c3 := brkga.Chromossome(0.1)
			chromossome3 = &c3
			c4 := brkga.Chromossome(0.7)
			chromossome4 = &c4

			individual = &brkga.Individual{
				Chromosomes: []*brkga.Chromossome{
					chromossome1, chromossome2, chromossome3, chromossome4},
			}

			customer1 = gps.Point{Name: "Customer 1"}
			customer2 = gps.Point{Name: "Customer 2"}
			customer3 = gps.Point{Name: "Customer 3"}
			customer4 = gps.Point{Name: "Customer 4"}

			gpsMap.Clients = []gps.Point{customer1, customer2, customer3, customer4}

			sut.individual = individual
			sut.gpsMap = gpsMap
		})

		It("map customers to chromossomes", func() {
			expectedCustomerByChromossomes := map[*brkga.Chromossome]gps.Point{
				chromossome1: customer1,
				chromossome2: customer2,
				chromossome3: customer3,
				chromossome4: customer4,
			}
			sut.mapCustomerByChromossome()

			Expect(sut.customerByChromossome).To(Equal(expectedCustomerByChromossomes))
		})
	})

	Describe("orderChromossomes", func() {
		var individual *brkga.Individual
		var chromossome1 *brkga.Chromossome
		var chromossome2 *brkga.Chromossome
		var chromossome3 *brkga.Chromossome
		var chromossome4 *brkga.Chromossome

		BeforeEach(func() {
			c1 := brkga.Chromossome(0.1)
			chromossome1 = &c1
			c2 := brkga.Chromossome(0.3)
			chromossome2 = &c2
			c3 := brkga.Chromossome(0.5)
			chromossome3 = &c3
			c4 := brkga.Chromossome(0.7)
			chromossome4 = &c4

			individual = &brkga.Individual{
				Chromosomes: []*brkga.Chromossome{
					chromossome4, chromossome2, chromossome1, chromossome3},
			}

			sut.individual = individual
		})

		It("should order chromossomes asc by their genes", func() {
			expectedOrderedChromossomes := []*brkga.Chromossome{
				chromossome1, chromossome2, chromossome3, chromossome4,
			}
			sut.orderChromossomes()

			Expect(sut.orderedChromossomes).To(Equal(expectedOrderedChromossomes))
		})
	})

	Describe("mapChromossomeByVehicle", func() {
		var individual *brkga.Individual
		var chromossome1 *brkga.Chromossome
		var chromossome2 *brkga.Chromossome
		var chromossome3 *brkga.Chromossome
		var chromossome4 *brkga.Chromossome

		BeforeEach(func() {
			c1 := brkga.Chromossome(0.13)
			chromossome1 = &c1
			c2 := brkga.Chromossome(0.14)
			chromossome2 = &c2
			c3 := brkga.Chromossome(0.17)
			chromossome3 = &c3
			c4 := brkga.Chromossome(0.19)
			chromossome4 = &c4

			individual = &brkga.Individual{
				Chromosomes: []*brkga.Chromossome{
					chromossome1, chromossome2, chromossome3, chromossome4},
			}

			sut.cachedGeneAmplifier = 60
			sut.cachedGeneModule = 6
			sut.carList = []vehicle.ICar{mockedCar1, mockedCar2}
			sut.individual = individual
		})

		It("should map chromossomes to vehicles", func() {
			mockedCar1.EXPECT().Storage().Return(2.0).AnyTimes()
			mockedCar2.EXPECT().Storage().Return(2.0).AnyTimes()
			mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1}).AnyTimes()
			mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone2}).AnyTimes()
			mockedDrone1.EXPECT().Storage().Return(1.0).AnyTimes()
			mockedDrone2.EXPECT().Storage().Return(1.0).AnyTimes()

			expectedCarMap := map[*brkga.Chromossome]vehicle.ICar{
				chromossome1: mockedCar1,
				chromossome2: mockedCar1,
				chromossome3: mockedCar2,
				chromossome4: mockedCar2,
			}

			expectedDroneMap := map[*brkga.Chromossome]vehicle.IDrone{
				chromossome3: mockedDrone1,
				chromossome4: mockedDrone2,
			}

			sut.mapChromossomeByVehicle()

			Expect(sut.carByChromossome).To(Equal(expectedCarMap))
			Expect(sut.droneByChromossome).To(Equal(expectedDroneMap))
		})
	})

	Describe("defineVehicle", func() {
		BeforeEach(func() {
			sut.cachedGeneAmplifier = 60
			sut.cachedGeneModule = 12
			sut.carList = []vehicle.ICar{mockedCar1}

		})

		Context("when there is a car chromossome", func() {
			It("should return only car", func() {
				c := brkga.Chromossome(0.1)
				chromossome1 := &c

				mockedCar1.EXPECT().Storage().Return(10.0)

				receivedCar, receivedDrone := sut.defineVehicle(chromossome1)
				Expect(receivedCar).To(Equal(mockedCar1))
				Expect(receivedDrone).To(BeNil())
			})
		})

		Context("when there is a drone chromossome", func() {
			It("should return drone and car", func() {
				c := brkga.Chromossome(0.18)
				chromossome1 := &c

				mockedCar1.EXPECT().Storage().Return(10.0)
				mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1})
				mockedDrone1.EXPECT().Storage().Return(2.0)

				receivedCar, receivedDrone := sut.defineVehicle(chromossome1)
				Expect(receivedCar).To(Equal(mockedCar1))
				Expect(receivedDrone).To(Equal(mockedDrone1))
			})
		})

	})

	Describe("decodeDroneChromossome", func() {
		var chromossome1 *brkga.Chromossome
		var customer1 = gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3}
		var mockedCarStop *mockroute.MockIMainStop

		BeforeEach(func() {
			mockedCarStop = mockroute.NewMockIMainStop(mockedCtrl)

			c := brkga.Chromossome(0.1)
			chromossome1 = &c

			sut.droneByChromossome = map[*brkga.Chromossome]vehicle.IDrone{chromossome1: mockedDrone1}
			sut.itineraryByDrone = map[vehicle.IDrone]itinerary.Itinerary{mockedDrone1: mockedItinerary1}
			sut.customerByChromossome = map[*brkga.Chromossome]gps.Point{chromossome1: customer1}
		})

		Context("when drone is flying", func() {
			It("should move chromossome's drone to chromossome's customer", func() {
				mockedDrone1.EXPECT().IsFlying().Return(true)
				mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
				mockedConstructor1.EXPECT().MoveDrone(mockedDrone1, customer1)

				sut.decodeDroneChromossome(chromossome1)
			})
		})

		Context("when drone is not flying", func() {
			It("should start flight and move chromossome's drone to chromossome's customer", func() {
				mockedDrone1.EXPECT().IsFlying().Return(false)
				mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
				mockedConstructor1.EXPECT().ActualCarStop().Return(mockedCarStop)
				mockedConstructor1.EXPECT().StartDroneFlight(mockedDrone1, mockedCarStop)
				mockedConstructor1.EXPECT().MoveDrone(mockedDrone1, customer1)

				sut.decodeDroneChromossome(chromossome1)
			})
		})

	})

	Describe("decodeCarChromossome", func() {
		var chromossome1 *brkga.Chromossome
		var customer1 = gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3}
		var mockedCarStop *mockroute.MockIMainStop

		BeforeEach(func() {
			mockedCarStop = mockroute.NewMockIMainStop(mockedCtrl)

			c := brkga.Chromossome(0.1)
			chromossome1 = &c

			sut.carByChromossome = map[*brkga.Chromossome]vehicle.ICar{chromossome1: mockedCar1}
			sut.itineraryByCar = map[vehicle.ICar]itinerary.Itinerary{mockedCar1: mockedItinerary1}
			sut.customerByChromossome = map[*brkga.Chromossome]gps.Point{chromossome1: customer1}
		})

		It("should move chromossome's car to chromossome's customer", func() {
			mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1)
			mockedConstructor1.EXPECT().ActualCarStop().Return(mockedCarStop)
			mockedConstructor1.EXPECT().LandAllDrones(mockedCarStop)
			mockedConstructor1.EXPECT().MoveCar(customer1)

			sut.decodeCarChromossome(chromossome1)
		})
	})

	Describe("geneAmplifier", func() {
		BeforeEach(func() {
			sut.carList = []vehicle.ICar{mockedCar1, mockedCar2}
			sut.gpsMap.Clients = []gps.Point{{}, {}, {}}
		})

		Context("when cache is empty", func() {
			It("should calc gene amplifier", func() {
				mockedCar1.EXPECT().Storage().Return(2.0)
				mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1})
				mockedDrone1.EXPECT().Storage().Return(2.0)
				mockedCar2.EXPECT().Storage().Return(1.0)
				mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone2})
				mockedDrone2.EXPECT().Storage().Return(1.0)

				Expect(sut.geneAmplifier()).To(Equal(18.0))
				Expect(sut.cachedGeneAmplifier).To(Equal(18.0))
			})
		})

		Context("when cache is not empty", func() {
			It("should not calc gene amplifier", func() {
				sut.cachedGeneAmplifier = 1.0

				mockedCar1.EXPECT().Storage().Return(2.0).Times(0)
				mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1}).Times(0)
				mockedDrone1.EXPECT().Storage().Return(2.0).Times(0)
				mockedCar2.EXPECT().Storage().Return(1.0).Times(0)
				mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone2}).Times(0)
				mockedDrone2.EXPECT().Storage().Return(1.0).Times(0)

				Expect(sut.geneAmplifier()).To(Equal(1.0))
				Expect(sut.cachedGeneAmplifier).To(Equal(1.0))
			})
		})
	})

	Describe("geneModule", func() {
		BeforeEach(func() {
			sut.carList = []vehicle.ICar{mockedCar1, mockedCar2}
			sut.gpsMap.Clients = []gps.Point{{}, {}, {}}
		})

		Context("when cache is empty", func() {
			It("should calc gene amplifier", func() {
				mockedCar1.EXPECT().Storage().Return(2.0)
				mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1})
				mockedDrone1.EXPECT().Storage().Return(2.0)
				mockedCar2.EXPECT().Storage().Return(1.0)
				mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone2})
				mockedDrone2.EXPECT().Storage().Return(1.0)

				Expect(sut.geneModule()).To(Equal(6.0))
				Expect(sut.cachedGeneModule).To(Equal(6.0))
			})
		})

		Context("when cache is not empty", func() {
			It("should not calc gene amplifier", func() {
				sut.cachedGeneModule = 1.0

				mockedCar1.EXPECT().Storage().Return(2.0).Times(0)
				mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1}).Times(0)
				mockedDrone1.EXPECT().Storage().Return(2.0).Times(0)
				mockedCar2.EXPECT().Storage().Return(1.0).Times(0)
				mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone2}).Times(0)
				mockedDrone2.EXPECT().Storage().Return(1.0).Times(0)

				Expect(sut.geneModule()).To(Equal(1.0))
				Expect(sut.cachedGeneModule).To(Equal(1.0))
			})
		})
	})

	Describe("calcTotalStorage", func() {

		BeforeEach(func() {
			sut.carList = []vehicle.ICar{mockedCar1, mockedCar2}
		})

		It("should the sum of storage of all vehicles", func() {
			mockedCar1.EXPECT().Storage().Return(10.1)
			mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
			mockedDrone1.EXPECT().Storage().Return(1.1)
			mockedDrone2.EXPECT().Storage().Return(2.1)
			mockedCar2.EXPECT().Storage().Return(20.1)
			mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{})

			Expect(sut.calcTotalStorage()).To(Equal(33.4))
		})

	})

})
