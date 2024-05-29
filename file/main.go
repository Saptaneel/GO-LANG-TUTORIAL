package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	/*
		file, err := os.Create("sapta.txt")

		if err != nil {
			fmt.Println("Error :", err)
			return
		}
		defer file.Close()

		fmt.Println("Success")

		content := "HELLO WORLD"

		_, error := io.WriteString(file, content)

		if error != nil {
			fmt.Println("Error :", error)
			return
		}
		fmt.Println("Success")
	*/

	file, err := os.Open("sapta.txt")

	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error", err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}

	content, err := ioutil.ReadFile("sapta.txt")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(string(content))
}
