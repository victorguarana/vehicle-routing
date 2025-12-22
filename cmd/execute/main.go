package main

import (
	"log"
	"os"

	"github.com/victorguarana/vehicle-routing/cmd/execute/instances/eilon"
	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/brkga/decoder"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/measure"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

var dronePercentage = 0.0
var iterations = 10

var brkgaBaseParams = brkga.BRKGAParams[itinerary.ItineraryList]{
	BiasPercentage:      0.75,
	CrossoverPercentage: 0.6,
	TopPercentage:       0.1,
	MaxPop:              100,
	GenerationLimit:     2000,
}

var distanceMeasurer = measure.NewMeasurer(measure.TotalDistance, "TotalDistance")
var fuelMeasurer = measure.NewMeasurer(measure.SpentFuel, "SpentFuel")
var timeMeasurer = measure.NewMeasurer(measure.TimeSpent, "TimeSpent")

func main() {
	f, err := os.OpenFile("resultados_tabela4_semdrone.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Println("Bias Percentage", brkgaBaseParams.BiasPercentage)
	log.Println("Crossover Percentage", brkgaBaseParams.CrossoverPercentage)
	log.Println("Top Percentage", brkgaBaseParams.TopPercentage)
	log.Println("Max Population", brkgaBaseParams.MaxPop)
	log.Println("Generation Limit", brkgaBaseParams.GenerationLimit)
	log.Println("Drone Percentage", dronePercentage)

	log.Println("Starting execution...")
	executeBRKGA(eilon.LoadEIL22())
	executeBRKGA(eilon.LoadEIL23())
	executeBRKGA(eilon.LoadEIL30())
	executeBRKGA(eilon.LoadEIL33())
	executeBRKGA(eilon.LoadEIL51())
	executeBRKGA(eilon.LoadEIL76A())
	executeBRKGA(eilon.LoadEIL76B())
	executeBRKGA(eilon.LoadEIL76C())
	executeBRKGA(eilon.LoadEIL76D())
	executeBRKGA(eilon.LoadEIL101A())
	executeBRKGA(eilon.LoadEIL101B())
	log.Println("Finishing execution...")
}

func executeBRKGA(mapFilename string, gpsMap gps.Map, carList []vehicle.ICar) {
	BRKGA(
		measure.NewMeasurer(measure.TotalDistance, "TotalDistance"),
		decoder.NewPositionalDecoderWithVehicleByPercentage(carList, gpsMap, dronePercentage),
		gpsMap,
		mapFilename,
		"positional_by_percentage_total_distance")

	BRKGA(
		measure.NewMeasurer(measure.TotalDistance, "TotalDistance"),
		decoder.NewPositionalDecoderWithVehicleByStorage(carList, gpsMap),
		gpsMap,
		mapFilename,
		"positional_by_storage_total_distance")

	BRKGA(
		measure.NewMeasurer(measure.TotalDistance, "TotalDistance"),
		decoder.NewTimeDecoderWithVehicleByPercentage(carList, gpsMap, dronePercentage),
		gpsMap,
		mapFilename,
		"time_by_percentage_total_distance")

	BRKGA(
		measure.NewMeasurer(measure.TotalDistance, "TotalDistance"),
		decoder.NewTimeDecoderWithVehicleByStorage(carList, gpsMap),
		gpsMap,
		mapFilename,
		"time_by_storage_total_distance")

	BRKGA(
		measure.NewMeasurer(measure.SpentFuel, "SpentFuel"),
		decoder.NewPositionalDecoderWithVehicleByPercentage(carList, gpsMap, dronePercentage),
		gpsMap,
		mapFilename,
		"positional_by_percentage_fuel_spent")

	BRKGA(
		measure.NewMeasurer(measure.SpentFuel, "SpentFuel"),
		decoder.NewPositionalDecoderWithVehicleByStorage(carList, gpsMap),
		gpsMap,
		mapFilename,
		"positional_by_storage_fuel_spent")

	BRKGA(
		measure.NewMeasurer(measure.SpentFuel, "SpentFuel"),
		decoder.NewTimeDecoderWithVehicleByPercentage(carList, gpsMap, dronePercentage),
		gpsMap,
		mapFilename,
		"time_by_percentage_fuel_spent")

	BRKGA(
		measure.NewMeasurer(measure.SpentFuel, "SpentFuel"),
		decoder.NewTimeDecoderWithVehicleByStorage(carList, gpsMap),
		gpsMap,
		mapFilename,
		"time_by_storage_fuel_spent")

	BRKGA(
		measure.NewMeasurer(measure.TimeSpent, "TimeSpent"),
		decoder.NewPositionalDecoderWithVehicleByPercentage(carList, gpsMap, dronePercentage),
		gpsMap,
		mapFilename,
		"positional_by_percentage_time_spent")

	BRKGA(
		measure.NewMeasurer(measure.TimeSpent, "TimeSpent"),
		decoder.NewPositionalDecoderWithVehicleByStorage(carList, gpsMap),
		gpsMap,
		mapFilename,
		"positional_by_storage_time_spent")

	BRKGA(
		measure.NewMeasurer(measure.TimeSpent, "TimeSpent"),
		decoder.NewTimeDecoderWithVehicleByPercentage(carList, gpsMap, dronePercentage),
		gpsMap,
		mapFilename,
		"time_by_percentage_time_spent")

	BRKGA(
		measure.NewMeasurer(measure.TimeSpent, "TimeSpent"),
		decoder.NewTimeDecoderWithVehicleByStorage(carList, gpsMap),
		gpsMap,
		mapFilename,
		"time_by_storage_time_spent")

	// wg.Wait()
}

func BRKGA(m measure.Measurer, d brkga.IDecoder[itinerary.ItineraryList], gpsMap gps.Map, mapFilename string, nameSuffix string) {
	// defer wg.Done()
	for i := 0; i < iterations; i++ {
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
			log.Println("No solution found for", mapFilename, d.Name())
			continue
		}

		score := m.Measure(itn)
		log.Printf("%s %s (%s): %d\n", mapFilename, d.Name(), m.Name(), int(score))

		score = distanceMeasurer.Measure(itn)
		log.Printf("%s %s (%s - %s): %d\n", mapFilename, d.Name(), m.Name(), "Total Distance", int(score))

		score = timeMeasurer.Measure(itn)
		log.Printf("%s %s (%s - %s): %d\n", mapFilename, d.Name(), m.Name(), "Time Spent", int(score))

		score = fuelMeasurer.Measure(itn)
		log.Printf("%s %s (%s - %s): %d\n", mapFilename, d.Name(), m.Name(), "Fuel Spent", int(score))
	}
}
