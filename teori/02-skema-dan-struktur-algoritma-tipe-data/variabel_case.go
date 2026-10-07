package main

import "fmt"

func main(){
	var name string
	var middleName string
	var lastName string

	name = "Faiq"
	middleName = "Fadhlul"
	lastName = "Aziz"

	fmt.Println("Name:", name)
	fmt.Println("Middle Name:", middleName)
	fmt.Println("Last Name:", lastName)

	var (
		fullName = "faiq fadhlul aziz"
	)

	fmt.Println("Full Name:", fullName)
}