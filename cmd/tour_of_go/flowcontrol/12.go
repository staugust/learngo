/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleep() {
	start := time.Now()

	defer func() {
		fmt.Printf("program sleeped for %g seconds\n", float64(time.Now().Sub(start).Nanoseconds())/1000000000.0)
	}()
	time.Sleep(time.Second * time.Duration(rand.Int63n(10)))
}

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
	rand.Seed(time.Now().UnixNano())
	sleep()
}
