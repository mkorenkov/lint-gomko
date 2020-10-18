package main

import "fmt"

func main() {
	//nolint:test1 ignore
	a := 1 // nolint
	b := 2 //nolint
	c := 3 //nolint ignore smth
	d := 4 // nolint:smth
	fmt.Println(a + b + c + d)
}
