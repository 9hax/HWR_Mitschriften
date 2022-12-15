package main

import (
	i "graphics/imaging"
)

func main() {
	println("Uebung 07 - Grafik")

	img := i.CreateBasicGif(100, 100)
	i.SetPixel(img, 50, 50, 1)
	i.SetPixel(img, 25, 25, 1)
	i.DrawLine(img, 10, 10, 10, 90, 1)

	i.WriteImage(img)

}
