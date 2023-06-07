package main

import "fmt"

//variable is declared with capital value
//which signifies the variable is of type !!!!! PUBLIC !!!!

const Variable = "woieuf"

func main() {
	var username string = "Aashrith"
	fmt.Println("username is " + username)
	fmt.Printf("Variable is of type : %T \n", username)

	// implicit type
	fmt.Println("")
	var website = "something.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", username)

	// No var type
	fmt.Println("")
	value := 23
	fmt.Println(value)
}
