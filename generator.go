package image_generator

import (
	"github.com/masakudou/image-generator/model"
	"golang.org/x/xerrors"
	"image"
	"image/color"

	"image/draw"
)

// Generator is main interface
type Generator interface {
	Generate() (*image.RGBA, error)
}

// Input uses for component config
type Input struct {
	PosX int
	PosY int
	Width int
	Height int
}

// NewGenerator returns Generator interface.
func NewGenerator(baseImage image.Image, componentConfig []Input) (Generator, error) {
	generator := &generator{}
	generator.baseImage = baseImage

	if len(componentConfig) == 0 {
		return nil, xerrors.New("no components config input")
	}

	for _, v := range componentConfig {
		comp := &model.Component{}
		comp.SetFrame(v.PosX, v.PosY, v.Width, v.Height)
		comp.FrameWidth = 10

		generator.components = append(generator.components, comp)
	}

	return generator, nil
}

type generator struct {
	baseImage image.Image

	components []*model.Component
}

// Generate returns an image with the information drawn on the baseImage
// based on the information set in the components.
func (g *generator) Generate() (img *image.RGBA, err error) {
	imgRect := image.Rectangle{image.Pt(0, 0), g.baseImage.Bounds().Size()}
	img = image.NewRGBA(imgRect)
	draw.Draw(img, img.Bounds(), g.baseImage, image.Point{}, draw.Src)

	for _, v := range g.components {
		// TODO: to frameColor
		src := &image.Uniform{color.RGBA{0, 0, 255, 255}}

		frame := &model.Frame{
			Rect:  image.Rect(v.PosX, v.PosY, v.PosX+v.Width, v.PosY+v.Height),
			Width: v.FrameWidth,
		}
		draw.DrawMask(img, img.Bounds(), src, image.Point{}, frame, image.Point{}, draw.Over)
	}

	return img, nil
}
