package model

import (
	"image"
	"image/color"
)

// ImageGeneratorComponent is component interface
type ImageGeneratorComponent interface {
	FrameColor() *image.Uniform

	ColorModel() color.Model
	Bounds() image.Rectangle
	At(x, y int) color.Color
}
