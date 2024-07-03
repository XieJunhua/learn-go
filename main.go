package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
)

const Pi = 3.14

var i = 1  //给定了initial value 可以省略类型，会自动进行推断
var j int

func add(x, y int) int {
	return x + y
}

// return multiple values
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int)(x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func test_defer()  {
	// defer 语句会将函数推迟到外层函数返回之后执行
	// defer会将调用压入到一个栈中，在外层函数返回的时候，会按照后进先出的顺序调用
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
}

func test_switch() {
	// switch的case语句从上到下顺序执行，直到匹配上为止。后面的不再执行
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:
			// freebsd, openbsd,
			// plan9, windows...
			fmt.Printf("%s.\n", os)
	}
}

func main() {
	fmt.Println("ddd")
	fmt.Println(rand.Intn(10))
	fmt.Println(math.Sqrt(16))
	fmt.Println(add(1, 2))

	a, b := swap("hello", "world") // 短变量声明只能用在函数中
	fmt.Println(a, b)

	x, y := split(17)
	fmt.Println(x, y)

	fmt.Println(i)
	fmt.Println(j)

	// 奇奇怪怪的for循环写法  大括号是必须的
	// 三个部分都是可选的。如果省略了, ;也可以省略
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum < 200 {
		sum += sum
	}
	fmt.Println(sum)

	test_switch()

	test_defer()
}
