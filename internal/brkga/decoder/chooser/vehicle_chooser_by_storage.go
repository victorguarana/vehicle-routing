package chooser

import (
	"math"

	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

type vehicleChooserByStorage struct {
	gpsMap gps.Map
}

func NewVehicleChooserByStorage(gpsMap gps.Map) *vehicleChooserByStorage {
	ch := &vehicleChooserByStorage{
		gpsMap: gpsMap,
	}

	return ch
}

func (c *vehicleChooserByStorage) DefineVehicle(carList []vehicle.ICar, chromossome *brkga.Chromossome) (vehicle.ICar, vehicle.IDrone) {
	modSum := 0.0
	moduledGene := c.calcModuledGene(carList, chromossome)

	for _, car := range carList {
		modSum += car.Storage()
		if moduledGene < modSum {
			return car, nil
		}
	}

	for _, car := range carList {
		for _, drone := range car.Drones() {
			modSum += drone.Storage()
			if moduledGene < modSum {
				return car, drone
			}
		}
	}

	return nil, nil
}
func (c *vehicleChooserByStorage) calcModuledGene(carList []vehicle.ICar, chromossome *brkga.Chromossome) float64 {
	amplifiedGene := chromossome.Gene() * c.calcGeneAmplifier(carList)
	return math.Mod(amplifiedGene, c.calcGeneModule(carList))
}

func (c *vehicleChooserByStorage) calcGeneAmplifier(carList []vehicle.ICar) float64 {
	return float64(len(c.gpsMap.Clients)) * c.calcTotalStorage(carList)
}

func (c *vehicleChooserByStorage) calcGeneModule(carList []vehicle.ICar) float64 {
	return c.calcTotalStorage(carList)
}

func (c *vehicleChooserByStorage) calcTotalStorage(carList []vehicle.ICar) float64 {
	totalStorage := 0.0
	for _, car := range carList {
		totalStorage += car.Storage()
		for _, drone := range car.Drones() {
			totalStorage += drone.Storage()
		}
	}
	return totalStorage
}
