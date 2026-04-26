package main

import "fmt"

type Vertex struct {
	Long, Lang float32
}

// instead of declaring a map variable
// and initilizing it with make.
// if we know the values upfront, then we can declare and initalize in one go
var m = map[string]Vertex{
	"Bell labs": Vertex{1.1, 1.2},
	"Google":    {9.3, 5.2}, // can skip the typeDef(Vertex), since this is a composite type and there is redundancy
}

func main() {
	fmt.Println(m)
	fmt.Println(len(m))
}
