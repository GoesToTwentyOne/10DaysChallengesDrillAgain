package main

import "fmt"

func main() {
	slice1 := []int{4, 5, 6, 78}
	goes(slice1)

}
func goes(n []int) {
	var totalSum int
	for _, value := range n {
		totalSum += value

	}
	fmt.Println(totalSum)
}
