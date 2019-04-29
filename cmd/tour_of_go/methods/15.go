/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import "fmt"

type AugInterface interface {
	print()
}
type TestInterface struct {
	val string
}

func (t TestInterface) print() {
	fmt.Println("TestInterface -> ", t.val)
}

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	if f, ok := i.(float64); ok {
		fmt.Println(f)
	} // panic

	t := TestInterface{
		val: "abc",
	}

	i = t
	fmt.Printf("%T %v\n", i, i)

	var k AugInterface = &t
	k.print()
	fmt.Printf("%T %v\n", k, k)

}
