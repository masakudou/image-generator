package model

import (
	"image"
	"image/color"
)

// Frame defines a rectangle with a Width
type Frame struct {
	Rect image.Rectangle

	Width int
}

func (f *Frame) ColorModel() color.Model { return color.RGBAModel }

func (f *Frame) Bounds() image.Rectangle { return f.Rect }

func (f *Frame) At(x, y int) color.Color {
	if (f.Rect.Min.X+f.Width <= x) && (f.Rect.Max.X-f.Width >= x) {
		if (f.Rect.Min.Y+f.Width <= y) && (f.Rect.Max.Y-f.Width >= y) {
			return color.Alpha{0}
		}
	}

	return color.Alpha{255}
}
