package image_generator

import (
	"github.com/masakudou/image-generator/model"
	"image"
	"image/draw"
)

// Generator is main interface
type Generator interface {
	Generate() *image.RGBA
}

// NewGenerator returns Generator interface.
func NewGenerator(baseImage image.Image, components []model.ImageGeneratorComponent) Generator {
	return &generator{baseImage, components}
}

type generator struct {
	baseImage image.Image

	components []model.ImageGeneratorComponent
}

// Generate returns an image with the information drawn on the baseImage
// based on the information set in the components.
func (g *generator) Generate() *image.RGBA {
	imgRect := image.Rectangle{image.Pt(0, 0), g.baseImage.Bounds().Size()}
	img := image.NewRGBA(imgRect)
	draw.Draw(img, img.Bounds(), g.baseImage, image.Point{}, draw.Src)

	for _, v := range g.components {
		draw.DrawMask(img, img.Bounds(), v.FrameColor(), image.Point{}, v, image.Point{}, draw.Over)
	}

	return img
}
