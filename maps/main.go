package main

import "fmt"

func main() {

	studentGrades := make(map[string]int)

	studentGrades["Virat"] = 90
	studentGrades["Rohit"] = 45
	studentGrades["Yashasvi"] = 77
	studentGrades["Rishabh"] = 40

	fmt.Println("Marks of Rohit :", studentGrades["Rohit"])

	delete(studentGrades, "Rohit")
	fmt.Println("Marks of Rohit :", studentGrades["Rohit"])

	grades, Exists := studentGrades["Rohit"]
	fmt.Println("Marks of Rohit :", grades)
	fmt.Println("Rohit exists :", Exists)

}
