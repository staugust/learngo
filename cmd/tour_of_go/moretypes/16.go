/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import "fmt"

func main() {
	var pow = make([]int, 0)
	var i = 1
	var k = 1
	for i < 12 {
		pow = append(pow, k<<1)
		k = k << 1
		i += 1
	}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
