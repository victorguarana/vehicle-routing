package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/victorguarana/vehicle-routing/internal/csp"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/greedy"
	"github.com/victorguarana/vehicle-routing/internal/ils"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/measure"
	"github.com/victorguarana/vehicle-routing/internal/output"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

const mapFilename = "maps/map1"

var allMeasures = map[string]func(itinerary.Info) float64{
	"Total Distance": measure.TotalDistance,
	"Total Time":     measure.TimeSpent,
	"Total Fuel":     measure.SpentFuel,
}

func main() {
	BestInsertion(mapFilename)
	BestInsertionWithDrones(mapFilename)
	BestInsertionWithDronesShiftC2D(mapFilename)
	BestInsertionWithDronesShiftD2C(mapFilename)
	BestInsertionWithDronesSwapCD(mapFilename)

	ClosestNeighbor(mapFilename)
	ClosestNeighborWithDrones(mapFilename)
	ClosestNeighborWithDronesShiftC2D(mapFilename)
	ClosestNeighborWithDronesShiftD2C(mapFilename)
	ClosestNeighborWithDronesSwapCD(mapFilename)

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
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s_closest_neighbor.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, outputInfos)
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
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s_closest_neighbor_with_drones.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, outputInfos)
}

func ClosestNeighborWithDronesShiftC2D(mapFilename string) {
	gpsMap := gps.LoadMap(mapFilename)
	initialPoint := gpsMap.Warehouses[0]
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	constructor := itn.Constructor()
	modifier := itn.Modifier()
	finder := itn.Finder()
	greedy.ClosestNeighbor([]itinerary.Constructor{constructor}, gpsMap)
	greedy.DroneStrikesInsertion(constructor, modifier)

	err := ils.ShiftCarToDrone(modifier, finder)
	if err != nil {
		log.Println("ClosestNeighborWithDronesShiftC2D:", err)
		return
	}

	itnInfo := itn.Info()
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s-closest-neighbor-with-drones-shift-c2d.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, outputInfos)
}

func ClosestNeighborWithDronesShiftD2C(mapFilename string) {
	gpsMap := gps.LoadMap(mapFilename)
	initialPoint := gpsMap.Warehouses[0]
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	constructor := itn.Constructor()
	modifier := itn.Modifier()
	finder := itn.Finder()
	greedy.ClosestNeighbor([]itinerary.Constructor{constructor}, gpsMap)
	greedy.DroneStrikesInsertion(constructor, modifier)

	err := ils.ShiftDroneToCar(modifier, finder)
	if err != nil {
		log.Println("ClosestNeighborWithDronesShiftD2C:", err)
		return
	}

	itnInfo := itn.Info()
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s-closest-neighbor-with-drones-shift-d2c.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, outputInfos)
}

func ClosestNeighborWithDronesSwapCD(mapFilename string) {
	gpsMap := gps.LoadMap(mapFilename)
	initialPoint := gpsMap.Warehouses[0]
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	constructor := itn.Constructor()
	modifier := itn.Modifier()
	finder := itn.Finder()
	greedy.ClosestNeighbor([]itinerary.Constructor{constructor}, gpsMap)
	greedy.DroneStrikesInsertion(constructor, modifier)

	err := ils.SwapCarAndDrone(modifier, finder)
	if err != nil {
		log.Println("ClosestNeighborWithDronesSwapCD:", err)
		return
	}

	itnInfo := itn.Info()
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s-closest-neighbor-with-drones-swap-cd.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, outputInfos)
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
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s_best_insertion.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, outputInfos)
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
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s_best_insertion_with_drones.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, outputInfos)
}

func BestInsertionWithDronesShiftC2D(mapFilename string) {
	gpsMap := gps.LoadMap(mapFilename)
	initialPoint := gpsMap.Warehouses[0]
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	constructor := itn.Constructor()
	modifier := itn.Modifier()
	finder := itn.Finder()
	greedy.BestInsertion([]itinerary.Constructor{constructor}, gpsMap)
	greedy.DroneStrikesInsertion(constructor, modifier)

	err := ils.ShiftCarToDrone(modifier, finder)
	if err != nil {
		log.Println("BestInsertionWithDronesShiftC2D:", err)
		return
	}

	itnInfo := itn.Info()
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s-best-insertion-with-drones-shift-c2d.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, outputInfos)
}

func BestInsertionWithDronesShiftD2C(mapFilename string) {
	gpsMap := gps.LoadMap(mapFilename)
	initialPoint := gpsMap.Warehouses[0]
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	constructor := itn.Constructor()
	modifier := itn.Modifier()
	finder := itn.Finder()
	greedy.BestInsertion([]itinerary.Constructor{constructor}, gpsMap)
	greedy.DroneStrikesInsertion(constructor, modifier)

	err := ils.ShiftDroneToCar(modifier, finder)
	if err != nil {
		log.Println("BestInsertionWithDronesShiftD2C:", err)
	}

	itnInfo := itn.Info()
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s-best-insertion-with-drones-shift-d2c.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, outputInfos)
}

func BestInsertionWithDronesSwapCD(mapFilename string) {
	gpsMap := gps.LoadMap(mapFilename)
	initialPoint := gpsMap.Warehouses[0]
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	constructor := itn.Constructor()
	modifier := itn.Modifier()
	finder := itn.Finder()
	greedy.BestInsertion([]itinerary.Constructor{constructor}, gpsMap)
	greedy.DroneStrikesInsertion(constructor, modifier)

	err := ils.SwapCarAndDrone(modifier, finder)
	if err != nil {
		log.Println("BestInsertionWithDronesSwapCD:", err)
	}

	itnInfo := itn.Info()
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s-best-insertion-with-drones-swap-cd.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, outputInfos)
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
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s_covering.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, outputInfos)
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
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s_covering_max_drones.png", removeExtentionFromFilename(mapFilename))
	output.ToImage(filename, itnInfo, outputInfos)
}

func removeExtentionFromFilename(filename string) string {
	return strings.Trim(strings.TrimSuffix(filename, ".txt"), ".csv")
}

func mountOutputInfo(itnInfo itinerary.Info) []output.Info {
	var infos []output.Info
	for measureName, measureFunc := range allMeasures {
		measureValue := measureFunc(itnInfo)
		measureStr := fmt.Sprintf("%s: %.2f", measureName, measureValue)
		infos = append(infos, output.Info{Str: measureStr})

		log.Println(measureName, measureValue)
	}

	return infos
}
