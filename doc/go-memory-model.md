# The Go Memory Model
Go内存模型讲述了在什么情况下, 一个goroutine中对变量的写操作在另一个goroutine的对变量的读操作中能观测到. 

## Happens Before
在一个goroutine中对变量的读写操作表现出以它们在源码中的先后顺序执行的. 换句话说, 编译器(compiler)和处理器(processors)可能重新调整读写操作的顺序, 但要保证这种顺序调整并不改变语言规范定义的goroutine内的行为. 

这种重新排序机制会导致在一个goroutine中观察到的读写顺序可能和另一个goroutine中观察到的不一样. 比如说, 一个goroutine执行代码`a = 1; b = 2;`, 可能在另外一个goroutine中先观察到`b`值更新发生在`a`值更新之前. 

为了指定读取和写入的要求，我们定义`happens before`，在Go程序中执行内存操作的部分顺序. 如果事件*e1*发生在事件*e2*之前，那么我们说*e2*发生在*e1*之后。 另外，如果*e1*既不在*e2*之前发生也不在*e2*之后发生，那么我们说*e1*和*e2*同时发生。

```
在单独一个goroutine中, `happens before`顺序即为程序源码所展示的顺序. 
```

如果以下两个条件成立, 则允许对变量*v*的读操作*r*观察到对*v*的写操作*w*:
1. 读操作*r*发生在写操作*w*之后
1. 在写操作*w*发生之后, 读操作*r*发生之前, 没有其他的对变量*v*的写操作.

为了保证对变量*v*的读操作*r*观察到对*v*的特定写操作*w*，确保*w*是允许*r*观察到的唯一写操作, 也就是说, 保证*r*观察到*w*必须满足以下两条件: 
1. *w*发生在*r*之前.
1. 对共享变量*v*的任何其他写操作要么发生在*w*之前, 要么发生在*r*之后. 

这两条比之前提交的两条的约束更强, 它要求没有其他写操作与*w*或*r*同时发生. 

在单个goroutine中, 没有并发, 因此这两个定义是等价的: 读操作*r*观察到最近的写操作*w*对变量*v*写入的值. 当多个goroutine访问共享变量*v*时，它们必须使用同步事件来建立`happens before`条件, 以确保读操作能观察到想要观察到的写操作. 

以变量类型的零值(`zero`)对变量*v*的初始化在内存模型中视为对变量*v*的一次写操作. 

大于单个机器字长(a single machine word)的值的读操作和写操作可视为以未指定的顺序进行的多次以单个机器字长度进行的读/写操作. 

## Synchronization 同步

### Initialization初始化
程序的初始化运行在一个单独的goroutine中, 但是这个goroutine可能创建新的goroutines, 这些goroutines可能同时运行.
如果一个包*p*导入另外一个包*q*, *q*的`init`方法执行完毕之后, 才会执行*p*的`init`方法. 
`main.main`(main包的main方法)在所有的`init`方法执行完毕之后, 才会开始执行. 

### Goroutine creation

`go`语句创建的goroutine在go语句执行完毕之后才开始执行. 
考虑到如下代码:
```
var a string

func f() {
	print(a)
}

func hello() {
	a = "hello, world"
	go f()
}
```
调用函数`hello`, 会在之后的某一时刻打印`"hello, world"`(极有可能在hello函数已经返回之后). 

### Goroutine destruction
一个goroutine的退出时间是不可确定的, 不能确保goroutine在程序的特定时间之前退出. 有如下代码:
```
var a string

func hello() {
	go func() { a = "hello" }()
	print(a)
}
```
对*a*的赋值操作并未跟随任何同步事件, 所以并不能保证在其他goroutine中能看到对*a*的赋值. 事实上, 激进的编译器甚至可能删除整个go语句. 
如果一个goroutine中对变量的更改需要被其他goroutine观察到, 需要使用一种同步机制, 比如说`lock`或者`channel communication`, 从而使操作相对有序.

### Channel communication
`Channel通信`是goroutines间实现同步机制的常见方法. 对特定channel的每个send操作都会匹配到相应的在此channel上的receive操作, 通过send和receive在不同的goroutine. 

在一个channel上的send操作再相应的receive操作之前发生. 换句话说, 数据总是先被send到channel, 才能从channel中receive此数据. 
有代码如下:
```
var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	c <- 0
}

func main() {
	go f()
	<-c
	print(a)
}
```
这段代码保证会打印`"hello, world"`. 对*a*的写操作发生在send `0`到c上之前, 也就先于从c上读取到此数据, 而`print`则在从c中读取此数据之后. 
The closing of a channel happens before a receive that returns a zero value because the channel is closed.
channel的关闭操作发生在从channel上receive到因channel关闭而产生的零值之前. 将上述代码`f`中的`c <- 0`改为 `close(c)`, 也能保证`print`打印的是`"hello, world"`. 
验证代码如下:
```
import "fmt"

var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	c <- 1
	close(c)
}

func main() {
	go f()
	for i := 0; i < 10; i++ {
		k, ok := <-c
		fmt.Println(k, ok)
	}
	fmt.Println(a)
}
```
输出为:
```
1 true
0 false
0 false
0 false
0 false
0 false
0 false
0 false
0 false
0 false
hello, world
```

无缓存的channel( `unbuffered channel`)中, receive操作发生在相应的send操作完成之前. 考虑到以下代码:
```
var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	<-c
}
func main() {
	go f()
	c <- 0
	print(a)
}
```
其仍能保证`print`打印的是`"hello, world"`. 这是因为对*c*的写入发生在对*c*的读取之前, 而对*c*的读取又在对*c*的写入完成之前发生, 而print发生在对*c*的写入完成之后, 对*a*的赋值操作发生在对*c*的读取之前, 所以对*a*的赋值操作发生在`print`之前, 故肯定会打印`"hello, world". 

如果channel是有缓存的, 比如说` c:= make(chan int, 1)`, 则上述代码不能保证必然打印`"hello, world"`. 程序可能输出空字符串, 崩溃, 或者其他异常行为. 

从无缓冲channel推广到有缓冲channel, 有
```
在容量为*C*的channel上的第`k`次读取操作发生在对该channel的第`k + C`次写入完成之前. 
```
可以由缓冲channel构建出计数信号量：信道中的项目数量对应于活动使用次数，信道容量对应于最大同时使用次数，发送项目获取信号量，以及 接收项目会释放信号量。 这是限制并发的常用习惯用法。

该程序为工作列表中的每个条目启动一个goroutine，但goroutine使用限制通道进行协调，以确保一次最多有三个正在运行工作函数。
```
var limit = make(chan int, 3)

func main() {
	for _, w := range work {
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select{}
}
```

### Locks 锁
`sync`包提供了两种锁, `sync.Mutex`和`sync.RWMutex`. 

### Once
`sync`包提供了`Once`用于初始化. 



