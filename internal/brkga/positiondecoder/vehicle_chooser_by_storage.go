package positiondecoder

import (
	"math"

	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

type vehicleChooserByStorage struct {
	carList []vehicle.ICar
	gpsMap  gps.Map

	geneAmplifier float64
	geneModule    float64
}

func NewVehicleChooserByStorage(carList []vehicle.ICar, gpsMap gps.Map) *vehicleChooserByStorage {
	ch := &vehicleChooserByStorage{
		carList: carList,
		gpsMap:  gpsMap,
	}

	ch.calcGeneAmplifier()
	ch.calcGeneModule()
	return ch
}

func (c *vehicleChooserByStorage) defineVehicle(chromossome *brkga.Chromossome) (vehicle.ICar, vehicle.IDrone) {
	modSum := 0.0
	amplifiedGene := chromossome.Gene() * c.geneAmplifier
	moduledGene := math.Mod(amplifiedGene, c.geneModule)

	for _, car := range c.carList {
		modSum += car.Storage()
		if moduledGene < modSum {
			return car, nil
		}
	}

	for _, car := range c.carList {
		for _, drone := range car.Drones() {
			modSum += drone.Storage()
			if moduledGene < modSum {
				return car, drone
			}
		}
	}

	return nil, nil
}

func (c *vehicleChooserByStorage) calcGeneAmplifier() {
	c.geneAmplifier = float64(len(c.gpsMap.Clients)) * c.calcTotalStorage()
}

func (c *vehicleChooserByStorage) calcGeneModule() {
	c.geneModule = float64(len(c.gpsMap.Clients))
}

func (c *vehicleChooserByStorage) calcTotalStorage() float64 {
	totalStorage := 0.0
	for _, car := range c.carList {
		totalStorage += car.Storage()
		for _, drone := range car.Drones() {
			totalStorage += drone.Storage()
		}
	}
	return totalStorage
}
