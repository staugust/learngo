/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	i := 10
	sum = 0
	for i > 0 {
		sum += i
		i--
	}
	fmt.Println(sum)

	sum = 0
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, val := range arr {
		sum += val
	}
	fmt.Println(sum)

	i = 10
	sum = 0
	for {
		sum += i
		i--
		if i < 0 {
			break
		}
	}
	fmt.Println(sum)
}
