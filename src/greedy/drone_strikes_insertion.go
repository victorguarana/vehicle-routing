package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

const maxStrikes = 3

type droneStrikes struct {
	drone   vehicles.IDrone
	strikes int
}

func DroneStrikesInsertion(car vehicles.ICar) {
	routeIterator := car.Route().Iterator()
	droneStrikes := initDroneStrikes(car)

	for routeIterator.HasNext() {
		actualStop := routeIterator.Actual()
		if actualStop.IsDeposit() {
			landAllFlyingDrones(droneStrikes, actualStop)
			routeIterator.GoToNext()
			continue
		}

		if actualStop.IsClient() {
			if anyDroneWasStriked(droneStrikes) || anyDroneNeedToLand(droneStrikes, routeIterator.Next()) {
				landAllFlyingDrones(droneStrikes, actualStop)
				routeIterator.GoToNext()
				continue
			}

			updateDroneStrikes(droneStrikes, actualStop)
			if drone := dockedDroneThatCanSupport(droneStrikes, actualStop); drone != nil {
				drone.Move(routes.NewSubStop(actualStop.Point()))
				car.Route().RemoveMainStop(routeIterator.Index())
				routeIterator.GoToNext()
				continue
			}

			if drone := flyingDroneThatCanSupport(droneStrikes, actualStop); drone != nil {
				drone.Move(routes.NewSubStop(actualStop.Point()))
				car.Route().RemoveMainStop(routeIterator.Index())
				routeIterator.GoToNext()
				continue
			}

		}
		routeIterator.GoToNext()
	}

	landAllFlyingDrones(droneStrikes, routeIterator.Next())
}

func initDroneStrikes(car vehicles.ICar) []droneStrikes {
	drones := car.Drones()
	dStrks := make([]droneStrikes, len(drones))
	for i, drone := range drones {
		dStrks[i] = droneStrikes{drone: drone}
	}
	return dStrks
}

func landAllFlyingDrones(dStrks []droneStrikes, landingPoint routes.IMainStop) {
	for _, dStrk := range dStrks {
		if dStrk.drone.IsFlying() {
			dStrk.drone.Land(landingPoint)
			dStrk.strikes = 0
		}
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

func anyDroneNeedToLand(dStrks []droneStrikes, stop routes.IMainStop) bool {
	point := stop.Point()
	for _, dStrk := range dStrks {
		if dStrk.drone.IsFlying() && !dStrk.drone.Support(point) {
			return true
		}
	}
	return false
}

func updateDroneStrikes(dStrks []droneStrikes, stop routes.IMainStop) {
	point := stop.Point()
	for i, dStrk := range dStrks {
		if dStrk.drone.IsFlying() && dStrk.drone.Support(point) {
			dStrk.strikes = 0
		} else {
			dStrk.strikes++
		}
		dStrks[i] = dStrk
	}
}

func flyingDroneThatCanSupport(dStrks []droneStrikes, stop routes.IMainStop) vehicles.IDrone {
	point := stop.Point()
	for _, dStrk := range dStrks {
		if dStrk.drone.IsFlying() && dStrk.drone.Support(point) {
			return dStrk.drone
		}
	}
	return nil
}

func dockedDroneThatCanSupport(dStrks []droneStrikes, stop routes.IMainStop) vehicles.IDrone {
	point := stop.Point()
	for _, dStrk := range dStrks {
		if !dStrk.drone.IsFlying() && dStrk.drone.Support(point) {
			return dStrk.drone
		}
	}
	return nil
}
