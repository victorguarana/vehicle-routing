package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

// TODO: Check for errors

// Insert 2 different drones in the same route
// Each drone will deliver a client and return immediately to the car
func DroneSimpleInsertion(route routes.IRoute) {
	car := route.Car()
	drones := car.Drones()

	for actualCarStopIndex := 1; actualCarStopIndex < route.Len(); actualCarStopIndex++ {
		previousCarStop, _ := route.AtIndex(actualCarStopIndex - 1)
		actualCarStop, _ := route.AtIndex(actualCarStopIndex)
		nextCarStop, _ := route.AtIndex(actualCarStopIndex + 1)

		// Case #1: Actual point is a Client and next is a Deposit -> Deliver actual point with the first available drone
		if actualCarStop.IsClient() && nextCarStop.IsDeposit() {
			if drones[0].Support(actualCarStop.Point()) {
				createFlight(drones[0], previousCarStop, nextCarStop, actualCarStop.Point())
				route.RemoveCarStop(actualCarStopIndex)
			}

			continue
		}

		// Case #2: Actual and next points are Clients
		if actualCarStop.IsClient() && nextCarStop.IsClient() {
			drone0SupportsActual := drones[0].Support(actualCarStop.Point())
			drone1SupportsNext := drones[1].Support(nextCarStop.Point())
			nextNextCarStop, _ := route.AtIndex(actualCarStopIndex + 2)

			switch {
			// Case #2.1: Both drones support both clients -> Deliver each client with an each drone
			case drone0SupportsActual && drone1SupportsNext:
				createFlight(drones[0], previousCarStop, nextNextCarStop, actualCarStop.Point())
				createFlight(drones[1], previousCarStop, nextNextCarStop, nextCarStop.Point())

				route.RemoveCarStop(actualCarStopIndex)
				route.RemoveCarStop(actualCarStopIndex)

				// Increment the index because the next point is now a landing point

			// Case #2.2: Just first drone supports actual -> Just deliver actual client with the first drone
			case drone0SupportsActual && !drone1SupportsNext:
				createFlight(drones[0], previousCarStop, nextCarStop, actualCarStop.Point())
				route.RemoveCarStop(actualCarStopIndex)

			// Case #2.3: Just second drone supports next client -> Just deliver next client with the second drone
			case !drone0SupportsActual && drone1SupportsNext:
				createFlight(drones[1], actualCarStop, nextNextCarStop, nextCarStop.Point())
				route.RemoveCarStop(actualCarStopIndex + 1)

				// Increment the index because we skipped the actual point
				// actualCarStopIndex++
			}

			continue
		}
	}
}

func createFlight(drone vehicles.IDrone, takeoffCarStop, landingCarStop routes.ICarStop, point *gps.Point) error {
	actualFlight, err := routes.NewFlight(drone, takeoffCarStop, landingCarStop)
	if err != nil {
		return err
	}

	err = actualFlight.Append(point)
	if err != nil {
		return err
	}

	err = drone.Move(point)
	if err != nil {
		return err
	}

	err = drone.Land(landingCarStop.Point())
	if err != nil {
		return err
	}

	return nil
}
