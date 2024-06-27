package output

import (
	"fmt"
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

func drawBackgound(ggCtx *gg.Context) {
	// Pushing and Popping the context to avoid changing color outsite this function
	ggCtx.Push()
	defer ggCtx.Pop()

	ggCtx.SetColor(backgroundColor)
	ggCtx.Clear()
}

func drawStop(ggCtx *gg.Context, stop Stop) {
	// Pushing and Popping the context to avoid changing color outsite this function
	ggCtx.Push()
	defer ggCtx.Pop()

	ggCtx.SetColor(textColor)
	stopImage := loadStopImage(stop)
	stopImageHeight := stopImage.Bounds().Max.Y
	centerAxis := 0.5
	ggCtx.DrawImageAnchored(
		stopImage,
		int(axisX(stop)), int(axisY(stop)),
		centerAxis, centerAxis,
	)
	ggCtx.DrawString(
		stop.Name(),
		axisX(stop), axisY(stop)-float64(stopImageHeight/2),
	)
}

func drawMovement(ggCtx *gg.Context, actual Stop, next Stop, img image.Image, clr color.Color) {
	// Pushing and Popping the context to avoid changing color outsite this function
	ggCtx.Push()
	defer ggCtx.Pop()

	grad := movementGradient(actual, next, clr)
	ggCtx.SetStrokeStyle(grad)

	ggCtx.DrawLine(
		axisX(actual), axisY(actual),
		axisX(next), axisY(next),
	)
	ggCtx.Stroke()
	centerAxis := 0.5
	ggCtx.DrawImageAnchored(
		img,
		int((axisX(actual)+axisX(next))/2), int((axisY(actual)+axisY(next))/2),
		centerAxis, centerAxis,
	)
}

func drawInfos(ggCtx *gg.Context, routeDistance float64, routeTime float64) {
	// Pushing and Popping the context to avoid changing color outsite this function
	ggCtx.Push()
	defer ggCtx.Pop()

	ggCtx.SetColor(textColor)
	totalDistanceString := fmt.Sprintf("Total Distance: %.2f", routeDistance)
	totalTimeString := fmt.Sprintf("Total Time: %.2f", routeTime)
	_, totalDistanceHeight := ggCtx.MeasureString(totalDistanceString)
	_, totalTimeHeight := ggCtx.MeasureString(totalTimeString)
	ggCtx.DrawString(
		fmt.Sprintln("Total Distance: ", routeDistance),
		0, totalDistanceHeight,
	)
	ggCtx.DrawString(
		fmt.Sprintln("Total Time: ", routeTime),
		0, totalDistanceHeight+totalTimeHeight,
	)
}

func axisX(stop Stop) float64 {
	return applyScale(stop.Longitude() + paddingLeft)
}

func axisY(stop Stop) float64 {
	return applyScale(stop.Latitude() + paddingUp)
}

func applyScale(in float64) float64 {
	return in * applyScaleValue
}

func movementGradient(actual Stop, next Stop, initialColor color.Color) gg.Gradient {
	grad := gg.NewLinearGradient(axisX(actual), axisY(actual), axisX(next), axisY(next))
	grad.AddColorStop(0, initialColor)
	grad.AddColorStop(0.7, initialColor)
	grad.AddColorStop(1, endGradLineColor)
	return grad
}
