# string 内部实现
Run code shown below, guess what's the output?
```
	var a, b string
	a = "abc"
	b = "def"
	fmt.Println(a, b, &a, &b)
	a, b = b, a
	fmt.Println(a, b, &a, &b)
```
Output is :
```
abc def 0xc0420561c0 0xc0420561d0
def abc 0xc0420561c0 0xc0420561d0
```

The pointer of a, b keeps the same as before, while the content is swapped. 
Let's explore the inside details about string. 

