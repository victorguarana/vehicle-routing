package output

import (
	"image"

	"github.com/fogleman/gg"
)

func loadCarImage() image.Image {
	im, err := gg.LoadPNG("assets/truck.png")
	if err != nil {
		panic(err)
	}
	return im
}

func loadDroneImage() image.Image {
	im, err := gg.LoadPNG("assets/drone.png")
	if err != nil {
		panic(err)
	}
	return im
}

func loadStopImage(stop Stop) image.Image {
	if stop.IsClient() {
		im, err := gg.LoadPNG("assets/client.png")
		if err != nil {
			panic(err)
		}
		return im
	}

	if stop.IsWarehouse() {
		im, err := gg.LoadPNG("assets/warehouse.png")
		if err != nil {
			panic(err)
		}
		return im
	}

	panic("Stop type not found")
}
