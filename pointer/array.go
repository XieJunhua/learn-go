package main

import "fmt"

// 切片类似数组的引用 
// 更改切片的元素会修改其底层数组对应的元素

// golang中数组的大小是固定的，切片为数组元素提供了动态大小

// append(s []T, vs ...T)函数可以为切片动态增加元素


func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func testCap()  {

	// 切片的长度就是它所包含的元素个数。

// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0] // 截取使其长度为0
	printSlice(s)

	s = s[:4] // 截取使其长度为4
	printSlice(s)

	s = s[2:] // 舍弃前两个
	printSlice(s)

	a := make([]int, 5)  
	b := make([]int, 0, 5)

	printSlice(a)
	printSlice(b)
}

func main()  {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	s := primes[2:4]

	fmt.Println(s)
	fmt.Println(len(s))

	// []bool{true, true, false}  // 切片字面量

	testCap()

}