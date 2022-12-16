package main

import (
	i "graphics/imaging"
	"image"
)

func main() {
	println("Uebung 07 - Grafik")

	img := i.CreateBasicGif(100, 100)
	i.SetPixel(img, 50, 50, 1)
	i.SetPixel(img, 25, 25, 1)
	i.DrawLine(img, 15, 10, 40, 0, 1)

	drawNikolaus(img, 10, 40, 3)
	drawFunnyCircles(img)

	i.WriteImage(img)

}

func drawNikolaus(img *image.Paletted, x, y int, col uint8) {
	// Box
	/*
		i.DrawLine(img, x, y+5, x+10, y+5, col)
		i.DrawLine(img, x, y+15, x+10, y+15, col)
		i.DrawLine(img, x, y+5, x, y+15, col)
		i.DrawLine(img, x+10, y+5, x+10, y+15, col)
	*/
	i.DrawRect(img, x, y+5, x+10, y+15, col)

	i.DrawLine(img, x, y+5, x+10, y+15, col)
	i.DrawLine(img, x, y+15, x+10, y+5, col)

	i.DrawLine(img, x, y+5, x+5, y, col)
	i.DrawLine(img, x+10, y+5, x+5, y, col)
}

func drawFunnyCircles(img *image.Paletted) {
	i.DrawCircle(img, 40, 50, 20, 1)
	i.DrawCircle(img, 70, 70, 70, 1)
	i.DrawCircle(img, 70, 70, 60, 2)
	i.DrawCircle(img, 70, 70, 50, 3)
	i.DrawCircle(img, 70, 70, 40, 4)
	i.DrawCircle(img, 70, 70, 30, 5)
}
