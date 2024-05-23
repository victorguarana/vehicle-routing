package greedy

import (
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/route"
)

const maxStrikes = 3

type droneStrikes struct {
	droneNumber itinerary.DroneNumber
	strikes     int
}

func DroneStrikesInsertion(itinerary itinerary.Itinerary) {
	routeIterator := itinerary.RouteIterator()
	droneStrikes := initDroneStrikes(itinerary)

	for routeIterator.HasNext() {
		actualStop := routeIterator.Actual()
		nextStop := routeIterator.Next()
		if actualStop.IsWarehouse() {
			itinerary.LandAllDrones(actualStop)
			resetDroneStrikes(droneStrikes)
			routeIterator.GoToNext()
			continue
		}

		if actualStop.IsClient() {
			if anyDroneWasStriked(droneStrikes) || anyDroneNeedToLand(itinerary, droneStrikes, nextStop) {
				itinerary.LandAllDrones(actualStop)
				resetDroneStrikes(droneStrikes)
				routeIterator.GoToNext()
				continue
			}

			updateDroneStrikes(itinerary, droneStrikes, actualStop, nextStop)
			if droneNumber, exists := dockedDroneThatCanSupport(itinerary, droneStrikes, actualStop, nextStop); exists {
				itinerary.StartDroneFlight(droneNumber, routeIterator.Previous())
				itinerary.MoveDrone(droneNumber, actualStop.Point())
				itinerary.RemoveMainStopFromRoute(routeIterator.Index())
				routeIterator.RemoveActualIndex()
				continue
			}

			if droneNumber, exists := flyingDroneThatCanSupport(itinerary, droneStrikes, actualStop, nextStop); exists {
				itinerary.MoveDrone(droneNumber, actualStop.Point())
				itinerary.RemoveMainStopFromRoute(routeIterator.Index())
				routeIterator.RemoveActualIndex()
				continue
			}
		}
		routeIterator.GoToNext()
	}
	itinerary.LandAllDrones(routeIterator.Actual())
}

func initDroneStrikes(itinerary itinerary.Itinerary) []droneStrikes {
	droneNumbers := itinerary.DroneNumbers()
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

func anyDroneNeedToLand(itinerary itinerary.Itinerary, dStrks []droneStrikes, next route.IMainStop) bool {
	nextPoint := next.Point()
	for _, dStrk := range dStrks {
		if itinerary.DroneIsFlying(dStrk.droneNumber) && !itinerary.DroneCanReach(dStrk.droneNumber, nextPoint) {
			return true
		}
	}
	return false
}

func updateDroneStrikes(itinerary itinerary.Itinerary, dStrks []droneStrikes, actual route.IMainStop, next route.IMainStop) {
	actualPoint := actual.Point()
	nextPoint := next.Point()
	for i, dStrk := range dStrks {
		if itinerary.DroneIsFlying(dStrk.droneNumber) {
			if itinerary.DroneSupport(dStrk.droneNumber, actualPoint, nextPoint) {
				dStrk.strikes = 0
			} else {
				dStrk.strikes++
			}
			dStrks[i] = dStrk
		}
	}
}

func flyingDroneThatCanSupport(itn itinerary.Itinerary, dStrks []droneStrikes, actual route.IMainStop, next route.IMainStop) (itinerary.DroneNumber, bool) {
	actualPoint := actual.Point()
	nextPoint := next.Point()
	nextPoint.PackageSize = 0
	for _, dStrk := range dStrks {
		if itn.DroneIsFlying(dStrk.droneNumber) && itn.DroneSupport(dStrk.droneNumber, actualPoint, nextPoint) {
			return dStrk.droneNumber, true
		}
	}
	return 0, false
}

func dockedDroneThatCanSupport(itn itinerary.Itinerary, dStrks []droneStrikes, actual route.IMainStop, next route.IMainStop) (itinerary.DroneNumber, bool) {
	actualPoint := actual.Point()
	nextPoint := next.Point()
	nextPoint.PackageSize = 0
	for _, dStrk := range dStrks {
		if !itn.DroneIsFlying(dStrk.droneNumber) && itn.DroneSupport(dStrk.droneNumber, actualPoint, nextPoint) {
			return dStrk.droneNumber, true
		}
	}
	return 0, false
}
