/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
	var k []int
	fmt.Println(reflect.DeepEqual(s, k))
	s = append(s, 1)
	k = append(k, 1)

	for i := 1; i < 100000; i++ {
		s = append(s, i)
		k = append(k, i)
	}
	fmt.Println(reflect.DeepEqual(s, k))

}
