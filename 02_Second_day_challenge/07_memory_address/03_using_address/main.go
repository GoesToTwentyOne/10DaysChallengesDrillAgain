package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("enter no 1 :")
	fmt.Scanln(&a)
	fmt.Println("enter no 2 :")
	fmt.Scanln(&b)

	fmt.Println("result  ", a+b)
}
