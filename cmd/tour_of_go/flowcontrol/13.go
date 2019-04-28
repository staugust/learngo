/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"fmt"
	"time"
)

func cal() {
	defer func() {
		fmt.Println("first defer")
	}()
	defer fmt.Println("second defer")
	for i := 0; i < 10; i++ {
		defer fmt.Println(time.Now())
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("counting")
	cal()
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
