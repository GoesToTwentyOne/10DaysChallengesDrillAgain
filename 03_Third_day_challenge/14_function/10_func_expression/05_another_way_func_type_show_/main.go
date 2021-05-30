package main

import "fmt"

func goes() func() string {
	return func() string {
		return "I'm Nihad Hossain"
	}

}
func main() {
	gfinal := goes()
	fmt.Println(gfinal())
	fmt.Printf("%T \n", gfinal)

}
