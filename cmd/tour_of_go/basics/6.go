/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"fmt"
	"unsafe"
)

func swap(a, b string) (string, string) {
	return b, a
}

func pswap(a, b *string) (*string, *string) {
	return b, a
}

func main() {
	var a, b string = "abc", "def"
	fmt.Printf("%s %s %v %v\n", a, b, &a, &b)
	a, b = swap(a, b)
	fmt.Printf("%s %s %v %v\n", a, b, &a, &b)
	x, y := pswap(&a, &b)
	fmt.Printf("%s %s %v %v\n", *x, *y, x, y)
	var str = "abcdef"
	fmt.Printf("%s %v\n", str, &str)
	str = "defxxx"
	fmt.Printf("%s %v\n", str, &str)
	ps := unsafe.Pointer(&str)
	fmt.Printf("%v\n", ps)
	c := a
	fmt.Printf("%s %s %s %v %v %v", a, b, c, &a, &b, &c)
}
