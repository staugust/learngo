# 成员函数什么时候选择指针作为接收者, 什么时候选择对象作为接收者?

```
There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both Scale and Abs are with receiver type *Vertex, even though the Abs method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)
```

# struct 与 interface转换
有下述代码, 可以看到, 如果以指针作为接收者, 则无论是`Point` 还是`*Point`均可赋值给*interface* `Abser` *a*. 而`Vertex`的`Abs`方法的接收者是对象, 则无法将`*Vertex`赋值给*interface* `Abser`.
```
import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v
	fmt.Println(a.Abs())
	p := Point{3, 4}
	a = p
	fmt.Println(a.Abs())
	a = &p
	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
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
```
