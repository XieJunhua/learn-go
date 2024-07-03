package main

import "fmt"

/**
指针
Go 拥有指针。指针保存了值的内存地址。

类型 *T 是指向 T 类型值的指针，其零值为 nil。

var p *int
& 操作符会生成一个指向其操作数的指针。

i := 42
p = &i
* 操作符表示指针指向的底层值。

fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
这也就是通常所说的「解引用」或「间接引用」。
*/

type Vertex struct {
	X int
	Y int
}


func main() {
	a := 2
	b := &a
	fmt.Println(a, b)
	fmt.Println(*b)
	*b = 5
	fmt.Println(a, b)
	fmt.Println(*b)

	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	fmt.Println(v.X)

	p := &v // p is a pointer to a Vertex
	p.X = 1e9
	fmt.Println(v)
}
