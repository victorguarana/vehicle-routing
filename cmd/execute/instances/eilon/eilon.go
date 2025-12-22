package eilon

import (
	"fmt"

	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

const carEfficiencyConst = 1.0
const carSpeedConst = 5.0
const droneEfficiencyConst = 5.0
const droneSpeedConst = 5.0
const rangeConst = 100000.0 // Range is not important for this instance

type fleetParams struct {
	nCars         int
	nDronesPerCar int
	carStorage    float64
	droneStorage  float64
	startingPoint gps.Point
}

func LoadEIL22() (string, gps.Map, []vehicle.ICar) {
	mapFilename := "cmd/execute/instances/eil22"
	gpsMap := gps.LoadMap(mapFilename)
	carList := createFleet(fleetParams{
		nCars:         4,
		nDronesPerCar: 2,
		carStorage:    6000.0,
		droneStorage:  1200.0,
		startingPoint: gpsMap.Warehouses[0],
	})
	return "EIL22", gpsMap, carList
}

func LoadEIL23() (string, gps.Map, []vehicle.ICar) {
	mapFilename := "cmd/execute/instances/eil23"
	gpsMap := gps.LoadMap(mapFilename)
	carList := createFleet(fleetParams{
		nCars:         3,
		nDronesPerCar: 2,
		carStorage:    4500.0,
		droneStorage:  900.0,
		startingPoint: gpsMap.Warehouses[0],
	})
	return "EIL23", gpsMap, carList
}

func LoadEIL30() (string, gps.Map, []vehicle.ICar) {
	mapFilename := "cmd/execute/instances/eil30"
	gpsMap := gps.LoadMap(mapFilename)
	carList := createFleet(fleetParams{
		nCars:         3,
		nDronesPerCar: 2,
		carStorage:    4500.0,
		droneStorage:  900.0,
		startingPoint: gpsMap.Warehouses[0],
	})
	return "EIL30", gpsMap, carList
}

func LoadEIL31() (string, gps.Map, []vehicle.ICar) {
	mapFilename := "cmd/execute/instances/eil31"
	gpsMap := gps.LoadMap(mapFilename)
	carList := createFleet(fleetParams{
		nCars:         7,
		nDronesPerCar: 2,
		carStorage:    140.0,
		droneStorage:  28.0,
		startingPoint: gpsMap.Warehouses[0],
	})
	return "EIL31", gpsMap, carList
}

func LoadEIL33() (string, gps.Map, []vehicle.ICar) {
	mapFilename := "cmd/execute/instances/eil33"
	gpsMap := gps.LoadMap(mapFilename)
	carList := createFleet(fleetParams{
		nCars:         4,
		nDronesPerCar: 2,
		carStorage:    8000.0,
		droneStorage:  1600.0,
		startingPoint: gpsMap.Warehouses[0],
	})
	return "EIL33", gpsMap, carList
}

func LoadEIL51() (string, gps.Map, []vehicle.ICar) {
	mapFilename := "cmd/execute/instances/eil51"
	gpsMap := gps.LoadMap(mapFilename)
	carList := createFleet(fleetParams{
		nCars:         5,
		nDronesPerCar: 2,
		carStorage:    160.0,
		droneStorage:  32.0,
		startingPoint: gpsMap.Warehouses[0],
	})
	return "EIL51", gpsMap, carList
}

func LoadEIL76A() (string, gps.Map, []vehicle.ICar) {
	mapFilename := "cmd/execute/instances/eil76"
	gpsMap := gps.LoadMap(mapFilename)
	carList := createFleet(fleetParams{
		nCars:         7,
		nDronesPerCar: 2,
		carStorage:    220.0,
		droneStorage:  44.0,
		startingPoint: gpsMap.Warehouses[0],
	})
	return "EIL76A", gpsMap, carList
}

func LoadEIL76B() (string, gps.Map, []vehicle.ICar) {
	mapFilename := "cmd/execute/instances/eil76"
	gpsMap := gps.LoadMap(mapFilename)
	carList := createFleet(fleetParams{
		nCars:         8,
		nDronesPerCar: 2,
		carStorage:    180.0,
		droneStorage:  36.0,
		startingPoint: gpsMap.Warehouses[0],
	})
	return "EIL76B", gpsMap, carList
}

func LoadEIL76C() (string, gps.Map, []vehicle.ICar) {
	mapFilename := "cmd/execute/instances/eil76"
	gpsMap := gps.LoadMap(mapFilename)
	carList := createFleet(fleetParams{
		nCars:         10,
		nDronesPerCar: 2,
		carStorage:    180.0,
		droneStorage:  36.0,
		startingPoint: gpsMap.Warehouses[0],
	})
	return "EIL76C", gpsMap, carList
}

func LoadEIL76D() (string, gps.Map, []vehicle.ICar) {
	mapFilename := "cmd/execute/instances/eil76"
	gpsMap := gps.LoadMap(mapFilename)
	carList := createFleet(fleetParams{
		nCars:         14,
		nDronesPerCar: 2,
		carStorage:    180.0,
		droneStorage:  36.0,
		startingPoint: gpsMap.Warehouses[0],
	})
	return "EIL76D", gpsMap, carList
}

func LoadEIL101A() (string, gps.Map, []vehicle.ICar) {
	mapFilename := "cmd/execute/instances/eil101"
	gpsMap := gps.LoadMap(mapFilename)
	carList := createFleet(fleetParams{
		nCars:         8,
		nDronesPerCar: 2,
		carStorage:    200.0,
		droneStorage:  40.0,
		startingPoint: gpsMap.Warehouses[0],
	})
	return "EIL101A", gpsMap, carList
}

func LoadEIL101B() (string, gps.Map, []vehicle.ICar) {
	mapFilename := "cmd/execute/instances/eil101"
	gpsMap := gps.LoadMap(mapFilename)
	carList := createFleet(fleetParams{
		nCars:         14,
		nDronesPerCar: 2,
		carStorage:    200.0,
		droneStorage:  40.0,
		startingPoint: gpsMap.Warehouses[0],
	})
	return "EIL101B", gpsMap, carList
}

func createFleet(fp fleetParams) []vehicle.ICar {
	cars := make([]vehicle.ICar, fp.nCars)
	for i := 0; i < fp.nCars; i++ {
		car := vehicle.NewCarLimited(vehicle.CarParams{
			Efficiency:    carEfficiencyConst,
			Speed:         carSpeedConst,
			Storage:       fp.carStorage,
			Range:         rangeConst,
			Name:          "car" + fmt.Sprint(i),
			StartingPoint: fp.startingPoint,
		})
		for j := 0; j < fp.nDronesPerCar; j++ {
			car.NewDroneWithParams(
				vehicle.DroneParams{
					Efficiency:    droneEfficiencyConst,
					Speed:         droneSpeedConst,
					Storage:       fp.droneStorage,
					Range:         rangeConst,
					Name:          "drone" + fmt.Sprint(j) + "_of_car" + fmt.Sprint(i),
					StartingPoint: fp.startingPoint,
				},
			)
		}
		cars[i] = car
	}
	return cars
}
