package itinerary

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/route"
)

//go:generate mockgen -source=validator.go -destination=mock/validatormock.go
type Validator interface {
	IsValid() bool
}

type validator struct {
	*info
}

type requirements struct {
	requiredStorage float64
	requiredRange   float64
}

func (v *validator) IsValid() bool {
	return v.carCanSupportRoute() && v.dronesCanSupportFlights()
}

func (v *validator) carCanSupportRoute() bool {
	requiredStorage := 0.0
	requiredRange := 0.0

	mainStopList := v.route.MainStopList()
	prevStop := mainStopList[0]
	for _, actualStop := range mainStopList {
		requiredRange += gps.ManhattanDistanceBetweenPoints(prevStop.Point(), actualStop.Point())
		if requiredRange > v.car.Range() {
			return false
		}

		if actualStop.IsWarehouse() {
			requiredRange = 0.0
			requiredStorage = 0.0
		}

		if actualStop.IsClient() {
			requiredStorage += v.calcMainStopRequiredStorage(actualStop)
			if requiredStorage > v.car.Storage() {
				return false
			}
		}
		prevStop = actualStop
	}

	return true
}

func (v *validator) calcMainStopRequiredStorage(customer route.IMainStop) float64 {
	requiredStorage := customer.Point().PackageSize

	for _, flight := range customer.StartingSubRoutes() {
		flightRequirements := v.calcFlightRequirements(flight)
		requiredStorage += flightRequirements.requiredStorage
	}

	return requiredStorage
}

func (*validator) calcFlightRequirements(flight route.ISubRoute) requirements {
	requiredRange := 0.0
	requiredStorage := 0.0

	prevPoint := flight.StartingStop().Point()
	subStops := flight.SubStopList()

	for _, actualStop := range subStops {
		actualPoint := actualStop.Point()
		requiredRange += gps.EuclideanDistanceBetweenPoints(prevPoint, actualPoint)
		requiredStorage += actualPoint.PackageSize
		prevPoint = actualPoint
	}

	actualPoint := flight.ReturningStop().Point()
	requiredRange += gps.EuclideanDistanceBetweenPoints(prevPoint, actualPoint)

	return requirements{
		requiredStorage: requiredStorage,
		requiredRange:   requiredRange,
	}
}

func (v *validator) dronesCanSupportFlights() bool {
	for _, subItinerary := range v.completedSubItineraryList {
		if !v.isValidSubItinerary(subItinerary) {
			return false
		}
	}
	return true
}

func (v *validator) isValidSubItinerary(subItinerary SubItinerary) bool {
	flightRequirements := v.calcFlightRequirements(subItinerary.Flight)

	if flightRequirements.requiredRange > subItinerary.Drone.Range() {
		return false
	}

	if flightRequirements.requiredStorage > subItinerary.Drone.Storage() {
		return false
	}

	return true
}
