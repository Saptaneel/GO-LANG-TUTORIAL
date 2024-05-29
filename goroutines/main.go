package main

import (
	"fmt"
	"time"
)

func sayhello() {
	fmt.Println("Hello")
	//time.Sleep(4000 * time.Millisecond)
	fmt.Println("Say Hello function")
}

func sayhi() {
	fmt.Println("goroutines")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Say Hi function")
}
func main() {
	fmt.Println("Hello Saptaneel")
	go sayhello()
	go sayhi()

	time.Sleep(1000 * time.Millisecond)

}
