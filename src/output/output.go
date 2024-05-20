package output

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/routes"
)

var textColor = color.Black
var backgroundColor = color.White
var flightLineColor = color.RGBA{0, 0, 255, 255}
var routeLineColor = color.RGBA{0, 255, 0, 255}
var endGradLineColor = color.RGBA{255, 0, 0, 255}

const imageSize = 550
const paddingLeft = 50
const paddingUp = 55
const applyScaleValue = 5
const mainLineWidth = 0.3

type Stop interface {
	IsClient() bool
	IsWarehouse() bool
	Latitude() float64
	Longitude() float64
	Name() string
}

func ToImage(fileName string, itn itinerary.Itinerary) {
	ggCtx := gg.NewContext(imageSize, imageSize)
	drawBackgound(ggCtx)
	setRouteValues(ggCtx)
	itineraryToImage(ggCtx, itn)
	err := ggCtx.SavePNG(fileName)
	if err != nil {
		panic(err)
	}
}

func setRouteValues(ggCtx *gg.Context) {
	ggCtx.SetLineWidth(applyScale(mainLineWidth))
}

func itineraryToImage(ggCtx *gg.Context, itn itinerary.Itinerary) {
	iterator := itn.RouteIterator()
	for iterator.HasNext() {
		actual := iterator.Actual()
		next := iterator.Next()
		drawMovement(ggCtx, actual, next, loadCarImage(), routeLineColor)
		drawStop(ggCtx, actual)
		flightsToImage(ggCtx, actual.StartingSubRoutes())
		iterator.GoToNext()
	}
	drawStop(ggCtx, iterator.Actual())
}

func flightsToImage(ggCtx *gg.Context, flights []routes.ISubRoute) {
	for _, flight := range flights {
		drawMovement(ggCtx, flight.StartingStop(), flight.First(), loadDroneImage(), flightLineColor)
		iterator := flight.Iterator()
		for iterator.HasNext() {
			actual := iterator.Actual()
			next := iterator.Next()
			drawMovement(ggCtx, actual, next, loadDroneImage(), flightLineColor)
			drawStop(ggCtx, actual)
			iterator.GoToNext()
		}
		drawMovement(ggCtx, flight.Last(), flight.ReturningStop(), loadDroneImage(), flightLineColor)
		drawStop(ggCtx, flight.Last())
	}
}
