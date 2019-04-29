/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"fmt"
	"go/types"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Point struct {
	X, Y float64
}

func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

type AugFloat float64

func (f AugFloat) Abs() float64 {
	return math.Abs(float64(f))
}

func do(i interface{}) {
	switch v := i.(type) {
	case float64:
		fmt.Printf("Float %T %v\n", v, v)
	//case AugFloat:
	//	fmt.Printf("AugFloat %T %v\n", v, v)
	case *Point:
		fmt.Printf("*Point %T %v\n", i, i)
	case Point:
		fmt.Printf("Point %T %v\n", i, i)
	case types.Pointer:
		fmt.Printf("Pointer %T %v\n", i, i)
	case Vertex:
		fmt.Printf("Vertex %T %v\n", i, i)
	case *Vertex:
		fmt.Printf("*Vertex %T %v\n", i, i)
	case Abser:
		fmt.Printf("Abser %T %v\n", i, i)
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
	do(Vertex{6, 8})
	do(&Vertex{6, 8})
	do(Point{3, 4})
	do(&Point{3, 4})

	var i Abser
	i = &Vertex{6, 8}
	do(i)
	i = Point{3, 4}
	do(i)
	i = &Point{3, 4}
	do(i)
	do(AugFloat(1))
	var af AugFloat
	do(&af)
}
