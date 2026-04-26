package main

import "fmt"

type Vertex struct {
	lat, long float64
}

// here we are only declaring a map
// the zero value of map will be nil
// if we assing a k-v to nil map we get an error
var m map[string]Vertex

func main() {
	// creating a map synt
	// map[keyType]valueType
	m = make(map[string]Vertex)
	m["Bell labs"] = Vertex{1.2, 3.3}
	fmt.Println(m["Bell labs"])

	// comma-ok idiom in golang
	// non exiting fieds and their zero value
	v, ok := m["hi"]
	fmt.Println(v, ok)

	// deleteing a value in map
	m["1"] = Vertex{1, 1}
	m["2"] = Vertex{2, 2}
	m["3"] = Vertex{3, 3}
	delete(m, "1")
	delete(m, "non-key-no-error")
	fmt.Println("deleted map", m)

	// looping through map
	for key, val := range m {
		fmt.Println("loop:", key, val)
	}

}
