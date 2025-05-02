package decoder

import (
	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

//go:generate mockgen -source=decoder.go -destination=strategy_mock_test.go -package=decoder
type strategy interface {
	DefineVehicle(carList []vehicle.ICar, chromossome *brkga.Chromossome) (vehicle.ICar, vehicle.IDrone)
}
