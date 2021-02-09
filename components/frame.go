package components

import (
	"image"
	"image/color"
)

type SquareFrame struct {
	Color color.RGBA
	Rect  image.Rectangle

	Width int
}

func (f *SquareFrame) FrameColor() *image.Uniform {
	return &image.Uniform{f.Color}
}

func (f *SquareFrame) ColorModel() color.Model { return color.RGBAModel }

func (f *SquareFrame) Bounds() image.Rectangle { return f.Rect }

func (f *SquareFrame) At(x, y int) color.Color {
	if (f.Rect.Min.X+f.Width <= x) && (f.Rect.Max.X-f.Width >= x) {
		if (f.Rect.Min.Y+f.Width <= y) && (f.Rect.Max.Y-f.Width >= y) {
			return color.Alpha{0}
		}
	}

	return color.Alpha{255}
}
