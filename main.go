package main

import (
	"fmt"
	"log"

	"github.com/victorguarana/vehicle-routing/internal/csp"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/greedy"
	"github.com/victorguarana/vehicle-routing/internal/ils"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/measure"
	"github.com/victorguarana/vehicle-routing/internal/output"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

const mapFilename = "example/map"

var allMeasures = map[string]func(itinerary.Info) float64{
	"Total Distance": measure.TotalDistance,
	"Total Time":     measure.TimeSpent,
	"Total Fuel":     measure.SpentFuel,
}

func main() {
	BestInsertion()
	BestInsertionWithDrones()
	BestInsertionWithDronesShiftC2D()
	BestInsertionWithDronesShiftD2C()
	BestInsertionWithDronesSwapCD()

	ClosestNeighbor()
	ClosestNeighborWithDrones()
	ClosestNeighborWithDronesShiftC2D()
	ClosestNeighborWithDronesShiftD2C()
	ClosestNeighborWithDronesSwapCD()

	Covering()
	CoveringMaxDrones()
}

func ClosestNeighbor() {
	gpsMap, itn := loadEnvironment()
	constructor := itn.Constructor()
	greedy.ClosestNeighbor([]itinerary.Constructor{constructor}, gpsMap)

	itnInfo := itn.Info()
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s_closest_neighbor.png", mapFilename)
	output.ToImage(filename, itnInfo, outputInfos)
}

func ClosestNeighborWithDrones() {
	gpsMap, itn := loadEnvironment()
	constructor := itn.Constructor()
	modifier := itn.Modifier()
	greedy.ClosestNeighbor([]itinerary.Constructor{constructor}, gpsMap)
	greedy.DroneStrikesInsertion(constructor, modifier)

	itnInfo := itn.Info()
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s_closest_neighbor_with_drones.png", mapFilename)
	output.ToImage(filename, itnInfo, outputInfos)
}

func ClosestNeighborWithDronesShiftC2D() {
	gpsMap, itn := loadEnvironment()
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
	filename := fmt.Sprintf("%s_closest_neighbor_with_drones_shift_c2d.png", mapFilename)
	output.ToImage(filename, itnInfo, outputInfos)
}

func ClosestNeighborWithDronesShiftD2C() {
	gpsMap, itn := loadEnvironment()
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
	filename := fmt.Sprintf("%s_closest_neighbor_with_drones_shift_d2c.png", mapFilename)
	output.ToImage(filename, itnInfo, outputInfos)
}

func ClosestNeighborWithDronesSwapCD() {
	gpsMap, itn := loadEnvironment()
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
	filename := fmt.Sprintf("%s_closest_neighbor_with_drones_swap_cd.png", mapFilename)
	output.ToImage(filename, itnInfo, outputInfos)
}

func BestInsertion() {
	gpsMap, itn := loadEnvironment()
	constructor := itn.Constructor()
	greedy.BestInsertion([]itinerary.Constructor{constructor}, gpsMap)

	itnInfo := itn.Info()
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s_best_insertion.png", mapFilename)
	output.ToImage(filename, itnInfo, outputInfos)
}

func BestInsertionWithDrones() {
	gpsMap, itn := loadEnvironment()
	constructor := itn.Constructor()
	modifier := itn.Modifier()
	greedy.BestInsertion([]itinerary.Constructor{constructor}, gpsMap)
	greedy.DroneStrikesInsertion(constructor, modifier)

	itnInfo := itn.Info()
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s_best_insertion_with_drones.png", mapFilename)
	output.ToImage(filename, itnInfo, outputInfos)
}

func BestInsertionWithDronesShiftC2D() {
	gpsMap, itn := loadEnvironment()
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
	filename := fmt.Sprintf("%s_best_insertion_with_drones_shift_c2d.png", mapFilename)
	output.ToImage(filename, itnInfo, outputInfos)
}

func BestInsertionWithDronesShiftD2C() {
	gpsMap, itn := loadEnvironment()
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
	filename := fmt.Sprintf("%s_best_insertion_with_drones_shift_d2c.png", mapFilename)
	output.ToImage(filename, itnInfo, outputInfos)
}

func BestInsertionWithDronesSwapCD() {
	gpsMap, itn := loadEnvironment()
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
	filename := fmt.Sprintf("%s_best_insertion_with_drones_swap_cd.png", mapFilename)
	output.ToImage(filename, itnInfo, outputInfos)
}

func Covering() {
	gpsMap, itn := loadEnvironment()
	neighorhoodDistance := vehicle.DroneRange / 4
	constructor := itn.Constructor()
	csp.CoveringWithDrones([]itinerary.Constructor{constructor}, gpsMap, neighorhoodDistance)

	itnInfo := itn.Info()
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s_covering.png", mapFilename)
	output.ToImage(filename, itnInfo, outputInfos)
}

func CoveringMaxDrones() {
	gpsMap, itn := loadEnvironment()
	neighorhoodDistance := vehicle.DroneRange / 2
	constructor := itn.Constructor()
	csp.CoveringWithDrones([]itinerary.Constructor{constructor}, gpsMap, neighorhoodDistance)

	itnInfo := itn.Info()
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s_covering_max_drones.png", mapFilename)
	output.ToImage(filename, itnInfo, outputInfos)
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

func loadEnvironment() (gps.Map, itinerary.Itinerary) {
	gpsMap := gps.LoadMap(mapFilename)
	initialPoint := gpsMap.Warehouses[0]
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)

	return gpsMap, itn
}
