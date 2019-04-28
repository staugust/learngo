/*
 * https://tour.golang.org/basics/4
 */

package main

import "fmt"

func add(a, b int32) int32 {
	return a + b
}

func main() {
	fmt.Println(add(42, 23))
	var a int32 = 0x7fffffff
	var b int32 = 0x7fffffff
	fmt.Printf("%d + %d = %d ", a, b, add(a, b))
}
