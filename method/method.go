package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (x MyFloat) Abs() float64  {
	if x < 0 {
		return float64(-x)
	} 

	return float64(x)
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


// 指针类型的接受者， 指针类型的接受者可以修改接受者的值
// 这里如果把*去掉，并不会报错，但是结果会是5,没有scale

func (v *Vertex) Scale(f float64)  { 
	v.X = v.X * f
	v.Y = v.Y * f
}

func testClass()  {
		var i interface{} = "hello"

	s := i.(string) 
	fmt.Println(s)

	s, ok := i.(string) // 类型断言，判断i的类型是否是string
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}
func main()  {

	v := Vertex{3, 4}
	fmt.Println(v.Abs())


	f := MyFloat(-4.0)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	p := &v
	fmt.Println(p.Abs()) // ==> p.Abs() = (*p).Abs() golang会自动解引用

	testClass()
	
}