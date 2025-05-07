package decoder

import (
	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

//go:generate mockgen -source=decoder.go -destination=strategy_mock_test.go -package=decoder
type strategy interface {
	DefineVehicle(carList []vehicle.ICar, chromossome *brkga.Chromossome) (vehicle.ICar, vehicle.IDrone)
}

func cloneCars(carList []vehicle.ICar) []vehicle.ICar {
	clonedCarList := make([]vehicle.ICar, len(carList))
	for i, c := range carList {
		clonedCarList[i] = c.Clone()
	}
	return clonedCarList
}

func mapItineraryByCar(carList []vehicle.ICar) map[vehicle.ICar]itinerary.Itinerary {
	itineraryByCar := make(map[vehicle.ICar]itinerary.Itinerary, len(carList))
	for _, car := range carList {
		itn := itinerary.New(car)
		itineraryByCar[car] = itn
	}
	return itineraryByCar
}

func finalizeItineraries(itineraryList []itinerary.Itinerary, gpsMap gps.Map) {
	for _, itinerary := range itineraryList {
		constructor := itinerary.Constructor()
		constructor.MoveCar(closestWarehouse(constructor.ActualCarPoint(), gpsMap))
		constructor.LandAllDrones(constructor.ActualCarStop())
	}
}

func closestWarehouse(actualPoint gps.Point, gpsMap gps.Map) gps.Point {
	return gps.ClosestPoint(actualPoint, gpsMap.Warehouses)
}

func isValidSolution(itineraryList []itinerary.Itinerary) bool {
	for _, itn := range itineraryList {
		if !itn.Validator().IsValid() {
			return false
		}
	}

	return true
}
