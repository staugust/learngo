/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	stuff := make(chan int, 5)
	go func() {
		for i := 0; i < 6; i++ {
			select {
			case stuff <- i:
				fmt.Printf("Sent %v\n", i)
			default:
				fmt.Printf("Default on %v\n", i)
			}
		}
		fmt.Println("Closing")
		close(stuff)
	}()
	time.Sleep(time.Second)
	fmt.Println(<-stuff)
	fmt.Println(<-stuff)
	fmt.Println(<-stuff)
	fmt.Println(<-stuff)
	fmt.Println(<-stuff)
}
