package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Foo interface {
	Abs() float64
	Scale(f float64)
}

func ys(f Foo) float64 {

	return f.Abs()
}

func dv(f Foo, kali float64) {
	f.Scale(kali)
}

func main() {
	v := Vertex{3, 4}

	v.Scale(10)

	dv(&v, 50)
	fmt.Println(v)
	fmt.Println(ys(&v))

}
