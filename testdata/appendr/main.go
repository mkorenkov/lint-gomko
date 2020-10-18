package main

import "fmt"

func main() {
	values1 := []int{1, 2, 3}
	values2 := []int{4, 5, 6}
	values3 := append(values1, values2...)
	fmt.Println(values3)
}
