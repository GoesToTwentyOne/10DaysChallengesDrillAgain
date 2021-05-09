package main

import "fmt"

func main() {
	value := []int{1, 2, 3, 4}
	goes(value...)

}

func goes(n ...int) {
	var totalSum int

	for _, value := range n {
		totalSum += value
	}
	fmt.Println(totalSum)
}
