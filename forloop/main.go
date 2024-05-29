package main

import "fmt"

func main() {
	nos := []int{1, 2, 3, 4, 5}
	for index, value := range nos {
		fmt.Printf("index is %d, and value is %d\n", index, value)
	}
}
