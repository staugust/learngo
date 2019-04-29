/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	data [600][400][4]uint8
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(img.data), len(img.data[0]))
}
func (img Image) At(x int, y int) color.Color {
	return color.RGBA{
		R: img.data[x][y][0],
		G: img.data[x][y][1],
		B: img.data[x][y][2],
		A: img.data[x][y][3],
	}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
