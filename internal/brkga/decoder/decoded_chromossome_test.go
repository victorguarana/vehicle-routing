package decoder

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/victorguarana/vehicle-routing/internal/brkga"
	mockvehicle "github.com/victorguarana/vehicle-routing/internal/vehicle/mock"
	"go.uber.org/mock/gomock"
)

var _ = Describe("DecodedChromossome", func() {

	Describe("orderDecodedChromossomesByChromossome", func() {
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

			receivedDecodedChromossomeList := orderDecodedChromossomesByChromossome(decodedChromossomeList)

			Expect(receivedDecodedChromossomeList).To(HaveExactElements(expectedDecodedChromossomeList))
		})
	})

	Describe("collectDroneChromossomes", func() {
		var mockedCtrl *gomock.Controller
		var mockedCar *mockvehicle.MockICar
		var mockedDrone *mockvehicle.MockIDrone

		var carDecodedChromossome1 *decodedChromossome
		var carDecodedChromossome2 *decodedChromossome
		var droneDecodedChromossome1 *decodedChromossome
		var droneDecodedChromossome2 *decodedChromossome

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			mockedDrone = mockvehicle.NewMockIDrone(mockedCtrl)

			carDecodedChromossome1 = &decodedChromossome{
				car: mockedCar,
			}
			carDecodedChromossome2 = &decodedChromossome{
				car: mockedCar,
			}
			droneDecodedChromossome1 = &decodedChromossome{
				car:   mockedCar,
				drone: mockedDrone,
			}
			droneDecodedChromossome2 = &decodedChromossome{
				car:   mockedCar,
				drone: mockedDrone,
			}

		})

		It("should return drone chromossome list", func() {
			decodedChromossomeList := []*decodedChromossome{
				carDecodedChromossome1, carDecodedChromossome2, droneDecodedChromossome1, droneDecodedChromossome2,
			}

			reveicedDroneChromossomeList := collectDroneChromossomes(decodedChromossomeList)

			Expect(reveicedDroneChromossomeList).To(HaveLen(2))
			Expect(reveicedDroneChromossomeList[0]).To(BeIdenticalTo(droneDecodedChromossome1))
			Expect(reveicedDroneChromossomeList[1]).To(BeIdenticalTo(droneDecodedChromossome2))
		})

		It("should return empty drone chromossome list", func() {
			reveicedDroneChromossomeList := collectDroneChromossomes([]*decodedChromossome{})

			Expect(reveicedDroneChromossomeList).To(BeEmpty())
		})

		It("should return empty drone chromossome list when no drone chromossome is present", func() {
			decodedChromossomeList := []*decodedChromossome{
				carDecodedChromossome1, carDecodedChromossome2,
			}

			reveicedDroneChromossomeList := collectDroneChromossomes(decodedChromossomeList)

			Expect(reveicedDroneChromossomeList).To(BeEmpty())
		})
	})

	Describe("collectCarChromossomes", func() {
		var mockedCtrl *gomock.Controller
		var mockedCar *mockvehicle.MockICar
		var mockedDrone *mockvehicle.MockIDrone

		var carDecodedChromossome1 *decodedChromossome
		var carDecodedChromossome2 *decodedChromossome
		var droneDecodedChromossome1 *decodedChromossome
		var droneDecodedChromossome2 *decodedChromossome

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			mockedDrone = mockvehicle.NewMockIDrone(mockedCtrl)

			carDecodedChromossome1 = &decodedChromossome{
				car: mockedCar,
			}
			carDecodedChromossome2 = &decodedChromossome{
				car: mockedCar,
			}
			droneDecodedChromossome1 = &decodedChromossome{
				car:   mockedCar,
				drone: mockedDrone,
			}
			droneDecodedChromossome2 = &decodedChromossome{
				car:   mockedCar,
				drone: mockedDrone,
			}

		})
		It("should return car chromossome list", func() {
			decodedChromossomeList := []*decodedChromossome{
				carDecodedChromossome1, carDecodedChromossome2, droneDecodedChromossome1, droneDecodedChromossome2,
			}

			receivedCarChromossomeList := collectCarChromossomes(decodedChromossomeList)

			Expect(receivedCarChromossomeList).To(HaveLen(2))
			Expect(receivedCarChromossomeList[0]).To(BeIdenticalTo(carDecodedChromossome1))
			Expect(receivedCarChromossomeList[1]).To(BeIdenticalTo(carDecodedChromossome2))
		})

		It("should return empty car chromossome list", func() {
			receivedCarChromossomeList := collectCarChromossomes([]*decodedChromossome{})

			Expect(receivedCarChromossomeList).To(BeEmpty())
		})

		It("should return empty car chromossome list when no car chromossome is present", func() {
			decodedChromossomeList := []*decodedChromossome{
				droneDecodedChromossome1, droneDecodedChromossome2,
			}

			receivedCarChromossomeList := collectCarChromossomes(decodedChromossomeList)

			Expect(receivedCarChromossomeList).To(BeEmpty())
		})
	})
})
