package agatz

import (
	"fmt"

	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

// Agatz et al. instance constants
// Fleet Constants
const agatzCarEfficiencyConst = 1.0 // Not important for this instance
const agatzDroneSpeedAlpha1Const = 1.0
const agatzDroneSpeedAlpha2Const = 0.5
const agatzDroneSpeedAlpha3Const = 0.333
const agatzDroneEfficiencyConst = 1.0 // Not important for this instance
const agatzCarSpeedConst = 1.0
const agatzRangeConst = 100000.0 // Range is not important for this instance
const agatzCarQuantityConst = 1
const agatzDroneQuantityPerCarConst = 1
const agatzDroneStorage = 1      // Fixed to 1 since drones can only do one delivery at a time
const agatzCarStorage = 100000.0 // Not important for this instance

var AgatzInstanceUniformN10 = map[string]string{
	"AGATZ - Uniform - N10 - 51": "cmd/execute/instances/agatz/uniform/uniform-51-n10",
	"AGATZ - Uniform - N10 - 52": "cmd/execute/instances/agatz/uniform/uniform-52-n10",
	"AGATZ - Uniform - N10 - 53": "cmd/execute/instances/agatz/uniform/uniform-53-n10",
	"AGATZ - Uniform - N10 - 54": "cmd/execute/instances/agatz/uniform/uniform-54-n10",
	"AGATZ - Uniform - N10 - 55": "cmd/execute/instances/agatz/uniform/uniform-55-n10",
	"AGATZ - Uniform - N10 - 56": "cmd/execute/instances/agatz/uniform/uniform-56-n10",
	"AGATZ - Uniform - N10 - 57": "cmd/execute/instances/agatz/uniform/uniform-57-n10",
	"AGATZ - Uniform - N10 - 58": "cmd/execute/instances/agatz/uniform/uniform-58-n10",
	"AGATZ - Uniform - N10 - 59": "cmd/execute/instances/agatz/uniform/uniform-59-n10",
	"AGATZ - Uniform - N10 - 60": "cmd/execute/instances/agatz/uniform/uniform-60-n10",
}
var AgatzInstanceUniformN20 = map[string]string{
	"AGATZ - Uniform - N20 - 61": "cmd/execute/instances/agatz/uniform/uniform-61-n20",
	"AGATZ - Uniform - N20 - 62": "cmd/execute/instances/agatz/uniform/uniform-62-n20",
	"AGATZ - Uniform - N20 - 63": "cmd/execute/instances/agatz/uniform/uniform-63-n20",
	"AGATZ - Uniform - N20 - 64": "cmd/execute/instances/agatz/uniform/uniform-64-n20",
	"AGATZ - Uniform - N20 - 65": "cmd/execute/instances/agatz/uniform/uniform-65-n20",
	"AGATZ - Uniform - N20 - 66": "cmd/execute/instances/agatz/uniform/uniform-66-n20",
	"AGATZ - Uniform - N20 - 67": "cmd/execute/instances/agatz/uniform/uniform-67-n20",
	"AGATZ - Uniform - N20 - 68": "cmd/execute/instances/agatz/uniform/uniform-68-n20",
	"AGATZ - Uniform - N20 - 69": "cmd/execute/instances/agatz/uniform/uniform-69-n20",
	"AGATZ - Uniform - N20 - 70": "cmd/execute/instances/agatz/uniform/uniform-70-n20",
}
var AgatzInstanceUniformN50 = map[string]string{
	"AGATZ - Uniform - N50 - 71": "cmd/execute/instances/agatz/uniform/uniform-71-n50",
	"AGATZ - Uniform - N50 - 72": "cmd/execute/instances/agatz/uniform/uniform-72-n50",
	"AGATZ - Uniform - N50 - 73": "cmd/execute/instances/agatz/uniform/uniform-73-n50",
	"AGATZ - Uniform - N50 - 74": "cmd/execute/instances/agatz/uniform/uniform-74-n50",
	"AGATZ - Uniform - N50 - 75": "cmd/execute/instances/agatz/uniform/uniform-75-n50",
	"AGATZ - Uniform - N50 - 76": "cmd/execute/instances/agatz/uniform/uniform-76-n50",
	"AGATZ - Uniform - N50 - 77": "cmd/execute/instances/agatz/uniform/uniform-77-n50",
	"AGATZ - Uniform - N50 - 78": "cmd/execute/instances/agatz/uniform/uniform-78-n50",
	"AGATZ - Uniform - N50 - 79": "cmd/execute/instances/agatz/uniform/uniform-79-n50",
	"AGATZ - Uniform - N50 - 80": "cmd/execute/instances/agatz/uniform/uniform-80-n50",
}
var AgatzInstanceUniformN75 = map[string]string{
	"AGATZ - Uniform - N75 - 81": "cmd/execute/instances/agatz/uniform/uniform-81-n75",
	"AGATZ - Uniform - N75 - 82": "cmd/execute/instances/agatz/uniform/uniform-82-n75",
	"AGATZ - Uniform - N75 - 83": "cmd/execute/instances/agatz/uniform/uniform-83-n75",
	"AGATZ - Uniform - N75 - 84": "cmd/execute/instances/agatz/uniform/uniform-84-n75",
	"AGATZ - Uniform - N75 - 85": "cmd/execute/instances/agatz/uniform/uniform-85-n75",
	"AGATZ - Uniform - N75 - 86": "cmd/execute/instances/agatz/uniform/uniform-86-n75",
	"AGATZ - Uniform - N75 - 87": "cmd/execute/instances/agatz/uniform/uniform-87-n75",
	"AGATZ - Uniform - N75 - 88": "cmd/execute/instances/agatz/uniform/uniform-88-n75",
	"AGATZ - Uniform - N75 - 89": "cmd/execute/instances/agatz/uniform/uniform-89-n75",
	"AGATZ - Uniform - N75 - 90": "cmd/execute/instances/agatz/uniform/uniform-90-n75",
}
var AgatzInstanceUniformN100 = map[string]string{
	"AGATZ - Uniform - N100 - 91":  "cmd/execute/instances/agatz/uniform/uniform-91-n100",
	"AGATZ - Uniform - N100 - 92":  "cmd/execute/instances/agatz/uniform/uniform-92-n100",
	"AGATZ - Uniform - N100 - 93":  "cmd/execute/instances/agatz/uniform/uniform-93-n100",
	"AGATZ - Uniform - N100 - 94":  "cmd/execute/instances/agatz/uniform/uniform-94-n100",
	"AGATZ - Uniform - N100 - 95":  "cmd/execute/instances/agatz/uniform/uniform-95-n100",
	"AGATZ - Uniform - N100 - 96":  "cmd/execute/instances/agatz/uniform/uniform-96-n100",
	"AGATZ - Uniform - N100 - 97":  "cmd/execute/instances/agatz/uniform/uniform-97-n100",
	"AGATZ - Uniform - N100 - 98":  "cmd/execute/instances/agatz/uniform/uniform-98-n100",
	"AGATZ - Uniform - N100 - 99":  "cmd/execute/instances/agatz/uniform/uniform-99-n100",
	"AGATZ - Uniform - N100 - 100": "cmd/execute/instances/agatz/uniform/uniform-100-n100",
}
var AgatzInstanceUniformN175 = map[string]string{
	"AGATZ - Uniform - N175 - 101": "cmd/execute/instances/agatz/uniform/uniform-101-n175",
	"AGATZ - Uniform - N175 - 102": "cmd/execute/instances/agatz/uniform/uniform-102-n175",
	"AGATZ - Uniform - N175 - 103": "cmd/execute/instances/agatz/uniform/uniform-103-n175",
	"AGATZ - Uniform - N175 - 104": "cmd/execute/instances/agatz/uniform/uniform-104-n175",
	"AGATZ - Uniform - N175 - 105": "cmd/execute/instances/agatz/uniform/uniform-105-n175",
	"AGATZ - Uniform - N175 - 106": "cmd/execute/instances/agatz/uniform/uniform-106-n175",
	"AGATZ - Uniform - N175 - 107": "cmd/execute/instances/agatz/uniform/uniform-107-n175",
	"AGATZ - Uniform - N175 - 108": "cmd/execute/instances/agatz/uniform/uniform-108-n175",
	"AGATZ - Uniform - N175 - 109": "cmd/execute/instances/agatz/uniform/uniform-109-n175",
	"AGATZ - Uniform - N175 - 110": "cmd/execute/instances/agatz/uniform/uniform-110-n175",
}
var AgatzInstanceUniformN250 = map[string]string{
	"AGATZ - Uniform - N250 - 1":   "cmd/execute/instances/agatz/uniform/uniform-1-n250",
	"AGATZ - Uniform - N250 - 2":   "cmd/execute/instances/agatz/uniform/uniform-2-n250",
	"AGATZ - Uniform - N250 - 3":   "cmd/execute/instances/agatz/uniform/uniform-3-n250",
	"AGATZ - Uniform - N250 - 4":   "cmd/execute/instances/agatz/uniform/uniform-4-n250",
	"AGATZ - Uniform - N250 - 5":   "cmd/execute/instances/agatz/uniform/uniform-5-n250",
	"AGATZ - Uniform - N250 - 6":   "cmd/execute/instances/agatz/uniform/uniform-6-n250",
	"AGATZ - Uniform - N250 - 7":   "cmd/execute/instances/agatz/uniform/uniform-7-n250",
	"AGATZ - Uniform - N250 - 8":   "cmd/execute/instances/agatz/uniform/uniform-8-n250",
	"AGATZ - Uniform - N250 - 9":   "cmd/execute/instances/agatz/uniform/uniform-9-n250",
	"AGATZ - Uniform - N250 - 10":  "cmd/execute/instances/agatz/uniform/uniform-10-n250",
	"AGATZ - Uniform - N250 - 111": "cmd/execute/instances/agatz/uniform/uniform-111-n250",
	"AGATZ - Uniform - N250 - 112": "cmd/execute/instances/agatz/uniform/uniform-112-n250",
	"AGATZ - Uniform - N250 - 113": "cmd/execute/instances/agatz/uniform/uniform-113-n250",
	"AGATZ - Uniform - N250 - 114": "cmd/execute/instances/agatz/uniform/uniform-114-n250",
	"AGATZ - Uniform - N250 - 115": "cmd/execute/instances/agatz/uniform/uniform-115-n250",
	"AGATZ - Uniform - N250 - 116": "cmd/execute/instances/agatz/uniform/uniform-116-n250",
	"AGATZ - Uniform - N250 - 117": "cmd/execute/instances/agatz/uniform/uniform-117-n250",
	"AGATZ - Uniform - N250 - 118": "cmd/execute/instances/agatz/uniform/uniform-118-n250",
	"AGATZ - Uniform - N250 - 119": "cmd/execute/instances/agatz/uniform/uniform-119-n250",
	"AGATZ - Uniform - N250 - 120": "cmd/execute/instances/agatz/uniform/uniform-120-n250",
}

type Alpha int

var Alpha1 Alpha = 1
var Alpha2 Alpha = 2
var Alpha3 Alpha = 3

var mapAlphaToDroneSpeed = map[Alpha]float64{
	Alpha1: agatzDroneSpeedAlpha1Const,
	Alpha2: agatzDroneSpeedAlpha2Const,
	Alpha3: agatzDroneSpeedAlpha3Const,
}

type agatzFleetParams struct {
	carSpeed      float64
	startingPoint gps.Point
	alpha         Alpha
}

func LoadAgatz(mapFilename string, name string, alpha Alpha) (string, gps.Map, []vehicle.ICar) {
	gpsMap := gps.LoadMap(mapFilename)
	carList := createAgatzFleet(agatzFleetParams{
		carSpeed:      agatzCarSpeedConst,
		startingPoint: gpsMap.Warehouses[0],
		alpha:         alpha,
	})
	return name, gpsMap, carList
}

func createAgatzFleet(fp agatzFleetParams) []vehicle.ICar {
	cars := make([]vehicle.ICar, agatzCarQuantityConst)
	for i := 0; i < agatzCarQuantityConst; i++ {
		car := vehicle.NewCarLimited(vehicle.CarParams{
			Efficiency:    agatzCarEfficiencyConst,
			Speed:         fp.carSpeed,
			Storage:       agatzCarStorage,
			Range:         agatzRangeConst,
			Name:          "car" + fmt.Sprint(i),
			StartingPoint: fp.startingPoint,
		})
		for j := 0; j < agatzDroneQuantityPerCarConst; j++ {
			car.NewDroneWithParams(
				vehicle.DroneParams{
					Efficiency:    agatzDroneEfficiencyConst,
					Speed:         mapAlphaToDroneSpeed[fp.alpha],
					Storage:       agatzDroneStorage,
					Range:         agatzRangeConst,
					Name:          "drone" + fmt.Sprint(j) + "_of_car" + fmt.Sprint(i),
					StartingPoint: fp.startingPoint,
				},
			)
		}
		cars[i] = car
	}
	return cars
}
