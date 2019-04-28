/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = 1.0
	for i := 1; i < 40; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z, math.Sqrt(x))
	}
	return
}

func main() {
	for i := 1; i < 100; i++ {
		fmt.Println(Sqrt(float64(i)))
		fmt.Println("----")
	}
}
