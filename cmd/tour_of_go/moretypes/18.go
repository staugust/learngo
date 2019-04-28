/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var h = 400
	var w = 600
	res := make([][]uint8, w)
	for i := 0; i < w; i++ {
		x := make([]uint8, h)
		for j := 0; j < h; j++ {
			x[j] = uint8(i*i*j%128 + i + j)
		}
		res[i] = x
	}
	return res
}

func main() {
	pic.Show(Pic)
}
