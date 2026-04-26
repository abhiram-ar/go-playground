package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {

	u := User{name: "abhiram", age: 20}
	fmt.Println("before:", u.age)
	fmt.Println(u.Birthday())
	fmt.Println("after:", u.age)

}

// pointer reciver => means this methods receices a reference of the use
func (u *User) Birthday() bool {
	u.age++
	return true

}
