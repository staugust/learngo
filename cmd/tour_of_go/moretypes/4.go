/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func (v Vertex) print() {
	fmt.Printf("Vertex : {%d, %d}\n", v.X, v.Y)
}

func (v *Vertex) String() string {
	return fmt.Sprintf("Vertex    : {%d, %d}", v.X, v.Y)
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(&v)
}
