package main

import "fmt"

func tester() {
	var name string
	fmt.Print("Input Your Name : ")
	fmt.Scanln(&name)
	fmt.Println(name)

	var age int
	fmt.Print("Input Your Age : ")
	fmt.Scanln(&age)
	fmt.Println("Your current age is " + fmt.Sprint(age))

	var ageAfter5Years int = age + 5
	fmt.Println("Your age after 5 years is " + fmt.Sprint(ageAfter5Years))
}