/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func sending(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		c <- i
	}
}

func consuming(c <-chan int) {
	for {
		select {
		case x, val := <-c:
			fmt.Println(time.Now(), x, val)
		default:
			fmt.Println("default")
		}
		time.Sleep(time.Second)
		fmt.Println("run an iteration")
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
	c = make(chan int)
	go consuming(c)
	go sending(c)
	time.Sleep(time.Second * 20)

}
