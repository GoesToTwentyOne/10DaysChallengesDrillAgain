package main

import "fmt"

type contact struct {
	name    string
	address string
}

func main() {
	switchontype("Nihad")
	switchontype(7)
	stype := contact{name: "Nihad", address: "Hazipara"}
	switchontype(stype)
	switchontype(stype.name)
	switchontype(stype.address)

}
func switchontype(x interface{}) {
	switch x.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float")
	case bool:
		fmt.Println("bool")

	default:
		fmt.Println("unknown type")

	}

}
