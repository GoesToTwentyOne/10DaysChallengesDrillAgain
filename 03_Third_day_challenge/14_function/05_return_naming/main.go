package main

import "fmt"

func main() {
	finalValue := goes("Nihad", "Hossain")
	fmt.Println(finalValue)

}
func goes(fname, lname string) (s string) {
	s = fmt.Sprintf("My name is \"%s %s\"", fname, lname)
	return
}
