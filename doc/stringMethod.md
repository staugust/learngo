# an object's String method 
Given code below:

```go
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

```

The output is 
```
{1000000000 2}
Vertex    : {1000000000, 2}
```

When `v.String` function is declared as `func (v Vertex) String() string`,  the output is :
```
Vertex    : {1000000000, 2}
Vertex    : {1000000000, 2}
```

This is the difference of receiver of functions. 

