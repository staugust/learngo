/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	defer fmt.Println(time.Now(), " channel closed")
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		time.Sleep(time.Second)
		fmt.Println(time.Now(), i)
	}

	c = make(chan int, 5)
	for i := 1; i < 4; i++ {
		c <- i
	}
	close(c)
	for i := range c {
		time.Sleep(time.Second)
		fmt.Println(time.Now(), i)
	}
}
