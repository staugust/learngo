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
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	var c rune
	res, err := fmt.Scanf("%c", &c)
	for err == nil {
		fmt.Printf("%d\n", uint32(c))
		res, err = fmt.Scanf("%c", &c)
		if c == 'q' {
			break
		}
	}
	fmt.Println(err)
	fmt.Scanln()

	var x int
	res, err = fmt.Scanf("%d", &x)
	for err == nil {
		fmt.Println(x, res, err)
		switch 0 {
		case x % 2:
			fmt.Println(" x is a multiple of 2")
		case x%2 - 1:
			fmt.Println(" divide x by 2 remains 1")
		}
		res, err = fmt.Scanf("%d", &x)
	}
	fmt.Println(err)
}
