package main

import "fmt"

func main() {

	//numbers := []int{1, 2, 3, 4}
	//numbers = append(numbers, 5, 6, 7, 8, 9)
	//fmt.Printf("Data type is %T\n", numbers)
	//fmt.Println("Numbers are :", numbers)

	numbers := make([]int, 3, 5)

	numbers = append(numbers, 5, 6, 7, 8, 9)
	fmt.Println("Capacity is :", cap(numbers))
	fmt.Println("Length is :", len(numbers))
	fmt.Println("Slice :", numbers)

}
