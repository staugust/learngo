/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import "fmt"

var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	c <- 1
	close(c)
}

func main() {
	go f()
	for i := 0; i < 10; i++ {
		k, ok := <-c
		fmt.Println(k, ok)
	}
	fmt.Println(a)
}
