package main

import (
	image_generator "github.com/masakudou/image-generator"
	"image"
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

	input := []image_generator.Input{
		{
			Pos: []int{350, 200},
			Res: []int{720, 480},
		},
	}

	g, err := image_generator.NewGenerator(img, input)
	if err != nil {
		log.Fatal(err)
	}

	result, err := g.Generate()
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
