package output

import (
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

	ggCtx.SetColor(clr)
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

func axisX(stop Stop) float64 {
	return applyScale(stop.Longitude() + paddingLeft)
}

func axisY(stop Stop) float64 {
	return applyScale(stop.Latitude() + paddingUp)
}

func applyScale(in float64) float64 {
	return in * applyScaleValue
}
