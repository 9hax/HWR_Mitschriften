package imaging

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
	"strings"
	"unicode"
)

var palette []color.Color = []color.Color{
	color.Transparent,                  // 0 Base color
	color.White,                        // 1 White
	color.Black,                        // 2 Black
	color.RGBA{0xff, 0x00, 0x00, 0xff}, // 3 #f00f Full Red
	color.RGBA{0x00, 0xff, 0x00, 0xff}, // 4 #0f0f Full Green
	color.RGBA{0x00, 0x00, 0xff, 0xff}, // 5 #00ff Full Blue
	color.RGBA{0xff, 0x80, 0x80, 0xff}, // 6 #f88f Light Red
	color.RGBA{0x80, 0xff, 0x80, 0xff}, // 7 #8f8f Light Green
	color.RGBA{0x80, 0x80, 0xff, 0xff}, // 8 #88ff Light Blue
}

func AddColorToPalette(color color.Color) {
	// This allows adding a new color to a palette before initializing a new Image.
	palette = append(palette, color)
}

func CreateBasicGif(width, height int) *image.Paletted {
	// Create a new paletted image with the current palette.
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, palette)
	return img
}

func WriteGifImage(img *image.Paletted, filename string) {
	// Write the passed image to a gif file
	f, errors := os.Create(filename + ".gif")
	if errors != nil {
		panic(errors)
	}
	anim := gif.GIF{Delay: []int{0}, Image: []*image.Paletted{img}}
	gif.EncodeAll(f, &anim)
	f.Close()
}

func SetPixel(img *image.Paletted, x int, y int, col uint8) {
	// This wrapper unifies function calling across the library and allows setting pixels from outside the library.
	img.SetColorIndex(x, y, col)
}

func DrawLine(img *image.Paletted, x0 int, y0 int, x1 int, y1 int, col uint8) {
	// Bresenthams' Line-Drawing Algorithm

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
	// Get abs of x
	if x < 0 {
		return -x
	}
	return x
}

func DrawRect(img *image.Paletted, x0, y0, x1, y1 int, col uint8) {
	// Draw a rectangle by drawing 4 lines.
	DrawLine(img, x0, y0, x1, y0, col)
	DrawLine(img, x0, y1, x1, y1, col)
	DrawLine(img, x1, y0, x1, y1, col)
	DrawLine(img, x0, y0, x0, y1, col)
}

func DrawCircle(img *image.Paletted, x, y, dia int, filled bool, col uint8) {
	if filled {
		// Draw a filled circle by connecting the points on the circumference
		for i := 0; i < dia; i++ {
			for j := 0; j < dia; j++ {
				if math.Sqrt(float64((i-dia/2)*(i-dia/2)+(j-dia/2)*(j-dia/2))) < float64(dia/2) {
					img.SetColorIndex(x+i-dia/2, y+j-dia/2, col)
				}
			}
		}
	} else {
		// Draw only the circumference of the circle
		for i := 0; i < 360; i++ {
			angle := float64(i) * math.Pi / 180
			xx := int(float64(dia/2) * math.Cos(angle))
			yy := int(float64(dia/2) * math.Sin(angle))
			img.SetColorIndex(x+xx, y+yy, col)
		}
	}
}

func DrawBalloon(img *image.Paletted, x, y, dia, xshear int, col, col2 uint8) {
	// Draw a balloon by first drawing a circle, a light reflection and then drawing a lot of lines to make the body shape of the balloon.
	DrawCircle(img, x, y, dia, true, col)
	rad := dia / 2
	DrawCircle(img, x+(rad/2), y-(rad/2), rad/2, true, col2)
	for i := x - rad; i < x+rad; i++ {
		DrawLine(img, x+xshear, y+dia, i, y, col)
	}
}

func DrawEllipse(img *image.Paletted, x, y, xdia, ydia int, filled bool, col uint8) {
	if filled {
		// Draw a filled ellipse by connecting the points on the circumference
		for i := 0; i < xdia; i++ {
			for j := 0; j < ydia; j++ {
				if math.Sqrt(float64((i-xdia/2)*(i-xdia/2)+(j-ydia/2)*(j-ydia/2))) < float64((xdia+ydia)/4) {
					img.SetColorIndex(x+i-xdia/2, y+j-ydia/2, col)
				}
			}
		}
	} else {
		// Draw the circumference of the ellipse.
		for i := 0; i < 360; i++ {
			angle := float64(i) * math.Pi / 180
			xx := int(float64(xdia/2) * math.Cos(angle))
			yy := int(float64(ydia/2) * math.Sin(angle))
			img.SetColorIndex(x+xx, y+yy, col)
		}
	}
}

func DrawSpiral(img *image.Paletted, x, y, dia, length, spiralityx, spiralityy int, col uint8) {
	{
		// Draw the circumference of the spiral with an offset each iteration.
		for i := 0; i < 360*length; i++ {
			angle := float64(i) * math.Pi / 180
			xx := int(float64(dia/2) * math.Cos(angle))
			yy := int(float64(dia/2) * math.Sin(angle))
			sx := (i / (360 / spiralityx))
			sy := (i / (360 / spiralityy))
			img.SetColorIndex(x+xx+sx, y+yy+sy, col)
		}
	}
}

func DrawRune(img *image.Paletted, offsetx, offsety int, char rune, c uint8) {
	for y, row := range Font[unicode.ToUpper(char)] {
		for x, bit := range row {
			if bit {
				img.SetColorIndex(x+offsetx, y+offsety, c)
			}
		}
	}
}

func DrawString(img *image.Paletted, centerx, centery int, content string, c uint8) {
	content = strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(content, "ö", "oe", -1), "Ö", "Oe", -1), "ä", "ae", -1), "Ä", "Ae", -1), "ü", "ue", -1), "Ü", "Ue", -1)
	length := len(content)
	width := length * CharWidth
	startx := centerx - (width / 2)
	for x, char := range content {
		if char != ' ' {
			DrawRune(img, startx+(CharWidth*x), centery-CharMiddle, char, c)
		}
	}
}

func DrawPOI(img *image.Paletted, x, y int, c uint8) {
	// This draws a Point of Interest Icon (POI) using rectangles.
	DrawRect(img, x+2, y+5, x+12, y+25, c)
	DrawRect(img, x+28, y+5, x+38, y+25, c)
	DrawRect(img, x+13, y, x+27, y+25, c)
	DrawRect(img, x, y+25, x+40, y+28, c)

	DrawRect(img, x+5, y+8, x+9, y+12, c)
	DrawRect(img, x+5, y+16, x+9, y+20, c)
	DrawRect(img, x+31, y+8, x+35, y+12, c)
	DrawRect(img, x+31, y+16, x+35, y+20, c)

	DrawRect(img, x+15, y+2, x+19, y+6, c)
	DrawRect(img, x+19, y+2, x+25, y+6, c)
	DrawRect(img, x+15, y+8, x+21, y+12, c)
	DrawRect(img, x+21, y+8, x+25, y+12, c)

	DrawRect(img, x+16, y+15, x+24, y+25, c)

}
