package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v

	// (*p).X = 1000 --> while dereferencing a  struct pointer and accessing its fields is valid. it not neat
	p.X = 1000 // --> go allows to access the fields of a struct pointer wihtout explict dereferencing

	fmt.Println(v)

}
