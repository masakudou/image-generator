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

func (c *Component) SetFrame(x,y,w,h int) {
	c.PosX = x
	c.PosY = y
	c.Width = w
	c.Height = h
}
