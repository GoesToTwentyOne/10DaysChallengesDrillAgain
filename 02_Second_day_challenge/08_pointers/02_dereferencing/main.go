package main

import "fmt"

func main() {
	var a *int
	var b int = 45
	a = &b

	fmt.Println(a)
	//dereferencing
	fmt.Println(*a)

}
