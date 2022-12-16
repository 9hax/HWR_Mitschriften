package imaging

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

func CreateBasicGif(width, height int) *image.Paletted {
	palette := []color.Color{
		color.Transparent,
		color.White,
		color.Black,
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff}}
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
	// Find the difference between the x and y coordinates
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)

	// Determine which coordinate has the greatest difference
	var sx, sy int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}

	// Initialize the error value to 0
	err := dx - dy

	for {
		// Set the pixel at the current position
		SetPixel(img, x0, y0, col)

		// If the current position is the same as the ending position, the line is drawn
		if x0 == x1 && y0 == y1 {
			break
		}

		// Calculate the new error value
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
func DrawLine(img *image.Paletted, x0 int, y0 int, x1 int, y1 int, col uint8) {
	dx := x1 - x0
	dy := y1 - y0
	D := 2*dy - dx
	y := y0
	for x := x0; x < x1; x++ {
		println("Setting Pixel at ", x, ",", y, "to", col)
		SetPixel(img, x, y, col)
		if D > 0 {
			y = y + 1
			D = D - 2*dx
		}
		D = D + 2*dy
	}
}*/

func DrawRect(img *image.Paletted, x0, y0, x1, y1 int, col uint8) {
	DrawLine(img, x0, y0, x1, y0, col)
	DrawLine(img, x0, y1, x1, y1, col)
	DrawLine(img, x1, y0, x1, y1, col)
	DrawLine(img, x0, y0, x0, y1, col)
}

func DrawCircle(img *image.Paletted, x, y, dia int, col uint8) {
	// Calculate the radius of the circle
	radius := dia / 2
	// Iterate over the bounding box of the circle
	for i := x - radius; i <= x+radius; i++ {
		for j := y - radius; j <= y+radius; j++ {
			// If the point is close to the radius, set the pixel
			if abs((i-x)*(i-x)+(j-y)*(j-y)-(radius*radius)) <= radius {
				SetPixel(img, i, j, col)
			}
		}
	}
}
