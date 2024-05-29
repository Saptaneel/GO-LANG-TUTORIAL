package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("URL")

	myurl := "https://jsonplaceholder.typicode.com/todos/1"

	parse, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(parse.Scheme)
	fmt.Println(parse.Host)
	fmt.Println(parse.Path)
	fmt.Println(parse.RawQuery)

	parse.Path = "/saptaneel"
	parse.RawQuery = "user=sap"

	newurl := parse.String()

	fmt.Println(newurl)

}
