package main

import "fmt"

func main() {
	switch "Alex" {
	case "Dejiya":
		fmt.Println("Whats up Dejiya")
	case "Tahsin":
		fmt.Println("Whats up Tahsin")
	case "Patric":
		fmt.Println("Whats up Patric")
	default:
		fmt.Println("you haven't any friend")
	}
}
