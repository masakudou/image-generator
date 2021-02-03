package model

import (
	"github.com/golang/freetype/truetype"
	"image"
)

type Component struct {
	Font      *truetype.Font
	FontSize  float64
	FontColor *image.Uniform

	FrameWidth int
	FrameColor *image.Uniform

	PosX int
	PosY int

	Width  int
	Height int
}

func (c *Component) SetFrame(xy, wh []int) {
	c.PosX = xy[0]
	c.PosY = xy[1]
	c.Width = wh[0]
	c.Height = wh[1]
}
