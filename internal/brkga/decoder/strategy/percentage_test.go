package strategy

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
	mockvehicle "github.com/victorguarana/vehicle-routing/internal/vehicle/mock"
	"go.uber.org/mock/gomock"
)

var _ = Describe("vehicleChooserByPercentage", func() {
	var sut vehicleChooserByPercentage
	var mockedCtrl *gomock.Controller
	var mockedCar1 *mockvehicle.MockICar
	var mockedCar2 *mockvehicle.MockICar
	var mockedDrone1 *mockvehicle.MockIDrone
	var mockedDrone2 *mockvehicle.MockIDrone

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedCar1 = mockvehicle.NewMockICar(mockedCtrl)
		mockedCar2 = mockvehicle.NewMockICar(mockedCtrl)
		mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)

		sut = vehicleChooserByPercentage{}
	})

	Describe("NewVehicleChooserByPercentage", func() {
		It("should create vehicleChooserByPercentage with correct params", func() {
			gpsMap := gps.Map{
				Customers: []gps.Point{
					{Name: "customer1"},
					{Name: "customer2"},
					{Name: "customer3"},
				},
			}

			receivedVehicleChooser := NewVehicleChooserByPercentage(gpsMap, 0.6)
			expectedVehicleChooser := &vehicleChooserByPercentage{
				gpsMap:          gpsMap,
				dronePercentage: 0.6,
				carPercentage:   0.4,
			}

			Expect(receivedVehicleChooser).To(Equal(expectedVehicleChooser))
		})
	})

	Describe("DefineVehicle", func() {
		Context("when the gene corresponds to a car", func() {
			It("should return only car", func() {
				sut.gpsMap.Customers = []gps.Point{{}, {}, {}}
				sut.carPercentage = 0.7
				carList := []vehicle.ICar{mockedCar1, mockedCar2}
				c := brkga.Chromossome(0.1)
				chromossome1 := &c

				mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{}).AnyTimes()
				mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{}).AnyTimes()

				receivedCar, receivedDrone := sut.DefineVehicle(carList, chromossome1)
				Expect(receivedCar).To(Equal(mockedCar1))
				Expect(receivedDrone).To(BeNil())
			})
		})

		Context("when the gene corresponds to a drone", func() {
			It("should return drone and car", func() {
				sut.gpsMap.Customers = []gps.Point{{}, {}, {}}
				sut.carPercentage = 0.4
				sut.dronePercentage = 0.6
				carList := []vehicle.ICar{mockedCar1, mockedCar2}
				c := brkga.Chromossome(0.9)
				chromossome1 := &c

				mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1}).AnyTimes()
				mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone2}).AnyTimes()

				receivedCar, receivedDrone := sut.DefineVehicle(carList, chromossome1)
				Expect(receivedCar).To(Equal(mockedCar2))
				Expect(receivedDrone).To(Equal(mockedDrone2))
			})
		})
	})

	Describe("DefineWindowTime", func() {
		Context("when chromossome is min value", func() {
			It("should calc window time", func() {
				sut.gpsMap.Customers = []gps.Point{{}, {}}
				c := brkga.Chromossome(0)
				chromossome1 := &c

				receivedTimeWindowIndex := sut.DefineWindowTime(nil, chromossome1)
				Expect(receivedTimeWindowIndex).To(Equal(0))
			})
		})

		Context("when chromossome is max value", func() {
			It("should calc window time", func() {
				sut.gpsMap.Customers = []gps.Point{{}, {}}
				c := brkga.Chromossome(1)
				chromossome1 := &c

				receivedTimeWindowIndex := sut.DefineWindowTime(nil, chromossome1)
				Expect(receivedTimeWindowIndex).To(Equal(2))
			})
		})

		Context("when chromossome is any between min and max value", func() {
			It("should calc window time", func() {
				sut.gpsMap.Customers = []gps.Point{{}, {}}
				c := brkga.Chromossome(0.3)
				chromossome1 := &c

				receivedTimeWindowIndex := sut.DefineWindowTime(nil, chromossome1)
				Expect(receivedTimeWindowIndex).To(Equal(0))
			})

			It("should calc window time", func() {
				sut.gpsMap.Customers = []gps.Point{{}, {}}
				c := brkga.Chromossome(0.7)
				chromossome1 := &c

				receivedTimeWindowIndex := sut.DefineWindowTime(nil, chromossome1)
				Expect(receivedTimeWindowIndex).To(Equal(1))
			})
		})
	})

	Describe("calcModuledGene", func() {
		It("should calculate the moduled gene", func() {
			sut.gpsMap.Customers = []gps.Point{{}, {}, {}, {}}
			c := brkga.Chromossome(0.8)
			chromossome1 := &c

			receivedGene := sut.calcModuledGene(chromossome1)
			roundedReceivedGene := math.Round(receivedGene*10) / 10
			Expect(roundedReceivedGene).To(Equal(0.2))
		})
	})

	Describe("calcGeneAmplifier", func() {
		It("should calculate the gene amplifier", func() {
			sut.gpsMap.Customers = []gps.Point{{}, {}, {}}
			Expect(sut.calcGeneAmplifier()).To(Equal(3.0))
		})
	})

	Describe("defineDrone", func() {
		It("should return the first drone and car", func() {
			sut.dronePercentage = 0.3
			sut.carPercentage = 0.7
			carList := []vehicle.ICar{mockedCar1, mockedCar2}

			mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1}).AnyTimes()
			mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone2}).AnyTimes()

			receivedCar, receivedDrone := sut.defineDrone(carList, 0.85)
			Expect(receivedCar).To(BeIdenticalTo(mockedCar1))
			Expect(receivedDrone).To(BeIdenticalTo(mockedDrone1))
		})

		It("should return the second drone and car", func() {
			sut.dronePercentage = 0.3
			sut.carPercentage = 0.7
			carList := []vehicle.ICar{mockedCar1, mockedCar2}

			mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1}).AnyTimes()
			mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone2}).AnyTimes()

			receivedCar, receivedDrone := sut.defineDrone(carList, 1)
			Expect(receivedCar).To(BeIdenticalTo(mockedCar2))
			Expect(receivedDrone).To(BeIdenticalTo(mockedDrone2))
		})
	})

	Describe("defineCar", func() {
		It("should return the first car", func() {
			sut.carPercentage = 0.4
			carList := []vehicle.ICar{mockedCar1, mockedCar2}

			receivedCar := sut.defineCar(carList, 0.2)
			Expect(receivedCar).To(BeIdenticalTo(mockedCar1))
		})

		It("should return the second car", func() {
			sut.carPercentage = 0.4
			carList := []vehicle.ICar{mockedCar1, mockedCar2}

			receivedCar := sut.defineCar(carList, 0.4)
			Expect(receivedCar).To(BeIdenticalTo(mockedCar2))
		})
	})
})
