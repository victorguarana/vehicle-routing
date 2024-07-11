package output

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/route"
)

var primaryColor = color.Black
var backgroundColor = color.White
var flightLineColor = color.RGBA{0, 0, 255, 255}
var routeLineColor = color.RGBA{0, 255, 0, 255}

const imageSize = 550
const paddingLeft = 250
const paddingUp = 250
const applyScaleValue = 5
const mainLineWidth = 1.5

type Info struct {
	Str string
}

type Stop interface {
	IsClient() bool
	IsWarehouse() bool
	Latitude() float64
	Longitude() float64
	Name() string
}

func ToImage(fileName string, itineraryInfo itinerary.Info, infos []Info) {
	ggCtx := gg.NewContext(imageSize, imageSize)
	ggCtx.SetLineWidth(mainLineWidth)
	ggCtx.SetColor(primaryColor)

	drawBackgound(ggCtx)
	_, infosHeight := drawInfos(ggCtx, infos)
	ggCtx.Translate(paddingLeft, paddingUp+infosHeight)
	itineraryToImage(ggCtx, itineraryInfo)
	err := ggCtx.SavePNG(fileName)
	if err != nil {
		panic(err)
	}
}

func itineraryToImage(ggCtx *gg.Context, itineraryInfo itinerary.Info) {
	iterator := itineraryInfo.RouteIterator()
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

func flightsToImage(ggCtx *gg.Context, flights []route.ISubRoute) {
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
