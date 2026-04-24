package main

import "fmt"

func main() {
	isLogged := true
	isAdmin := false
	hasSubScription := true

	canOpenDashboard := isLogged && hasSubScription

	canDeletePost := isAdmin || (isLogged && hasSubScription)

	fmt.Println(canOpenDashboard, canDeletePost)
}
