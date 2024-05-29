package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()

	ftime := time.Format("02-01-2006, 15:04:05, Monday")

	fmt.Println(time)
	fmt.Println(ftime)

	// Add 24 hours to the current time
	new := time.AddDate(0, 0, 3)

	fmt.Println(new)

	// Format the new time as a string
	ctime := new.Format("02-01-2006, 15:04:05, Monday")

	fmt.Println(new)
	fmt.Println(ctime)

}
