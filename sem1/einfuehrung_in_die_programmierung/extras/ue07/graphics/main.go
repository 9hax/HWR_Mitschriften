package main

import (
	i "graphics/imaging"
	"image"
)

func main() {
	println("Uebung 07 - Grafik")

	img := i.CreateBasicGif(100, 100)

	i.DrawCircle(img, 50, 50, 300, true, 6)
	i.DrawCircle(img, 60, 150, 150, true, 4)

	drawNikolaus(img, 40, 60, 2)
	i.DrawLine(img, 60, 80, 60, 90, 3)
	i.DrawLine(img, 80, 85, 80, 95, 3)
	i.DrawSpiral(img, 55, 75, 5, 5, 10, 5, 2)

	drawClouds(img)
	drawBalloons(img)

	//drawNikolaus(img, 10, 40, 3)
	//drawFunnyCircles(img)
	//drawFunnyEllipses(img)
	//i.DrawBalloon(img, 60, 13, 20, 0, 4, 1)
	i.DrawEgg(img, 10, 70, 1)
	i.WriteImage(img)

}

func drawNikolaus(img *image.Paletted, x, y int, col uint8) {
	i.DrawRect(img, x, y+5, x+10, y+15, col)

	i.DrawLine(img, x, y+5, x+10, y+15, col)
	i.DrawLine(img, x, y+15, x+10, y+5, col)

	i.DrawLine(img, x, y+5, x+5, y, col)
	i.DrawLine(img, x+10, y+5, x+5, y, col)
}

func drawClouds(img *image.Paletted) {
	i.DrawCircle(img, 20, 20, 10, true, 1)
	i.DrawCircle(img, 30, 20, 20, true, 1)
	i.DrawCircle(img, 40, 25, 15, true, 1)

	i.DrawCircle(img, 60, 20, 15, true, 1)
	i.DrawEllipse(img, 70, 20, 30, 20, true, 1)
	i.DrawCircle(img, 85, 15, 20, true, 1)
}

func drawBalloons(img *image.Paletted) {
	i.DrawLine(img, 42+20, 50+20, 70, 80, 1)
	i.DrawLine(img, 50+20, 50+20, 70, 80, 1)
	i.DrawLine(img, 58+20, 50+20, 70, 80, 1)
	i.DrawBalloon(img, 35+20, 30+20, 20, 7, 4, 1)
	i.DrawBalloon(img, 65+20, 30+20, 20, -7, 5, 1)
	i.DrawBalloon(img, 50+20, 30+20, 25, 0, 3, 1)
}
