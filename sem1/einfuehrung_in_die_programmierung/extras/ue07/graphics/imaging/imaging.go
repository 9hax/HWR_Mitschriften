package imaging

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

func CreateBasicGif(width, height int) *image.Paletted {
	palette := []color.Color{color.Transparent, color.White, color.Black}
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, palette)
	return img
}

func WriteImage(img *image.Paletted) {
	// Programmcode für das Zeichnen
	f, errors := os.Create("myImage.gif")
	if errors != nil {
		panic(errors)
	}
	anim := gif.GIF{Delay: []int{0}, Image: []*image.Paletted{img}}
	gif.EncodeAll(f, &anim)
	f.Close()
}

func SetPixel(img *image.Paletted, x int, y int, col uint8) {
	img.SetColorIndex(x, y, col)
}

func DrawLine(img *image.Paletted, x0 int, y0 int, x1 int, y1 int, col uint8) {
	dx := x1 - x0
	dy := y1 - y0
	D := 2*dy - dx
	y := y0
	for x := x0; x < x1; x++ {
		img.SetColorIndex(x, y, col)
		if D > 0 {
			y = y + 1
			D = D - 2*dx
		}
		D = D + 2*dy
	}
}

func DrawRect(img *image.Paletted, x0, y0, x1, y1 int, col uint8) {
	DrawLine(img, x0, y0, x1, y0, col)
	DrawLine(img, x0, y1, x1, y1, col)
	DrawLine(img, x1, y0, x1, y1, col)
	DrawLine(img, x0, y0, x0, y1, col)
}
