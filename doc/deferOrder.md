# order of defer function

```
func cal() {
	defer func() {
		fmt.Println("first defer")
	}()
	defer fmt.Println("second defer")
	for i := 0; i < 10; i++ {
		defer fmt.Println(time.Now())
		time.Sleep(time.Second)
	}
}
```

Output is 

```
2019-04-28 17:35:15.0718526 +0800 CST m=+9.005911901
2019-04-28 17:35:14.0716882 +0800 CST m=+8.005748101
2019-04-28 17:35:13.071406 +0800 CST m=+7.005466501
2019-04-28 17:35:12.0710283 +0800 CST m=+6.005089301
2019-04-28 17:35:11.0707467 +0800 CST m=+5.004808301
2019-04-28 17:35:10.0703814 +0800 CST m=+4.004443601
2019-04-28 17:35:09.0701072 +0800 CST m=+3.004170001
2019-04-28 17:35:08.0697756 +0800 CST m=+2.003838901
2019-04-28 17:35:07.0694212 +0800 CST m=+1.003485001
2019-04-28 17:35:06.0689262 +0800 CST m=+0.002990701
second defer
first defer
```

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

To learn more about defer statements read this [blog post](https://blog.golang.org/defer-panic-and-recover).


