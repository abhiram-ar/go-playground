package main

import "fmt"

type Vertex struct {
	X, Y int
}

func (v Vertex) increment() Vertex {
	v.X = v.X + 1
	v.Y = v.Y + 1
	return v
}

func (v *Vertex) incrementP() *Vertex {
	v.X = v.X + 1
	v.Y = v.Y + 1
	return v
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	v1 := v.increment()
	fmt.Println(v)
	fmt.Println(v1)

	fmt.Println("")

	vp := Vertex{1, 2}
	fmt.Println(vp)

	vp2 := vp.incrementP()
	fmt.Println(vp)
	fmt.Println(*vp2)
}
