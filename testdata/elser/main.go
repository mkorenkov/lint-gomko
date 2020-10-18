package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range values {
		if v%2 == 1 {
			fmt.Printf("%d is odd\n", v)
		} else {
			fmt.Printf("%d is even\n", v)
		}
	}
}
