package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/brkga/decoder"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/measure"
	"github.com/victorguarana/vehicle-routing/internal/output"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

const mapFilename = "map_r101_25"

var allMeasures = map[string]func(itinerary.Info) float64{
	"Total Distance": measure.TotalDistance,
	"Total Time":     measure.TimeSpent,
	"Total Fuel":     measure.SpentFuel,
}

var brkgaBaseParams = brkga.BRKGAParams[itinerary.ItineraryList]{
	MaxPop:              100,
	TopPercentage:       0.2,
	CrossoverPercentage: 0.5,
	BiasPercentage:      0.7,
	GenerationLimit:     10000,
}

var wg sync.WaitGroup

func main() {
	executeBRKGA()
}

func executeBRKGA() {
	gpsMap, carList := loadBRKGAEnvironment()

	wg.Add(1)
	go BRKGA(
		measure.NewMeasurer(measure.TimeSpent, "TimeSpent"),
		decoder.NewPositionalDecoderWithVehicleByStorage(carList, gpsMap),
		gpsMap,
		"positional_by_storage_time_spent")

	wg.Add(1)
	go BRKGA(
		measure.NewMeasurer(measure.TotalDistance, "TotalDistance"),
		decoder.NewPositionalDecoderWithVehicleByStorage(carList, gpsMap),
		gpsMap,
		"positional_by_storage_total_distance")

	wg.Add(1)
	go BRKGA(
		measure.NewMeasurer(measure.SpentFuel, "SpentFuel"),
		decoder.NewPositionalDecoderWithVehicleByStorage(carList, gpsMap),
		gpsMap,
		"positional_by_storage_fuel_spent")

	wg.Add(1)
	go BRKGA(
		measure.NewMeasurer(measure.TimeSpent, "TimeSpent"),
		decoder.NewPositionalDecoderWithVehicleByPercentage(carList, gpsMap, 0.15),
		gpsMap,
		"positional_by_percentage_time_spent")

	wg.Add(1)
	go BRKGA(
		measure.NewMeasurer(measure.TotalDistance, "TotalDistance"),
		decoder.NewPositionalDecoderWithVehicleByPercentage(carList, gpsMap, 0.15),
		gpsMap,
		"positional_by_percentage_total_distance")

	wg.Add(1)
	go BRKGA(
		measure.NewMeasurer(measure.SpentFuel, "SpentFuel"),
		decoder.NewPositionalDecoderWithVehicleByPercentage(carList, gpsMap, 0.15),
		gpsMap,
		"positional_by_percentage_fuel_spent")

	wg.Add(1)
	go BRKGA(
		measure.NewMeasurer(measure.TimeSpent, "TimeSpent"),
		decoder.NewTimeDecoderWithVehicleByStorage(carList, gpsMap),
		gpsMap,
		"time_by_storage_time_spent")

	wg.Add(1)
	go BRKGA(
		measure.NewMeasurer(measure.TotalDistance, "TotalDistance"),
		decoder.NewTimeDecoderWithVehicleByStorage(carList, gpsMap),
		gpsMap,
		"time_by_storage_total_distance")

	wg.Add(1)
	go BRKGA(
		measure.NewMeasurer(measure.SpentFuel, "SpentFuel"),
		decoder.NewTimeDecoderWithVehicleByStorage(carList, gpsMap),
		gpsMap,
		"time_by_storage_fuel_spent")

	wg.Add(1)
	go BRKGA(
		measure.NewMeasurer(measure.TimeSpent, "TimeSpent"),
		decoder.NewTimeDecoderWithVehicleByPercentage(carList, gpsMap, 0.15),
		gpsMap,
		"time_by_percentage_time_spent")

	wg.Add(1)
	go BRKGA(
		measure.NewMeasurer(measure.TotalDistance, "TotalDistance"),
		decoder.NewTimeDecoderWithVehicleByPercentage(carList, gpsMap, 0.15),
		gpsMap,
		"time_by_percentage_total_distance")

	wg.Add(1)
	go BRKGA(
		measure.NewMeasurer(measure.SpentFuel, "SpentFuel"),
		decoder.NewTimeDecoderWithVehicleByPercentage(carList, gpsMap, 0.15),
		gpsMap,
		"time_by_percentage_fuel_spent")

	wg.Wait()
}

func BRKGA(m measure.Measurer, d brkga.IDecoder[itinerary.ItineraryList], gpsMap gps.Map, nameSuffix string) {
	defer wg.Done()
	itn := brkga.NewBRKGA(brkga.BRKGAParams[itinerary.ItineraryList]{
		MaxPop:              brkgaBaseParams.MaxPop,
		TopPercentage:       brkgaBaseParams.TopPercentage,
		CrossoverPercentage: brkgaBaseParams.CrossoverPercentage,
		BiasPercentage:      brkgaBaseParams.BiasPercentage,
		GenerationLimit:     brkgaBaseParams.GenerationLimit,
		ChromosomeLen:       len(gpsMap.Customers),
		Decoder:             d,
		Measurer:            m,
		OptimizationGoal:    brkga.Minimize,
	}).Execute()

	if itn == nil {
		fmt.Println("No solution found for", d.Name())
		return
	}

	score := m.Measure(itn)
	fmt.Printf("BRKGA %s: %.4f (%s) \n", d.Name(), score, m.Name())

	itnInfo := itn[0].Info()
	outputInfos := mountOutputInfo(itnInfo)
	filename := fmt.Sprintf("%s_brkga_%s.png", mapFilename, nameSuffix)
	output.ToImage(filename, itnInfo, outputInfos)
}

func loadBRKGAEnvironment() (gps.Map, []vehicle.ICar) {
	gpsMap := gps.LoadMap(mapFilename)
	car1 := vehicle.NewCarLimited(vehicle.CarParams{
		Efficiency:    1.0,
		Speed:         5.0,
		Storage:       200.0,
		Range:         700.0,
		Name:          "car1",
		StartingPoint: gpsMap.Warehouses[0],
	})
	car1.NewDroneWithParams(
		vehicle.DroneParams{
			Efficiency:    5.0,
			Speed:         5.0,
			Storage:       100.0,
			Range:         50.0,
			Name:          "drone1",
			StartingPoint: gpsMap.Warehouses[0],
		},
	)
	car2 := vehicle.NewCarLimited(vehicle.CarParams{
		Efficiency:    1.0,
		Speed:         5.0,
		Storage:       200.0,
		Range:         700.0,
		Name:          "car2",
		StartingPoint: gpsMap.Warehouses[0],
	})
	car2.NewDroneWithParams(
		vehicle.DroneParams{
			Efficiency:    5.0,
			Speed:         5.0,
			Storage:       100.0,
			Range:         50.0,
			Name:          "drone2",
			StartingPoint: gpsMap.Warehouses[0],
		},
	)
	return gpsMap, []vehicle.ICar{car1, car2}
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
