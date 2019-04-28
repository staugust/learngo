/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"fmt"
)

func main() {
	var x []int
	for i := 1; i < 10; i++ {
		x = append(x, i)
	}

	y := x[1:6]
	z := x[3:7]
	cpy := make([]int, len(y))
	copy(cpy, y)
	for i := 0; i < len(y); i++ {
		y[i] += 1
	}

	for i := 0; i < len(z); i++ {
		z[i] *= 10
	}

	for i := 0; i < len(cpy); i++ {
		cpy[i] *= 100
	}

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(cpy)

}
