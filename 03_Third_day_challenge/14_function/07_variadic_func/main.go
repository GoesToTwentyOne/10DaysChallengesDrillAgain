package main

import "fmt"

func main() {
	goes(4, 5, 6)

}
func goes(n ...int) {
	var total int

	for _, value := range n {
		total += value
	}
	fmt.Println(total)
	fmt.Printf("%T", n)
}
