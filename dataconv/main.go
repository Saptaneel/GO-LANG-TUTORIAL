package main

import (
	"fmt"
	"strconv"
)

func main() {
	no := "12340"
	noint, _ := strconv.Atoi(no)
	nofloat, _ := strconv.ParseFloat(no, 64)

	fmt.Println(noint)
	fmt.Println(nofloat)
	fmt.Println(nofloat)
}
