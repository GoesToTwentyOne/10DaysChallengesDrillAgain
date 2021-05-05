package main

import "fmt"

func main() {
	var a, b, c = 45, "Nihad", true
	fmt.Println(a, b, c)
	//automatin gaze/guess type
	fmt.Printf(" %T ", a)
	fmt.Printf(" %T ", b)
	fmt.Printf(" %T ", c)
}
