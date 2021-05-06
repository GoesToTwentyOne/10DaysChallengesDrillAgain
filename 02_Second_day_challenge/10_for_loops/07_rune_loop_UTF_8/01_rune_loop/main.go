package main

import "fmt"

func main() {
	for i := 0; i <= 255; i++ {
		fmt.Println(i, " - ", string(rune(i)), " - ", []byte(string(rune(i))))
	}
	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
}
