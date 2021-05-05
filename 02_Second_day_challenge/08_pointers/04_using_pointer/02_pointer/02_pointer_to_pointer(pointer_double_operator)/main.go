package main

import "fmt"

func main() {
	var a int = 45
	var singlePointer *int = &a
	var pointerTopointer **int = &singlePointer
	fmt.Println(a)
	fmt.Println(singlePointer)
	fmt.Println(*singlePointer)
	fmt.Println("-------------------------------------------")
	fmt.Println(pointerTopointer)
	fmt.Println(**pointerTopointer)

}
