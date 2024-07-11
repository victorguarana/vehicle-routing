package output

import (
	"image"
	"image/color"
	"math"

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
	drawArrow(ggCtx, actual, next)
	drawVehicle(ggCtx, actual, next, img)
}

func drawInfos(ggCtx *gg.Context, infos []Info) (float64, float64) {
	infosWidth := 0.0
	infosHeight := 0.0

	for _, info := range infos {
		width, height := ggCtx.MeasureString(info.Str)
		infosHeight += height
		infosWidth = math.Max(infosWidth, width)
		ggCtx.DrawString(info.Str, 0, infosHeight)
	}

	return infosWidth, infosHeight

}

func axisX(stop Stop) float64 {
	return applyScale(stop.Longitude())
}

func axisY(stop Stop) float64 {
	return applyScale(stop.Latitude())
}

func applyScale(in float64) float64 {
	return in * applyScaleValue
}

func drawArrow(ggCtx *gg.Context, actual Stop, next Stop) {
	// Pushing and Popping the context to avoid rotating outsite this function
	ggCtx.Push()
	defer ggCtx.Pop()

	arrowImg := loadArrowImage()
	ang := math.Atan2(axisY(actual)-axisY(next), axisX(actual)-axisX(next))
	centerAxis := 0.5
	posX := (axisX(actual) + axisX(next)*2) / 3
	posY := (axisY(actual) + axisY(next)*2) / 3
	// Subtracting Pi/2 because the arrow image is pointing up
	ggCtx.RotateAbout(ang-math.Pi/2, posX, posY)
	ggCtx.DrawImageAnchored(arrowImg, int(posX), int(posY), centerAxis, centerAxis)
}

func drawVehicle(ggCtx *gg.Context, actual Stop, next Stop, img image.Image) {
	centerAxis := 0.5
	posX := (axisX(actual) + axisX(next)) / 2
	posY := (axisY(actual) + axisY(next)) / 2
	ggCtx.DrawImageAnchored(img, int(posX), int(posY), centerAxis, centerAxis)
}
