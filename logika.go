package main

import "fmt"

func logika() {

	// Latihan 1
	// var n int = 6

	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("")
	// }

	// Latihan 2
	// var n int = 10

	// if n%3 == 0 && n%5 == 0 {
	// 	fmt.Println("FizzBuzz")
	// }else if n%3 == 0{
	// 	fmt.Println("Fizz")
	// }else if n%5 == 0 {
	// 	fmt.Println("Buzz")
	// }

	// Latihan 3
	// var n int = 5
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < (i+1); j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("")
	// }

	// Latihan 4
	var celcius float64
	fmt.Println("========================")
	fmt.Println("Aplikasi Konversi Suhu")
	fmt.Println("========================")
	fmt.Println("Input Konversi Suhu : ")
	fmt.Scanln(&celcius)

	fmt.Println("1. Fahrenhit")
	fmt.Println("2. Reamur")
	fmt.Println("3. Kelvin")

	var userContiniue string = "y"
	for userContiniue == "y" {
		
		var choice int
		fmt.Println("Input Nomor Suhu Konversi : ")
		fmt.Scanln(&choice)

		if choice == 1{
			result := (celcius * (9.0 / 5.0) + 32.0)
			fmt.Println("Konversi Suhu Fahrenhit : "+ fmt.Sprint(result))
		}else if choice == 2{
			result := celcius * (4.0 / 5.0)
			fmt.Println("Konversi Suhu Reamur : " + fmt.Sprint(result))
		}else if choice == 3{
			result := celcius + 273.15
			fmt.Println("Konversi Suhu Kelvin : " + fmt.Sprint(result))
		}
		fmt.Print("")
		fmt.Print("Apakah Anda Ingin Lanjut ? (y/n)")
		fmt.Scanln(&userContiniue)
	}
}
