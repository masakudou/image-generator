package components

import (
    "image"
    "image/color"
)

type AnotherSquareFrame struct {
    Color color.RGBA
    Rect  image.Rectangle

    Width int
}

func (f *AnotherSquareFrame) FrameColor() *image.Uniform {
    return &image.Uniform{f.Color}
}

func (f *AnotherSquareFrame) ColorModel() color.Model { return color.RGBAModel }

func (f *AnotherSquareFrame) Bounds() image.Rectangle { return f.Rect }

func (f *AnotherSquareFrame) At(x, y int) color.Color {
    if (f.Rect.Min.X+f.Width <= x) && (f.Rect.Max.X-f.Width >= x) {
        if (f.Rect.Min.Y+f.Width <= y) && (f.Rect.Max.Y-f.Width >= y) {
            return color.Alpha{100}
        }
    }

    return color.Alpha{255}
}
