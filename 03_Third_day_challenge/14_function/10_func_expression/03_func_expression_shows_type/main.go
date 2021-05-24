package main

import "fmt"

func main() {
	goes := func() {
		fmt.Println("Hello World")
	}
	goes()
	fmt.Printf(" %T \n", goes)

}
