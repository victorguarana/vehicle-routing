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
const agatzDroneSpeedAlpha2Const = 2.0
const agatzDroneSpeedAlpha3Const = 3.0
const agatzDroneEfficiencyConst = 1.0 // Not important for this instance
const agatzCarSpeedConst = 1.0
const agatzRangeConst = 100000.0 // Range is not important for this instance
const agatzCarQuantityConst = 1
const agatzDroneQuantityPerCarConst = 1
const agatzDroneStorage = 5      // Fixed to 1 since drones can only do one delivery at a time
const agatzCarStorage = 100000.0 // Not important for this instance

var AgatzInstanceUniformN10 = map[string]string{
	"AGATZ - Uniform - N010 - 51": "cmd/execute/instances/agatz/uniform/uniform-51-n10",
	"AGATZ - Uniform - N010 - 52": "cmd/execute/instances/agatz/uniform/uniform-52-n10",
	"AGATZ - Uniform - N010 - 53": "cmd/execute/instances/agatz/uniform/uniform-53-n10",
	"AGATZ - Uniform - N010 - 54": "cmd/execute/instances/agatz/uniform/uniform-54-n10",
	"AGATZ - Uniform - N010 - 55": "cmd/execute/instances/agatz/uniform/uniform-55-n10",
	"AGATZ - Uniform - N010 - 56": "cmd/execute/instances/agatz/uniform/uniform-56-n10",
	"AGATZ - Uniform - N010 - 57": "cmd/execute/instances/agatz/uniform/uniform-57-n10",
	"AGATZ - Uniform - N010 - 58": "cmd/execute/instances/agatz/uniform/uniform-58-n10",
	"AGATZ - Uniform - N010 - 59": "cmd/execute/instances/agatz/uniform/uniform-59-n10",
	"AGATZ - Uniform - N010 - 60": "cmd/execute/instances/agatz/uniform/uniform-60-n10",
}
var AgatzInstanceUniformN20 = map[string]string{
	"AGATZ - Uniform - N020 - 61": "cmd/execute/instances/agatz/uniform/uniform-61-n20",
	"AGATZ - Uniform - N020 - 62": "cmd/execute/instances/agatz/uniform/uniform-62-n20",
	"AGATZ - Uniform - N020 - 63": "cmd/execute/instances/agatz/uniform/uniform-63-n20",
	"AGATZ - Uniform - N020 - 64": "cmd/execute/instances/agatz/uniform/uniform-64-n20",
	"AGATZ - Uniform - N020 - 65": "cmd/execute/instances/agatz/uniform/uniform-65-n20",
	"AGATZ - Uniform - N020 - 66": "cmd/execute/instances/agatz/uniform/uniform-66-n20",
	"AGATZ - Uniform - N020 - 67": "cmd/execute/instances/agatz/uniform/uniform-67-n20",
	"AGATZ - Uniform - N020 - 68": "cmd/execute/instances/agatz/uniform/uniform-68-n20",
	"AGATZ - Uniform - N020 - 69": "cmd/execute/instances/agatz/uniform/uniform-69-n20",
	"AGATZ - Uniform - N020 - 70": "cmd/execute/instances/agatz/uniform/uniform-70-n20",
}
var AgatzInstanceUniformN50 = map[string]string{
	"AGATZ - Uniform - N050 - 71": "cmd/execute/instances/agatz/uniform/uniform-71-n50",
	"AGATZ - Uniform - N050 - 72": "cmd/execute/instances/agatz/uniform/uniform-72-n50",
	"AGATZ - Uniform - N050 - 73": "cmd/execute/instances/agatz/uniform/uniform-73-n50",
	"AGATZ - Uniform - N050 - 74": "cmd/execute/instances/agatz/uniform/uniform-74-n50",
	"AGATZ - Uniform - N050 - 75": "cmd/execute/instances/agatz/uniform/uniform-75-n50",
	"AGATZ - Uniform - N050 - 76": "cmd/execute/instances/agatz/uniform/uniform-76-n50",
	"AGATZ - Uniform - N050 - 77": "cmd/execute/instances/agatz/uniform/uniform-77-n50",
	"AGATZ - Uniform - N050 - 78": "cmd/execute/instances/agatz/uniform/uniform-78-n50",
	"AGATZ - Uniform - N050 - 79": "cmd/execute/instances/agatz/uniform/uniform-79-n50",
	"AGATZ - Uniform - N050 - 80": "cmd/execute/instances/agatz/uniform/uniform-80-n50",
}
var AgatzInstanceUniformN75 = map[string]string{
	"AGATZ - Uniform - N075 - 81": "cmd/execute/instances/agatz/uniform/uniform-81-n75",
	"AGATZ - Uniform - N075 - 82": "cmd/execute/instances/agatz/uniform/uniform-82-n75",
	"AGATZ - Uniform - N075 - 83": "cmd/execute/instances/agatz/uniform/uniform-83-n75",
	"AGATZ - Uniform - N075 - 84": "cmd/execute/instances/agatz/uniform/uniform-84-n75",
	"AGATZ - Uniform - N075 - 85": "cmd/execute/instances/agatz/uniform/uniform-85-n75",
	"AGATZ - Uniform - N075 - 86": "cmd/execute/instances/agatz/uniform/uniform-86-n75",
	"AGATZ - Uniform - N075 - 87": "cmd/execute/instances/agatz/uniform/uniform-87-n75",
	"AGATZ - Uniform - N075 - 88": "cmd/execute/instances/agatz/uniform/uniform-88-n75",
	"AGATZ - Uniform - N075 - 89": "cmd/execute/instances/agatz/uniform/uniform-89-n75",
	"AGATZ - Uniform - N075 - 90": "cmd/execute/instances/agatz/uniform/uniform-90-n75",
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

var AgatzInstanceSingleCenterN10 = map[string]string{
	"AGATZ - SingleCenter - N010 - 51": "cmd/execute/instances/agatz/singlecenter/singlecenter-51-n10",
	"AGATZ - SingleCenter - N010 - 52": "cmd/execute/instances/agatz/singlecenter/singlecenter-52-n10",
	"AGATZ - SingleCenter - N010 - 53": "cmd/execute/instances/agatz/singlecenter/singlecenter-53-n10",
	"AGATZ - SingleCenter - N010 - 54": "cmd/execute/instances/agatz/singlecenter/singlecenter-54-n10",
	"AGATZ - SingleCenter - N010 - 55": "cmd/execute/instances/agatz/singlecenter/singlecenter-55-n10",
	"AGATZ - SingleCenter - N010 - 56": "cmd/execute/instances/agatz/singlecenter/singlecenter-56-n10",
	"AGATZ - SingleCenter - N010 - 57": "cmd/execute/instances/agatz/singlecenter/singlecenter-57-n10",
	"AGATZ - SingleCenter - N010 - 58": "cmd/execute/instances/agatz/singlecenter/singlecenter-58-n10",
	"AGATZ - SingleCenter - N010 - 59": "cmd/execute/instances/agatz/singlecenter/singlecenter-59-n10",
	"AGATZ - SingleCenter - N010 - 60": "cmd/execute/instances/agatz/singlecenter/singlecenter-60-n10",
}
var AgatzInstanceSingleCenterN20 = map[string]string{
	"AGATZ - SingleCenter - N020 - 61": "cmd/execute/instances/agatz/singlecenter/singlecenter-61-n20",
	"AGATZ - SingleCenter - N020 - 62": "cmd/execute/instances/agatz/singlecenter/singlecenter-62-n20",
	"AGATZ - SingleCenter - N020 - 63": "cmd/execute/instances/agatz/singlecenter/singlecenter-63-n20",
	"AGATZ - SingleCenter - N020 - 64": "cmd/execute/instances/agatz/singlecenter/singlecenter-64-n20",
	"AGATZ - SingleCenter - N020 - 65": "cmd/execute/instances/agatz/singlecenter/singlecenter-65-n20",
	"AGATZ - SingleCenter - N020 - 66": "cmd/execute/instances/agatz/singlecenter/singlecenter-66-n20",
	"AGATZ - SingleCenter - N020 - 67": "cmd/execute/instances/agatz/singlecenter/singlecenter-67-n20",
	"AGATZ - SingleCenter - N020 - 68": "cmd/execute/instances/agatz/singlecenter/singlecenter-68-n20",
	"AGATZ - SingleCenter - N020 - 69": "cmd/execute/instances/agatz/singlecenter/singlecenter-69-n20",
	"AGATZ - SingleCenter - N020 - 70": "cmd/execute/instances/agatz/singlecenter/singlecenter-70-n20",
}
var AgatzInstanceSingleCenterN50 = map[string]string{
	"AGATZ - SingleCenter - N050 - 71": "cmd/execute/instances/agatz/singlecenter/singlecenter-71-n50",
	"AGATZ - SingleCenter - N050 - 72": "cmd/execute/instances/agatz/singlecenter/singlecenter-72-n50",
	"AGATZ - SingleCenter - N050 - 73": "cmd/execute/instances/agatz/singlecenter/singlecenter-73-n50",
	"AGATZ - SingleCenter - N050 - 74": "cmd/execute/instances/agatz/singlecenter/singlecenter-74-n50",
	"AGATZ - SingleCenter - N050 - 75": "cmd/execute/instances/agatz/singlecenter/singlecenter-75-n50",
	"AGATZ - SingleCenter - N050 - 76": "cmd/execute/instances/agatz/singlecenter/singlecenter-76-n50",
	"AGATZ - SingleCenter - N050 - 77": "cmd/execute/instances/agatz/singlecenter/singlecenter-77-n50",
	"AGATZ - SingleCenter - N050 - 78": "cmd/execute/instances/agatz/singlecenter/singlecenter-78-n50",
	"AGATZ - SingleCenter - N050 - 79": "cmd/execute/instances/agatz/singlecenter/singlecenter-79-n50",
	"AGATZ - SingleCenter - N050 - 80": "cmd/execute/instances/agatz/singlecenter/singlecenter-80-n50",
}
var AgatzInstanceSingleCenterN75 = map[string]string{
	"AGATZ - SingleCenter - N075 - 81": "cmd/execute/instances/agatz/singlecenter/singlecenter-81-n75",
	"AGATZ - SingleCenter - N075 - 82": "cmd/execute/instances/agatz/singlecenter/singlecenter-82-n75",
	"AGATZ - SingleCenter - N075 - 83": "cmd/execute/instances/agatz/singlecenter/singlecenter-83-n75",
	"AGATZ - SingleCenter - N075 - 84": "cmd/execute/instances/agatz/singlecenter/singlecenter-84-n75",
	"AGATZ - SingleCenter - N075 - 85": "cmd/execute/instances/agatz/singlecenter/singlecenter-85-n75",
	"AGATZ - SingleCenter - N075 - 86": "cmd/execute/instances/agatz/singlecenter/singlecenter-86-n75",
	"AGATZ - SingleCenter - N075 - 87": "cmd/execute/instances/agatz/singlecenter/singlecenter-87-n75",
	"AGATZ - SingleCenter - N075 - 88": "cmd/execute/instances/agatz/singlecenter/singlecenter-88-n75",
	"AGATZ - SingleCenter - N075 - 89": "cmd/execute/instances/agatz/singlecenter/singlecenter-89-n75",
	"AGATZ - SingleCenter - N075 - 90": "cmd/execute/instances/agatz/singlecenter/singlecenter-90-n75",
}
var AgatzInstanceSingleCenterN100 = map[string]string{
	"AGATZ - SingleCenter - N100 - 91":  "cmd/execute/instances/agatz/singlecenter/singlecenter-91-n100",
	"AGATZ - SingleCenter - N100 - 92":  "cmd/execute/instances/agatz/singlecenter/singlecenter-92-n100",
	"AGATZ - SingleCenter - N100 - 93":  "cmd/execute/instances/agatz/singlecenter/singlecenter-93-n100",
	"AGATZ - SingleCenter - N100 - 94":  "cmd/execute/instances/agatz/singlecenter/singlecenter-94-n100",
	"AGATZ - SingleCenter - N100 - 95":  "cmd/execute/instances/agatz/singlecenter/singlecenter-95-n100",
	"AGATZ - SingleCenter - N100 - 96":  "cmd/execute/instances/agatz/singlecenter/singlecenter-96-n100",
	"AGATZ - SingleCenter - N100 - 97":  "cmd/execute/instances/agatz/singlecenter/singlecenter-97-n100",
	"AGATZ - SingleCenter - N100 - 98":  "cmd/execute/instances/agatz/singlecenter/singlecenter-98-n100",
	"AGATZ - SingleCenter - N100 - 99":  "cmd/execute/instances/agatz/singlecenter/singlecenter-99-n100",
	"AGATZ - SingleCenter - N100 - 100": "cmd/execute/instances/agatz/singlecenter/singlecenter-100-n100",
}
var AgatzInstanceSingleCenterN175 = map[string]string{
	"AGATZ - SingleCenter - N175 - 101": "cmd/execute/instances/agatz/singlecenter/singlecenter-101-n175",
	"AGATZ - SingleCenter - N175 - 102": "cmd/execute/instances/agatz/singlecenter/singlecenter-102-n175",
	"AGATZ - SingleCenter - N175 - 103": "cmd/execute/instances/agatz/singlecenter/singlecenter-103-n175",
	"AGATZ - SingleCenter - N175 - 104": "cmd/execute/instances/agatz/singlecenter/singlecenter-104-n175",
	"AGATZ - SingleCenter - N175 - 105": "cmd/execute/instances/agatz/singlecenter/singlecenter-105-n175",
	"AGATZ - SingleCenter - N175 - 106": "cmd/execute/instances/agatz/singlecenter/singlecenter-106-n175",
	"AGATZ - SingleCenter - N175 - 107": "cmd/execute/instances/agatz/singlecenter/singlecenter-107-n175",
	"AGATZ - SingleCenter - N175 - 108": "cmd/execute/instances/agatz/singlecenter/singlecenter-108-n175",
	"AGATZ - SingleCenter - N175 - 109": "cmd/execute/instances/agatz/singlecenter/singlecenter-109-n175",
	"AGATZ - SingleCenter - N175 - 110": "cmd/execute/instances/agatz/singlecenter/singlecenter-110-n175",
}
var AgatzInstanceSingleCenterN250 = map[string]string{
	"AGATZ - SingleCenter - N250 - 1":   "cmd/execute/instances/agatz/singlecenter/singlecenter-1-n250",
	"AGATZ - SingleCenter - N250 - 2":   "cmd/execute/instances/agatz/singlecenter/singlecenter-2-n250",
	"AGATZ - SingleCenter - N250 - 3":   "cmd/execute/instances/agatz/singlecenter/singlecenter-3-n250",
	"AGATZ - SingleCenter - N250 - 4":   "cmd/execute/instances/agatz/singlecenter/singlecenter-4-n250",
	"AGATZ - SingleCenter - N250 - 5":   "cmd/execute/instances/agatz/singlecenter/singlecenter-5-n250",
	"AGATZ - SingleCenter - N250 - 6":   "cmd/execute/instances/agatz/singlecenter/singlecenter-6-n250",
	"AGATZ - SingleCenter - N250 - 7":   "cmd/execute/instances/agatz/singlecenter/singlecenter-7-n250",
	"AGATZ - SingleCenter - N250 - 8":   "cmd/execute/instances/agatz/singlecenter/singlecenter-8-n250",
	"AGATZ - SingleCenter - N250 - 9":   "cmd/execute/instances/agatz/singlecenter/singlecenter-9-n250",
	"AGATZ - SingleCenter - N250 - 10":  "cmd/execute/instances/agatz/singlecenter/singlecenter-10-n250",
	"AGATZ - SingleCenter - N250 - 111": "cmd/execute/instances/agatz/singlecenter/singlecenter-111-n250",
	"AGATZ - SingleCenter - N250 - 112": "cmd/execute/instances/agatz/singlecenter/singlecenter-112-n250",
	"AGATZ - SingleCenter - N250 - 113": "cmd/execute/instances/agatz/singlecenter/singlecenter-113-n250",
	"AGATZ - SingleCenter - N250 - 114": "cmd/execute/instances/agatz/singlecenter/singlecenter-114-n250",
	"AGATZ - SingleCenter - N250 - 115": "cmd/execute/instances/agatz/singlecenter/singlecenter-115-n250",
	"AGATZ - SingleCenter - N250 - 116": "cmd/execute/instances/agatz/singlecenter/singlecenter-116-n250",
	"AGATZ - SingleCenter - N250 - 117": "cmd/execute/instances/agatz/singlecenter/singlecenter-117-n250",
	"AGATZ - SingleCenter - N250 - 118": "cmd/execute/instances/agatz/singlecenter/singlecenter-118-n250",
	"AGATZ - SingleCenter - N250 - 119": "cmd/execute/instances/agatz/singlecenter/singlecenter-119-n250",
	"AGATZ - SingleCenter - N250 - 120": "cmd/execute/instances/agatz/singlecenter/singlecenter-120-n250",
}

var AgatzInstanceDoubleCenterN10 = map[string]string{
	"AGATZ - DoubleCenter - N010 - 51": "cmd/execute/instances/agatz/doublecenter/doublecenter-51-n10",
	"AGATZ - DoubleCenter - N010 - 52": "cmd/execute/instances/agatz/doublecenter/doublecenter-52-n10",
	"AGATZ - DoubleCenter - N010 - 53": "cmd/execute/instances/agatz/doublecenter/doublecenter-53-n10",
	"AGATZ - DoubleCenter - N010 - 54": "cmd/execute/instances/agatz/doublecenter/doublecenter-54-n10",
	"AGATZ - DoubleCenter - N010 - 55": "cmd/execute/instances/agatz/doublecenter/doublecenter-55-n10",
	"AGATZ - DoubleCenter - N010 - 56": "cmd/execute/instances/agatz/doublecenter/doublecenter-56-n10",
	"AGATZ - DoubleCenter - N010 - 57": "cmd/execute/instances/agatz/doublecenter/doublecenter-57-n10",
	"AGATZ - DoubleCenter - N010 - 58": "cmd/execute/instances/agatz/doublecenter/doublecenter-58-n10",
	"AGATZ - DoubleCenter - N010 - 59": "cmd/execute/instances/agatz/doublecenter/doublecenter-59-n10",
	"AGATZ - DoubleCenter - N010 - 60": "cmd/execute/instances/agatz/doublecenter/doublecenter-60-n10",
}
var AgatzInstanceDoubleCenterN20 = map[string]string{
	"AGATZ - DoubleCenter - N020 - 61": "cmd/execute/instances/agatz/doublecenter/doublecenter-61-n20",
	"AGATZ - DoubleCenter - N020 - 62": "cmd/execute/instances/agatz/doublecenter/doublecenter-62-n20",
	"AGATZ - DoubleCenter - N020 - 63": "cmd/execute/instances/agatz/doublecenter/doublecenter-63-n20",
	"AGATZ - DoubleCenter - N020 - 64": "cmd/execute/instances/agatz/doublecenter/doublecenter-64-n20",
	"AGATZ - DoubleCenter - N020 - 65": "cmd/execute/instances/agatz/doublecenter/doublecenter-65-n20",
	"AGATZ - DoubleCenter - N020 - 66": "cmd/execute/instances/agatz/doublecenter/doublecenter-66-n20",
	"AGATZ - DoubleCenter - N020 - 67": "cmd/execute/instances/agatz/doublecenter/doublecenter-67-n20",
	"AGATZ - DoubleCenter - N020 - 68": "cmd/execute/instances/agatz/doublecenter/doublecenter-68-n20",
	"AGATZ - DoubleCenter - N020 - 69": "cmd/execute/instances/agatz/doublecenter/doublecenter-69-n20",
	"AGATZ - DoubleCenter - N020 - 70": "cmd/execute/instances/agatz/doublecenter/doublecenter-70-n20",
}
var AgatzInstanceDoubleCenterN50 = map[string]string{
	"AGATZ - DoubleCenter - N050 - 71": "cmd/execute/instances/agatz/doublecenter/doublecenter-71-n50",
	"AGATZ - DoubleCenter - N050 - 72": "cmd/execute/instances/agatz/doublecenter/doublecenter-72-n50",
	"AGATZ - DoubleCenter - N050 - 73": "cmd/execute/instances/agatz/doublecenter/doublecenter-73-n50",
	"AGATZ - DoubleCenter - N050 - 74": "cmd/execute/instances/agatz/doublecenter/doublecenter-74-n50",
	"AGATZ - DoubleCenter - N050 - 75": "cmd/execute/instances/agatz/doublecenter/doublecenter-75-n50",
	"AGATZ - DoubleCenter - N050 - 76": "cmd/execute/instances/agatz/doublecenter/doublecenter-76-n50",
	"AGATZ - DoubleCenter - N050 - 77": "cmd/execute/instances/agatz/doublecenter/doublecenter-77-n50",
	"AGATZ - DoubleCenter - N050 - 78": "cmd/execute/instances/agatz/doublecenter/doublecenter-78-n50",
	"AGATZ - DoubleCenter - N050 - 79": "cmd/execute/instances/agatz/doublecenter/doublecenter-79-n50",
	"AGATZ - DoubleCenter - N050 - 80": "cmd/execute/instances/agatz/doublecenter/doublecenter-80-n50",
}
var AgatzInstanceDoubleCenterN75 = map[string]string{
	"AGATZ - DoubleCenter - N075 - 81": "cmd/execute/instances/agatz/doublecenter/doublecenter-81-n75",
	"AGATZ - DoubleCenter - N075 - 82": "cmd/execute/instances/agatz/doublecenter/doublecenter-82-n75",
	"AGATZ - DoubleCenter - N075 - 83": "cmd/execute/instances/agatz/doublecenter/doublecenter-83-n75",
	"AGATZ - DoubleCenter - N075 - 84": "cmd/execute/instances/agatz/doublecenter/doublecenter-84-n75",
	"AGATZ - DoubleCenter - N075 - 85": "cmd/execute/instances/agatz/doublecenter/doublecenter-85-n75",
	"AGATZ - DoubleCenter - N075 - 86": "cmd/execute/instances/agatz/doublecenter/doublecenter-86-n75",
	"AGATZ - DoubleCenter - N075 - 87": "cmd/execute/instances/agatz/doublecenter/doublecenter-87-n75",
	"AGATZ - DoubleCenter - N075 - 88": "cmd/execute/instances/agatz/doublecenter/doublecenter-88-n75",
	"AGATZ - DoubleCenter - N075 - 89": "cmd/execute/instances/agatz/doublecenter/doublecenter-89-n75",
	"AGATZ - DoubleCenter - N075 - 90": "cmd/execute/instances/agatz/doublecenter/doublecenter-90-n75",
}
var AgatzInstanceDoubleCenterN100 = map[string]string{
	"AGATZ - DoubleCenter - N100 - 91":  "cmd/execute/instances/agatz/doublecenter/doublecenter-91-n100",
	"AGATZ - DoubleCenter - N100 - 92":  "cmd/execute/instances/agatz/doublecenter/doublecenter-92-n100",
	"AGATZ - DoubleCenter - N100 - 93":  "cmd/execute/instances/agatz/doublecenter/doublecenter-93-n100",
	"AGATZ - DoubleCenter - N100 - 94":  "cmd/execute/instances/agatz/doublecenter/doublecenter-94-n100",
	"AGATZ - DoubleCenter - N100 - 95":  "cmd/execute/instances/agatz/doublecenter/doublecenter-95-n100",
	"AGATZ - DoubleCenter - N100 - 96":  "cmd/execute/instances/agatz/doublecenter/doublecenter-96-n100",
	"AGATZ - DoubleCenter - N100 - 97":  "cmd/execute/instances/agatz/doublecenter/doublecenter-97-n100",
	"AGATZ - DoubleCenter - N100 - 98":  "cmd/execute/instances/agatz/doublecenter/doublecenter-98-n100",
	"AGATZ - DoubleCenter - N100 - 99":  "cmd/execute/instances/agatz/doublecenter/doublecenter-99-n100",
	"AGATZ - DoubleCenter - N100 - 100": "cmd/execute/instances/agatz/doublecenter/doublecenter-100-n100",
}
var AgatzInstanceDoubleCenterN175 = map[string]string{
	"AGATZ - DoubleCenter - N175 - 101": "cmd/execute/instances/agatz/doublecenter/doublecenter-101-n175",
	"AGATZ - DoubleCenter - N175 - 102": "cmd/execute/instances/agatz/doublecenter/doublecenter-102-n175",
	"AGATZ - DoubleCenter - N175 - 103": "cmd/execute/instances/agatz/doublecenter/doublecenter-103-n175",
	"AGATZ - DoubleCenter - N175 - 104": "cmd/execute/instances/agatz/doublecenter/doublecenter-104-n175",
	"AGATZ - DoubleCenter - N175 - 105": "cmd/execute/instances/agatz/doublecenter/doublecenter-105-n175",
	"AGATZ - DoubleCenter - N175 - 106": "cmd/execute/instances/agatz/doublecenter/doublecenter-106-n175",
	"AGATZ - DoubleCenter - N175 - 107": "cmd/execute/instances/agatz/doublecenter/doublecenter-107-n175",
	"AGATZ - DoubleCenter - N175 - 108": "cmd/execute/instances/agatz/doublecenter/doublecenter-108-n175",
	"AGATZ - DoubleCenter - N175 - 109": "cmd/execute/instances/agatz/doublecenter/doublecenter-109-n175",
	"AGATZ - DoubleCenter - N175 - 110": "cmd/execute/instances/agatz/doublecenter/doublecenter-110-n175",
}
var AgatzInstanceDoubleCenterN250 = map[string]string{
	"AGATZ - DoubleCenter - N250 - 1":   "cmd/execute/instances/agatz/doublecenter/doublecenter-1-n250",
	"AGATZ - DoubleCenter - N250 - 2":   "cmd/execute/instances/agatz/doublecenter/doublecenter-2-n250",
	"AGATZ - DoubleCenter - N250 - 3":   "cmd/execute/instances/agatz/doublecenter/doublecenter-3-n250",
	"AGATZ - DoubleCenter - N250 - 4":   "cmd/execute/instances/agatz/doublecenter/doublecenter-4-n250",
	"AGATZ - DoubleCenter - N250 - 5":   "cmd/execute/instances/agatz/doublecenter/doublecenter-5-n250",
	"AGATZ - DoubleCenter - N250 - 6":   "cmd/execute/instances/agatz/doublecenter/doublecenter-6-n250",
	"AGATZ - DoubleCenter - N250 - 7":   "cmd/execute/instances/agatz/doublecenter/doublecenter-7-n250",
	"AGATZ - DoubleCenter - N250 - 8":   "cmd/execute/instances/agatz/doublecenter/doublecenter-8-n250",
	"AGATZ - DoubleCenter - N250 - 9":   "cmd/execute/instances/agatz/doublecenter/doublecenter-9-n250",
	"AGATZ - DoubleCenter - N250 - 10":  "cmd/execute/instances/agatz/doublecenter/doublecenter-10-n250",
	"AGATZ - DoubleCenter - N250 - 111": "cmd/execute/instances/agatz/doublecenter/doublecenter-111-n250",
	"AGATZ - DoubleCenter - N250 - 112": "cmd/execute/instances/agatz/doublecenter/doublecenter-112-n250",
	"AGATZ - DoubleCenter - N250 - 113": "cmd/execute/instances/agatz/doublecenter/doublecenter-113-n250",
	"AGATZ - DoubleCenter - N250 - 114": "cmd/execute/instances/agatz/doublecenter/doublecenter-114-n250",
	"AGATZ - DoubleCenter - N250 - 115": "cmd/execute/instances/agatz/doublecenter/doublecenter-115-n250",
	"AGATZ - DoubleCenter - N250 - 116": "cmd/execute/instances/agatz/doublecenter/doublecenter-116-n250",
	"AGATZ - DoubleCenter - N250 - 117": "cmd/execute/instances/agatz/doublecenter/doublecenter-117-n250",
	"AGATZ - DoubleCenter - N250 - 118": "cmd/execute/instances/agatz/doublecenter/doublecenter-118-n250",
	"AGATZ - DoubleCenter - N250 - 119": "cmd/execute/instances/agatz/doublecenter/doublecenter-119-n250",
	"AGATZ - DoubleCenter - N250 - 120": "cmd/execute/instances/agatz/doublecenter/doublecenter-120-n250",
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
