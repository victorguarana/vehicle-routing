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

var dronePercentage = 0.5
var iterations = 10

var alpha agatz.Alpha = agatz.Alpha3

// var allMeasures = map[string]func(itinerary.Info) float64{
// 	// "Total Distance": measure.TotalDistance,
// 	"Total Time": measure.TimeSpent,
// 	// "Total Fuel":     measure.SpentFuel,
// }

var brkgaBaseParams = brkga.BRKGAParams[itinerary.ItineraryList]{
	BiasPercentage:      0.75,
	CrossoverPercentage: 0.6,
	TopPercentage:       0.1,
	MaxPop:              100,
	GenerationLimit:     2000,
}

// var distanceMeasurer = measure.NewMeasurer(measure.TotalDistance, "TotalDistance")

// var timeMeasurer = measure.NewMeasurer(measure.TimeSpent, "TimeSpent")
var timeAgatzMeasurer = measure.NewMeasurer(measure.TimeSpentAgatz, "TimeSpent")

func main() {
	log.Println("Alpha", alpha)
	log.Println("Drone Percentage", dronePercentage)

	f, err := os.OpenFile("logs/agatz/ms/resultados_mothership_agatz_uniform_alpha3_05.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	// Agatz Instances
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
		measure.NewMeasurer(measure.TimeSpentAgatz, "TimeSpent"),
		decoder.NewPositionalDecoderWithVehicleByPercentageV2(carList, gpsMap, dronePercentage),
		gpsMap,
		mapFilename,
		"positional_agatz_by_percentage_time_spent")
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

		// score := m.Measure(itn)
		// log.Printf("%s %s (%s): %.2f\n", mapFilename, d.Name(), m.Name(), score)

		// score := distanceMeasurer.Measure(itn)
		// log.Printf("%s %s (%s - %s): %.2f\n", mapFilename, d.Name(), m.Name(), "Total Distance", score)

		// outputInfos := mountOutputInfo(itn[0].Info())
		// filename := fmt.Sprintf("%s_teste_%d_debug.png", mapFilename, i)
		// output.ToImage(filename, itn[0].Info(), outputInfos)

		score := timeAgatzMeasurer.Measure(itn)
		log.Printf("%s %s (%s - %s): %.2f\n", mapFilename, d.Name(), m.Name(), "Time Spent", score)
		// timeMeasurer.Measure(itn)
	}
}

// func mountOutputInfo(itnInfo itinerary.Info) []output.Info {
// 	var infos []output.Info
// 	for measureName, measureFunc := range allMeasures {
// 		measureValue := measureFunc(itnInfo)
// 		measureStr := fmt.Sprintf("%s: %.2f", measureName, measureValue)
// 		infos = append(infos, output.Info{Str: measureStr})

// 		log.Println(measureName, measureValue)
// 	}

// 	return infos
// }
