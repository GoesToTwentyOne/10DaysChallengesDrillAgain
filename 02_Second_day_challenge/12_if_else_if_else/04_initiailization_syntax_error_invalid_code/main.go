package main

import "fmt"

func main() {
	b := true
	if food := "rice"; b {
		fmt.Println(food)
	}
	fmt.Println(b)
}
