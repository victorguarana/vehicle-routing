package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/victorguarana/vehicle-routing/src/csp"
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/greedy"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/measure"
	"github.com/victorguarana/vehicle-routing/src/output"
	"github.com/victorguarana/vehicle-routing/src/vehicle"
)

const mapFilename = "maps/map1.csv"

func main() {
	BestInsertion(mapFilename)
	BestInsertionWithDrones(mapFilename)

	ClosestNeighbor(mapFilename)
	ClosestNeighborWithDrones(mapFilename)

	Covering(mapFilename)
	CoveringMaxDrones(mapFilename)
}

func ClosestNeighbor(mapFilename string) {
	gpsMap := gps.LoadMap(mapFilename)
	initialPoint := gpsMap.Warehouses[0]
	car := vehicle.NewCar("car1", initialPoint)
	itn := itinerary.New(car)
	constructor := itn.Constructor()
	greedy.ClosestNeighbor([]itinerary.Constructor{constructor}, gpsMap)

	itnInfo := itn.Info()
	totalDistance := measure.TotalDistance(itnInfo)
	totalTime := measure.TimeSpent(itnInfo)
	totalFuelSpent := measure.SpentFuel(itnInfo)
	log.Println("ClosestNeighbor: Total Distance:", totalDistance)
	log.Println("ClosestNeighbor: Total Time:", totalTime)
	log.Println("ClosestNeighbor: Total Fuel Spent:", totalFuelSpent)

	filename := fmt.Sprintf("%s_closest_neighbor.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, totalDistance, totalTime)
}

func ClosestNeighborWithDrones(mapFilename string) {
	gpsMap := gps.LoadMap(mapFilename)
	initialPoint := gpsMap.Warehouses[0]
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	constructor := itn.Constructor()
	modifier := itn.Modifier()
	greedy.ClosestNeighbor([]itinerary.Constructor{constructor}, gpsMap)
	greedy.DroneStrikesInsertion(constructor, modifier)

	itnInfo := itn.Info()
	totalDistance := measure.TotalDistance(itnInfo)
	totalTime := measure.TimeSpent(itnInfo)
	totalFuelSpent := measure.SpentFuel(itnInfo)
	log.Println("ClosestNeighborWithDrones: Total Distance:", totalDistance)
	log.Println("ClosestNeighborWithDrones: Total Time:", totalTime)
	log.Println("ClosestNeighborWithDrones: Total Fuel Spent:", totalFuelSpent)

	filename := fmt.Sprintf("%s_closest_neighbor_with_drones.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, totalDistance, totalTime)
}

func BestInsertion(mapFilename string) {
	gpsMap := gps.LoadMap(mapFilename)
	initialPoint := gpsMap.Warehouses[0]
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	constructor := itn.Constructor()
	greedy.BestInsertion([]itinerary.Constructor{constructor}, gpsMap)

	itnInfo := itn.Info()
	totalDistance := measure.TotalDistance(itnInfo)
	totalTime := measure.TimeSpent(itnInfo)
	totalFuelSpent := measure.SpentFuel(itnInfo)
	log.Println("BestInsertion: Total Distance:", totalDistance)
	log.Println("BestInsertion: Total Time:", totalTime)
	log.Println("BestInsertion: Total Fuel Spent:", totalFuelSpent)

	filename := fmt.Sprintf("%s_best_insertion.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, totalDistance, totalTime)
}

func BestInsertionWithDrones(mapFilename string) {
	gpsMap := gps.LoadMap(mapFilename)
	initialPoint := gpsMap.Warehouses[0]
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	constructor := itn.Constructor()
	modifier := itn.Modifier()
	greedy.BestInsertion([]itinerary.Constructor{constructor}, gpsMap)
	greedy.DroneStrikesInsertion(constructor, modifier)

	itnInfo := itn.Info()
	totalDistance := measure.TotalDistance(itnInfo)
	totalTime := measure.TimeSpent(itnInfo)
	totalFuelSpent := measure.SpentFuel(itnInfo)
	log.Println("BestInsertionWithDrones: Total Distance:", totalDistance)
	log.Println("BestInsertionWithDrones: Total Time:", totalTime)
	log.Println("BestInsertionWithDrones: Total Fuel Spent:", totalFuelSpent)

	filename := fmt.Sprintf("%s_best_insertion_with_drones.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, totalDistance, totalTime)
}

func Covering(mapFilename string) {
	gpsMap := gps.LoadMap(mapFilename)
	initialPoint := gpsMap.Warehouses[0]
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	neighorhoodDistance := vehicle.DroneRange / 4
	constructor := itn.Constructor()
	csp.CoveringWithDrones([]itinerary.Constructor{constructor}, gpsMap, neighorhoodDistance)

	itnInfo := itn.Info()
	totalDistance := measure.TotalDistance(itnInfo)
	totalTime := measure.TimeSpent(itnInfo)
	totalFuelSpent := measure.SpentFuel(itnInfo)
	log.Println("Covering: Total Distance:", totalDistance)
	log.Println("Covering: Total Time:", totalTime)
	log.Println("Covering: Total Fuel Spent:", totalFuelSpent)

	filename := fmt.Sprintf("%s_covering.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, totalDistance, totalTime)
}

func CoveringMaxDrones(mapFilename string) {
	gpsMap := gps.LoadMap(mapFilename)
	initialPoint := gpsMap.Warehouses[0]
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	neighorhoodDistance := vehicle.DroneRange / 2
	constructor := itn.Constructor()
	csp.CoveringWithDrones([]itinerary.Constructor{constructor}, gpsMap, neighorhoodDistance)

	itnInfo := itn.Info()
	totalDistance := measure.TotalDistance(itnInfo)
	totalTime := measure.TimeSpent(itnInfo)
	totalFuelSpent := measure.SpentFuel(itnInfo)
	log.Println("CoveringMaxDrones: Total Distance:", totalDistance)
	log.Println("CoveringMaxDrones: Total Time:", totalTime)
	log.Println("CoveringMaxDrones: Total Fuel Spent:", totalFuelSpent)

	filename := fmt.Sprintf("%s_covering_max_drones.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, totalDistance, totalTime)
}

func removeExtentionFromFilename(filename string) string {
	return strings.Trim(strings.TrimSuffix(filename, ".txt"), ".csv")
}
