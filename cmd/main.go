package main

import (
	ig "github.com/masakudou/image-generator"
	ig_model "github.com/masakudou/image-generator/model"
	ig_components "github.com/masakudou/image-generator/components"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func main() {
	f, err := os.Open("resource/sample.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	components := []ig_model.ImageGeneratorComponent{}
	s := &ig_components.SquareFrame{
		Color: color.RGBA{
			R: 0,
			G: 255,
			B: 255,
			A: 255,
		},
		Rect: image.Rect(240, 340, 1240, 840),

		Width: 10,
	}
	a := &ig_components.AnotherSquareFrame{
		Color: color.RGBA{
			R: 255,
			G: 0,
			B: 255,
			A: 255,
		},
		Rect: image.Rect(440, 540, 1340, 1040),

		Width: 10,
	}
	components = append(components, s)
	components = append(components, a)

	g := ig.NewGenerator(img, components)

	result := g.Generate()
	if err != nil {
		log.Fatal(err)
	}

	of, err := os.Create("resource/result.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer of.Close()

	if err = jpeg.Encode(of, result, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatal(err)
	}
}
