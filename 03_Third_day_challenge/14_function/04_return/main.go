package main

import "fmt"

func main() {
	value := goes("Nihad", "Hossain")
	fmt.Println(value)

}
func goes(fname, lname string) string {
	return fmt.Sprintf("My name is \"%s %s\"", fname, lname)

}
