package main

import "fmt"

func main() {
	for i := 0; i < 256; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
		fmt.Print("------------------------next----------------------------")
		fmt.Printf("%d \t %b \t %#x \n", i, i, i)
	}
}
