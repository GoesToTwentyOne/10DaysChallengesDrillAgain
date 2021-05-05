package main

import "fmt"

var x = 0

func increment() int {
	x++
	return x
}
func increment2() int {
	x += 2
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment2())
	fmt.Println(increment())
	fmt.Println(increment2())
}
