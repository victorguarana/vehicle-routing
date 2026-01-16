package main

import (
	"log"
	"os"

	"github.com/victorguarana/vehicle-routing/cmd/execute/instances/agatz"
	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/brkga/decoder"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/measure"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

var dronePercentage = 0.8
var iterations = 10
var alpha agatz.Alpha = agatz.Alpha3

var brkgaBaseParams = brkga.BRKGAParams[itinerary.ItineraryList]{
	BiasPercentage:      0.75,
	CrossoverPercentage: 0.6,
	TopPercentage:       0.1,
	MaxPop:              100,
	GenerationLimit:     2000,
}

var distanceMeasurer = measure.NewMeasurer(measure.TotalDistance, "TotalDistance")

var timeMeasurer = measure.NewMeasurer(measure.TimeSpent, "TimeSpent")

func main() {
	f, err := os.OpenFile("resultados_agatz_uniform_alpha3_08.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
	log.Println("Alpha", alpha)

	log.Println("Starting execution...")

	// Agatz Uniform N10 Instances Alpha1
	for name, mapFilename := range agatz.AgatzInstanceUniformN10 {
		executeBRKGA(agatz.LoadAgatz(mapFilename, name, alpha))
	}
	for name, mapFilename := range agatz.AgatzInstanceUniformN20 {
		executeBRKGA(agatz.LoadAgatz(mapFilename, name, alpha))
	}
	for name, mapFilename := range agatz.AgatzInstanceUniformN50 {
		executeBRKGA(agatz.LoadAgatz(mapFilename, name, alpha))
	}
	for name, mapFilename := range agatz.AgatzInstanceUniformN75 {
		executeBRKGA(agatz.LoadAgatz(mapFilename, name, alpha))
	}
	for name, mapFilename := range agatz.AgatzInstanceUniformN100 {
		executeBRKGA(agatz.LoadAgatz(mapFilename, name, alpha))
	}
	for name, mapFilename := range agatz.AgatzInstanceUniformN175 {
		executeBRKGA(agatz.LoadAgatz(mapFilename, name, alpha))
	}
	for name, mapFilename := range agatz.AgatzInstanceUniformN250 {
		executeBRKGA(agatz.LoadAgatz(mapFilename, name, alpha))
	}

	log.Println("Finishing execution...")
}

func executeBRKGA(mapFilename string, gpsMap gps.Map, carList []vehicle.ICar) {
	BRKGA(
		measure.NewMeasurer(measure.TimeSpent, "TimeSpent"),
		decoder.NewPositionalDecoderWithVehicleByPercentage(carList, gpsMap, dronePercentage),
		gpsMap,
		mapFilename,
		"positional_by_percentage_time_spent")

	BRKGA(
		measure.NewMeasurer(measure.TimeSpent, "TimeSpent"),
		decoder.NewTimeDecoderWithVehicleByPercentage(carList, gpsMap, dronePercentage),
		gpsMap,
		mapFilename,
		"time_by_percentage_time_spent")
}

func BRKGA(m measure.Measurer, d brkga.IDecoder[itinerary.ItineraryList], gpsMap gps.Map, mapFilename string, nameSuffix string) {
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

		score := distanceMeasurer.Measure(itn)
		log.Printf("%s %s (%s - %s): %.2f\n", mapFilename, d.Name(), m.Name(), "Total Distance", score)

		score = timeMeasurer.Measure(itn)
		log.Printf("%s %s (%s - %s): %.2f\n", mapFilename, d.Name(), m.Name(), "Time Spent", score)
	}
}
