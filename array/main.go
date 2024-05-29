package main

import "fmt"

func main() {
	fmt.Println("Array")

	var name [5]string
	name[0] = "saptaneel"
	name[1] = "rohan"
	name[4] = "virat"

	fmt.Println("Names are ", name)
	fmt.Printf("Names are %q\n", name)

	var nos = [3]int{89, 47, 01}
	fmt.Println("Numbers are :", nos)

	var ans = len(nos)
	fmt.Println("length is :", ans)

}
