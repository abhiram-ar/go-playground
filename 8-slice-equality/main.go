package main

import (
	"bytes"
	"fmt"
)

func main() {
	slice := []byte{'a', 'b', 'c', 'd', 'e'}
	fmt.Printf("%v \n", slice)

	res1 := bytes.Equal(slice[0:2], []byte{'a', 'b'})
	res2 := bytes.Equal(slice[1:3], []byte{'a', 'b'})

	fmt.Println(res1)
	fmt.Println(res2)
}
