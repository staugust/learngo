/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	if s != "hello" {
		time.Sleep(time.Second)
	}
	defer fmt.Println(time.Now(), s)
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
