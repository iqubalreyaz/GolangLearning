package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func main() {
	baseImage := image.NewRGBA(image.Rect(0, 0, 1000, 1000))

	file, err := os.Create("output.png")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, baseImage); err != nil {
		log.Fatalf("Error encoding image: %v", err)
	}
	addImage(baseImage, ".", image.Point{5, 5})
}

func addImage(baseImage *image.RGBA, path string, point image.Point) error {
	templateFile, err := os.Open(path)
	if err != nil {
		return err
	}

	template, _, err := image.Decode(templateFile)

	if err != nil {
		return err
	}

	draw.Draw(baseImage, baseImage.Bounds(), template, point, draw.Over)

	return nil
}

func addText(baseImage *image.RGBA, text string, point image.Point, col color.Color, fontSize float64) error {
	fontBytes, err := os.ReadFile("font.ttf")
	if err != nil {
		return err
	}

	ttf, err := opentype.Parse(fontBytes)
	if err != nil {
		return err
	}

	face, err := opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}

	drawer := &font.Drawer{
		Dst:  baseImage,
		Src:  image.NewUniform(col),
		Face: face,
		Dot: fixed.Point26_6{
			X: fixed.I(point.X),
			Y: fixed.I(point.Y),
		},
	}

	drawer.DrawString(text)

	return nil
}
