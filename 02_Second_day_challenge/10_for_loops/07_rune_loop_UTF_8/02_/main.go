package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(rune(i)), []byte(string(rune(i))))
	}
}
