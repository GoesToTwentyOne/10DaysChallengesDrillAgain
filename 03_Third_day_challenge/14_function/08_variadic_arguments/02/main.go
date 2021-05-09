package main

import (
	"fmt"
	"reflect"
)

func main() {
	goes(4444, "Nihad", []int{4545454545, 777777777777, 88888888}, map[int]string{01: "Nihad", 02: "Dhaka"})

}
func goes(n ...interface{}) {
	for _, value := range n {
		fmt.Println(value, " ---- :", reflect.TypeOf(value).Kind())

	}

}
