package main

import (
	"github.com/fogleman/gg"
	"image/color"
	"log"
	"render/components"
)

func main() {
	frameWidth := 1920
	frameHeight := 1080

	// Init frame
	frame := gg.NewContext(frameWidth, frameHeight)

	// Fonts
	font, err := gg.LoadFontFace("fonts/Roboto-Regular.ttf", 30)
	if err != nil {
		log.Fatal(err)
	}

	boldFont, err := gg.LoadFontFace("fonts/Roboto-Bold.ttf", 50)
	if err != nil {
		log.Fatal(err)
	}

	// Colors
	white := color.White
	red := color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
	green := color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 255,
	}

	// Clear frame
	frame.SetRGB(0.1, 0.1, 0.1)
	frame.Clear()


	// TEST
	rows := [][]components.Text {
		{
			components.NewText("", white, boldFont),
			components.NewText("Amazon", white, boldFont),
			components.NewText("Newegg", white, boldFont),
			components.NewText("Best Buy", white, boldFont),
			components.NewText("B&H", white, boldFont),
			components.NewText("Micro Center", white, boldFont),
		},
		{
			components.NewText("RTX 3060", white, boldFont),
			components.NewText("in stock!", green, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
		},
		{
			components.NewText("RTX 3060 TI", white, boldFont),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
		},
		{
			components.NewText("RTX 3070", white, boldFont),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
		},
		{
			components.NewText("RTX 3080", white, boldFont),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
		},
		{
			components.NewText("RTX 3090", white, boldFont),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
			components.NewText("out of stock", red, font),
		},
	}

	components.DrawTable(frame, rows, 0.0, 0.0, float64(frameWidth), float64(frameHeight), 5)
	// END TEST

	// Save img
	err = frame.SavePNG("test.png")

	if err != nil {
		log.Fatal(err)
	}
}
