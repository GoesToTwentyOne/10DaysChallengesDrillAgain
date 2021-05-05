package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	increment2 := func() int {
		x += 2
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment2())
}
