package strategy

import (
	"math"

	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

type vehicleChooserByPercentage struct {
	gpsMap          gps.Map
	dronePercentage float64
	carPercentage   float64
}

func NewVehicleChooserByPercentage(gpsMap gps.Map, dronePercentage float64) *vehicleChooserByPercentage {
	ch := &vehicleChooserByPercentage{
		gpsMap:          gpsMap,
		dronePercentage: dronePercentage,
		carPercentage:   1 - dronePercentage,
	}

	return ch
}

func (c *vehicleChooserByPercentage) DefineVehicle(carList []vehicle.ICar, chromossome *brkga.Chromossome) (vehicle.ICar, vehicle.IDrone) {
	moduledGene := c.calcModuledGene(chromossome)

	if moduledGene < c.carPercentage {
		return c.defineCar(carList, moduledGene), nil
	}

	return c.defineDrone(carList, moduledGene)
}

func (c *vehicleChooserByPercentage) DefineWindowTime(_ []vehicle.ICar, chromossome *brkga.Chromossome) int {
	amplifiedGene := chromossome.Gene() * c.calcGeneAmplifier()
	return int(amplifiedGene)
}

func (c *vehicleChooserByPercentage) calcModuledGene(chromossome *brkga.Chromossome) float64 {
	amplifiedGene := chromossome.Gene() * c.calcGeneAmplifier()
	return math.Mod(amplifiedGene, c.calcGeneModule())
}

func (c *vehicleChooserByPercentage) calcGeneAmplifier() float64 {
	return float64(len(c.gpsMap.Clients))
}

func (c *vehicleChooserByPercentage) calcGeneModule() float64 {
	// Returns 1.0 because we need to use the percentage of the car and drone
	return 1.0
}

func (c *vehicleChooserByPercentage) countDrones(carList []vehicle.ICar) int {
	droneCount := 0
	for _, car := range carList {
		droneCount += len(car.Drones())
	}
	return droneCount
}

func (c *vehicleChooserByPercentage) defineDrone(carList []vehicle.ICar, moduledGene float64) (vehicle.ICar, vehicle.IDrone) {
	droneCount := c.countDrones(carList)
	percentageForEachDrone := c.dronePercentage / float64(droneCount)
	dronePercentageSum := 0.0 + c.carPercentage

	for _, car := range carList {
		for _, drone := range car.Drones() {
			dronePercentageSum += percentageForEachDrone
			if dronePercentageSum >= moduledGene {
				return car, drone
			}
		}
	}

	return nil, nil
}

func (c *vehicleChooserByPercentage) defineCar(carList []vehicle.ICar, moduledGene float64) vehicle.ICar {
	carCount := len(carList)
	percentageForEachCar := c.carPercentage / float64(carCount)
	carPercentageSum := 0.0

	for _, car := range carList {
		carPercentageSum += percentageForEachCar
		if carPercentageSum >= moduledGene {
			return car
		}
	}

	return nil
}
