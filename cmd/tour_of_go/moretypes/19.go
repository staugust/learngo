/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var k = map[string]string{}
	if k == nil {
		fmt.Println("k is nil")
	}

	k["s"] = "s"
	fmt.Println(k)

	var arr []int
	arr = append(arr, 1)
	fmt.Println(arr)

}
