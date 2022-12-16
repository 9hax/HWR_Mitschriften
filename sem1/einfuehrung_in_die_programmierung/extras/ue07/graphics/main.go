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

	//drawNikolaus(img, 10, 40, 3)
	//drawFunnyCircles(img)
	drawFunnyEllipses(img)
	//drawBalloons(img)
	//i.DrawBalloon(img, 60, 13, 20, 0, 4, 1)

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
	i.DrawCircle(img, 40, 50, 20, true, 1)
	i.DrawCircle(img, 70, 70, 70, false, 1)
	i.DrawCircle(img, 70, 70, 60, false, 2)
	i.DrawCircle(img, 70, 70, 50, false, 3)
	i.DrawCircle(img, 70, 70, 40, true, 4)
	i.DrawCircle(img, 70, 70, 30, false, 5)
}

func drawFunnyEllipses(img *image.Paletted) {
	i.DrawEllipse(img, 40, 50, 20, 30, true, 1)
	i.DrawEllipse(img, 70, 70, 70, 60, false, 1)
	i.DrawEllipse(img, 70, 70, 60, 70, false, 2)
	i.DrawEllipse(img, 70, 70, 50, 40, false, 3)
	i.DrawEllipse(img, 70, 70, 40, 50, true, 4)
	i.DrawEllipse(img, 70, 70, 30, 20, false, 5)
}

func drawBalloons(img *image.Paletted) {
	i.DrawLine(img, 42, 50, 50, 80, 1)
	i.DrawLine(img, 50, 50, 50, 80, 1)
	i.DrawLine(img, 58, 50, 50, 80, 1)
	i.DrawBalloon(img, 35, 30, 20, 7, 4, 1)
	i.DrawBalloon(img, 65, 30, 20, -7, 5, 1)
	i.DrawBalloon(img, 50, 30, 25, 0, 3, 1)
}
