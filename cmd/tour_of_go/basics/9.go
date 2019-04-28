/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"fmt"
	"os"
)

type Test struct {
	name string
}

func NewTest(n string) Test {
	fmt.Println("initializing Test instance with name ", n)
	return Test{
		name: n,
	}
}

var t Test = NewTest("ahaha")

func main() {
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
	fmt.Println(t)
}
