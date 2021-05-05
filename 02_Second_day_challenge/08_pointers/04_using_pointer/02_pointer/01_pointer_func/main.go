package main

import "fmt"

func main() {
	x := 69
	fmt.Println("before referencing  ", x)
	fmt.Println("before referencing  ", &x)
	goes(&x)
	fmt.Println("after referencing  ", x)
	fmt.Println("after referencing  ", &x)

}
func goes(z *int) {
	fmt.Println("before dereferencing  ", z)
	fmt.Println("before dereferencing  ", *z)
	*z = 1
	fmt.Println("after dereferencing  ", z)
	fmt.Println("after dereferencing  ", *z)

}
