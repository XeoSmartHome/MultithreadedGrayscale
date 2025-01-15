package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path"
)

func processImage(fileName string) {
	inputFile, err := os.Open(getInputPath(fileName))

	if err != nil {
		log.Fatal("Error opening file", err)
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			log.Fatal("Error closing file", err)
		}
	}(inputFile)

	outputFile, err := os.Create(getOutputPath(fileName))

	if err != nil {
		log.Fatal("Error creating file", err)
	}

	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			log.Fatal("Error closing file", err)
		}
	}(outputFile)

	ext := path.Ext(fileName)

	var colorImage image.Image

	switch ext {
	case ".png":
		colorImage, err = png.Decode(inputFile)
		if err != nil {
			log.Fatal("Error decoding image", err)
		}
	case ".jpg", ".jpeg":
		colorImage, err = jpeg.Decode(inputFile)
		if err != nil {
			log.Fatal("Error decoding image", err)
		}
	default:
		log.Fatal("Unsupported file type", ext)
	}

	grayImage := image.NewGray(colorImage.Bounds())

	bounds := colorImage.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := colorImage.At(x, y).RGBA()

			white := (r + g + b) / 3

			grayImage.Set(x, y, color.Gray{Y: uint8(white >> 8)})
		}
	}

	err = png.Encode(outputFile, grayImage)
	if err != nil {
		log.Fatal("Error encoding image", err)
	}
}

func getInputPath(fileName string) string {
	return path.Join(InputDirectory, fileName)
}

func getOutputPath(fileName string) string {
	return path.Join(OutputDirectory, fileName)
}
