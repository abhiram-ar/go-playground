package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {

	u := User{name: "abhiram", age: 20}
	fmt.Println(u.intro())

}

// value reciver => means this methods receices a copy of the use
func (u User) intro() string {
	return fmt.Sprintf("Hi, I am %s", u.name)
}
