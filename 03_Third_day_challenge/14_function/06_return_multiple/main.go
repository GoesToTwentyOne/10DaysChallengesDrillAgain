package main

import "fmt"

func main() {
	final1, final2 := goes("Nihad", "Hossain")
	fmt.Println(final1, final2)

}
func goes(fname, lname string) (string, string) {
	f := fmt.Sprintf("My first name is %s\n", fname)
	l := fmt.Sprintf("your last name is %s", lname)
	return f, l

}
