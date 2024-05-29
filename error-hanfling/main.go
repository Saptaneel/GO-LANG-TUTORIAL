package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator zero hoyna")
	}
	return a / b, nil

}
func main() {
	ans, _ := divide(14, 0)
	//if err != nil {
	//fmt.Println("Error")

	fmt.Println("division result is ", ans)

}
