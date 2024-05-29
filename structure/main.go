package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}

func main() {
	var sap person
	//fmt.Println("sap person :", sap)
	sap.firstname = "Saptaneel"
	sap.lastname = "Chakraborty"
	sap.age = 22
	fmt.Println("sap person :", sap)

	person1 := person{
		firstname: "Virat",
		lastname:  "Kohli",
		age:       36,
	}
	fmt.Println(" person1 :", person1)

}
