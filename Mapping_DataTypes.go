package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"foo" : 42} //specifies a map of one element that has a key of 'foo' and number 42
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo") // deletes key foo
	fmt.Println(m)
}