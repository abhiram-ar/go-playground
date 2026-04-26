package main

import "fmt"

type User struct {
	id    int
	name  string
	email string
	age   int
}

func main() {
	u1 := User{
		id:    1,
		name:  "Abhiram",
		email: "abhiram@emai.com",
		age:   2,
	}

	fmt.Println(u1)
	fmt.Println(u1.email)

	// struct fields are mutable by default
	u1.age = 25
	fmt.Println(u1)

	// struct can be initilized partially
	// non-initialized fields will default to their zero value
	u2 := User{
		id:   2,
		name: "Jhon",
	}
	fmt.Println("partial user", u2)

}
