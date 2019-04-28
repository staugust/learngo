/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func counter(c *int) func(int) int {
	return func(x int) int {
		return (*c) * x
	}
}

func main() {
	var i int = 0
	pos, neg := adder(), adder()
	cnt := counter(&i)
	for i = 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
			cnt(i),
		)
	}
}
