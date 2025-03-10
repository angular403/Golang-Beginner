package main

import "fmt"

func aritmatika() {
	var tambah float64 = 5 + 10
	var kurang  float64= 20 - 5
	kali := 5 * 5
	bagi := 10 / 2
	var hasilBagi float64 = 10.0 / 3.0
	var complexResult float64 = 1.0/3.0*5.0 + 100.0

	var complexResultAgain = tambah * kurang + (complexResult) * 10.0

	var modolus = 10 % 6
	var modolus1 = 10 % 7

	fmt.Println(tambah)
	fmt.Println(kurang)
	fmt.Println(kali)
	fmt.Println(bagi)
	fmt.Println(hasilBagi)
	fmt.Println(complexResult)
	fmt.Println(complexResultAgain)
	fmt.Println(modolus)
	fmt.Println(modolus1)


	var name string = "Andrew" + " " + "Wiliam "
	var lastName string = "Napitupulu"
	var fullName = name + lastName

	// fmt.Println(name)
	// fmt.Println(lastName)
	fmt.Println(fullName)
}
