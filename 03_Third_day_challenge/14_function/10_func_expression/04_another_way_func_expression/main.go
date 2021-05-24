package main

import "fmt"

func makegoes() func() string {
	return func() string {
		return "Hello World"
	}

}

func main() {
	greet := makegoes()
	fmt.Println(greet())

}
