package greedy

import (
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/route"
)

const maxStrikes = 3

type droneStrikes struct {
	droneNumber itinerary.DroneNumber
	strikes     int
}

func DroneStrikesInsertion(constructor itinerary.Constructor, modifier itinerary.Modifier) {
	routeIterator := constructor.RouteIterator()
	droneStrikes := initDroneStrikes(constructor)

	for routeIterator.HasNext() {
		actualStop := routeIterator.Actual()
		nextStop := routeIterator.Next()
		if actualStop.IsWarehouse() {
			constructor.LandAllDrones(actualStop)
			resetDroneStrikes(droneStrikes)
			routeIterator.GoToNext()
			continue
		}

		if actualStop.IsClient() {
			if anyDroneWasStriked(droneStrikes) || anyDroneNeedToLand(constructor, droneStrikes, nextStop) {
				constructor.LandAllDrones(actualStop)
				resetDroneStrikes(droneStrikes)
				routeIterator.GoToNext()
				continue
			}

			updateDroneStrikes(constructor, droneStrikes, actualStop, nextStop)
			if droneNumber, exists := dockedDroneThatCanSupport(constructor, droneStrikes, actualStop, nextStop); exists {
				constructor.StartDroneFlight(droneNumber, routeIterator.Previous())
				constructor.MoveDrone(droneNumber, actualStop.Point())
				modifier.RemoveMainStopFromRoute(routeIterator.Index())
				routeIterator.RemoveActualIndex()
				continue
			}

			if droneNumber, exists := flyingDroneThatCanSupport(constructor, droneStrikes, actualStop, nextStop); exists {
				constructor.MoveDrone(droneNumber, actualStop.Point())
				modifier.RemoveMainStopFromRoute(routeIterator.Index())
				routeIterator.RemoveActualIndex()
				continue
			}
		}
		routeIterator.GoToNext()
	}
	constructor.LandAllDrones(routeIterator.Actual())
}

func initDroneStrikes(constructor itinerary.Constructor) []droneStrikes {
	droneNumbers := constructor.DroneNumbers()
	dStrks := make([]droneStrikes, len(droneNumbers))
	for i, droneNumber := range droneNumbers {
		dStrks[i] = droneStrikes{droneNumber: droneNumber}
	}
	return dStrks
}

func resetDroneStrikes(dStrks []droneStrikes) {
	for i, dStrk := range dStrks {
		dStrk.strikes = 0
		dStrks[i] = dStrk
	}
}

func anyDroneWasStriked(dStrks []droneStrikes) bool {
	for _, dStrk := range dStrks {
		if dStrk.strikes >= maxStrikes {
			return true
		}
	}
	return false
}

func anyDroneNeedToLand(constructor itinerary.Constructor, dStrks []droneStrikes, next route.IMainStop) bool {
	nextPoint := next.Point()
	for _, dStrk := range dStrks {
		if constructor.DroneIsFlying(dStrk.droneNumber) && !constructor.DroneCanReach(dStrk.droneNumber, nextPoint) {
			return true
		}
	}
	return false
}

func updateDroneStrikes(constructor itinerary.Constructor, dStrks []droneStrikes, actual route.IMainStop, next route.IMainStop) {
	actualPoint := actual.Point()
	nextPoint := next.Point()
	for i, dStrk := range dStrks {
		if constructor.DroneIsFlying(dStrk.droneNumber) {
			if constructor.DroneSupport(dStrk.droneNumber, actualPoint, nextPoint) {
				dStrk.strikes = 0
			} else {
				dStrk.strikes++
			}
			dStrks[i] = dStrk
		}
	}
}

func flyingDroneThatCanSupport(constructor itinerary.Constructor, dStrks []droneStrikes, actual route.IMainStop, next route.IMainStop) (itinerary.DroneNumber, bool) {
	actualPoint := actual.Point()
	nextPoint := next.Point()
	nextPoint.PackageSize = 0
	for _, dStrk := range dStrks {
		if constructor.DroneIsFlying(dStrk.droneNumber) && constructor.DroneSupport(dStrk.droneNumber, actualPoint, nextPoint) {
			return dStrk.droneNumber, true
		}
	}
	return 0, false
}

func dockedDroneThatCanSupport(constructor itinerary.Constructor, dStrks []droneStrikes, actual route.IMainStop, next route.IMainStop) (itinerary.DroneNumber, bool) {
	actualPoint := actual.Point()
	nextPoint := next.Point()
	nextPoint.PackageSize = 0
	for _, dStrk := range dStrks {
		if !constructor.DroneIsFlying(dStrk.droneNumber) && constructor.DroneSupport(dStrk.droneNumber, actualPoint, nextPoint) {
			return dStrk.droneNumber, true
		}
	}
	return 0, false
}
