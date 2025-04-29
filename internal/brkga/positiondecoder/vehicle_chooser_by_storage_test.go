package positiondecoder

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
	mockvehicle "github.com/victorguarana/vehicle-routing/internal/vehicle/mock"
	"go.uber.org/mock/gomock"
)

var _ = Describe("vehicleChooserByStorage", func() {
	var sut vehicleChooserByStorage
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

		sut = vehicleChooserByStorage{}
	})

	Describe("NewVehicleChooserByStorage", func() {
		It("should create vehicleChooserByStorage with correct params", func() {
			gpsMap := gps.Map{
				Clients: []gps.Point{
					{Name: "client1"},
					{Name: "client2"},
					{Name: "client3"},
				},
			}

			receivedVehicleChooser := NewVehicleChooserByStorage(gpsMap)
			expectedVehicleChooser := &vehicleChooserByStorage{
				gpsMap: gpsMap,
			}

			Expect(receivedVehicleChooser).To(Equal(expectedVehicleChooser))
		})
	})

	Describe("defineVehicle", func() {
		Context("when there is a car chromossome", func() {
			It("should return only car", func() {
				sut.gpsMap.Clients = []gps.Point{{}, {}, {}}
				carList := []vehicle.ICar{mockedCar1, mockedCar2}
				c := brkga.Chromossome(0.1)
				chromossome1 := &c

				mockedCar1.EXPECT().Storage().Return(2.0).AnyTimes()
				mockedCar2.EXPECT().Storage().Return(2.0).AnyTimes()
				mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1}).AnyTimes()
				mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone2}).AnyTimes()
				mockedDrone1.EXPECT().Storage().Return(1.0).AnyTimes()
				mockedDrone2.EXPECT().Storage().Return(1.0).AnyTimes()

				receivedCar, receivedDrone := sut.defineVehicle(carList, chromossome1)
				Expect(receivedCar).To(Equal(mockedCar1))
				Expect(receivedDrone).To(BeNil())
			})
		})

		Context("when there is a drone chromossome", func() {
			It("should return drone and car", func() {
				sut.gpsMap.Clients = []gps.Point{{}, {}, {}}
				carList := []vehicle.ICar{mockedCar1, mockedCar2}
				c := brkga.Chromossome(0.25)
				chromossome1 := &c

				mockedCar1.EXPECT().Storage().Return(2.0).AnyTimes()
				mockedCar2.EXPECT().Storage().Return(2.0).AnyTimes()
				mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1}).AnyTimes()
				mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone2}).AnyTimes()
				mockedDrone1.EXPECT().Storage().Return(1.0).AnyTimes()
				mockedDrone2.EXPECT().Storage().Return(1.0).AnyTimes()

				receivedCar, receivedDrone := sut.defineVehicle(carList, chromossome1)
				Expect(receivedCar).To(Equal(mockedCar1))
				Expect(receivedDrone).To(Equal(mockedDrone1))
			})
		})
	})

	Describe("calcModuledGene", func() {
		Context("when there is only cars and customers", func() {
			It("should calc car gene", func() {
				sut.gpsMap.Clients = []gps.Point{{}, {}}
				carList := []vehicle.ICar{mockedCar1, mockedCar2}
				c := brkga.Chromossome(0.75)
				chromossome1 := &c

				mockedCar1.EXPECT().Storage().Return(2.0).AnyTimes()
				mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{}).AnyTimes()
				mockedCar2.EXPECT().Storage().Return(1.0).AnyTimes()
				mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{}).AnyTimes()

				Expect(sut.calcModuledGene(carList, chromossome1)).To(Equal(1.5))
			})
		})

		Context("when there is cars, drones and customers", func() {
			It("should calc car gene", func() {
				sut.gpsMap.Clients = []gps.Point{{}}
				carList := []vehicle.ICar{mockedCar1, mockedCar2}
				c := brkga.Chromossome(0.5)
				chromossome1 := &c

				mockedCar1.EXPECT().Storage().Return(2.0).AnyTimes()
				mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1}).AnyTimes()
				mockedCar2.EXPECT().Storage().Return(2.0).AnyTimes()
				mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone2}).AnyTimes()
				mockedDrone1.EXPECT().Storage().Return(1.0).AnyTimes()
				mockedDrone2.EXPECT().Storage().Return(1.0).AnyTimes()

				Expect(sut.calcModuledGene(carList, chromossome1)).To(Equal(3.0))
			})
		})
	})

	Describe("calcGeneAmplifier", func() {
		BeforeEach(func() {
			sut.gpsMap.Clients = []gps.Point{{}, {}, {}}
		})

		It("should calc gene amplifier", func() {
			carList := []vehicle.ICar{mockedCar1, mockedCar2}

			mockedCar1.EXPECT().Storage().Return(2.0)
			mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1})
			mockedDrone1.EXPECT().Storage().Return(2.0)
			mockedCar2.EXPECT().Storage().Return(1.0)
			mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone2})
			mockedDrone2.EXPECT().Storage().Return(1.0)

			Expect(sut.calcGeneAmplifier(carList)).To(Equal(18.0))
		})
	})

	Describe("calcGeneModule", func() {
		It("should calc gene module", func() {
			carList := []vehicle.ICar{mockedCar1, mockedCar2}

			mockedCar1.EXPECT().Storage().Return(2.0)
			mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1})
			mockedDrone1.EXPECT().Storage().Return(2.0)
			mockedCar2.EXPECT().Storage().Return(1.0)
			mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone2})
			mockedDrone2.EXPECT().Storage().Return(1.0)

			Expect(sut.calcGeneModule(carList)).To(Equal(6.0))
		})
	})

	Describe("calcTotalStorage", func() {
		It("should the sum of storage of all vehicles", func() {
			carList := []vehicle.ICar{mockedCar1, mockedCar2}
			mockedCar1.EXPECT().Storage().Return(10.1)
			mockedCar1.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
			mockedDrone1.EXPECT().Storage().Return(1.1)
			mockedDrone2.EXPECT().Storage().Return(2.1)
			mockedCar2.EXPECT().Storage().Return(20.1)
			mockedCar2.EXPECT().Drones().Return([]vehicle.IDrone{})

			Expect(sut.calcTotalStorage(carList)).To(Equal(33.4))
		})
	})
})
