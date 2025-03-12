package main

import "fmt"

func statement() {
	// var number int

	// fmt.Print("Masukkan Angkamu : ")
	// fmt.Scanln(&number)
	// fmt.Println("Angkamu Adalah : " + fmt.Sprint(number))

	// if number % 2 == 0 {
	// 	fmt.Println("Bilanganmu adalah bilangan Genap")
	// }else{
	// 	fmt.Println("Bilanganmu adalah Bilangan Ganjil")
	// }

	// var score int

	// fmt.Print("Masukkan Angkamu : ")
	// fmt.Scanln(&score)

	// if score == 90{
	// 	fmt.Println("Nilai Siswamu Adalah A")
	// }else if score >= 75{
	// 	fmt.Println("Nilai Siswamu Adalah B")
	// }else if score >= 60{
	// 	fmt.Println("Nilai Siswamu Adalah C")
	// }else if score >= 50{
	// 	fmt.Println("Nilai Siswamu Adalah D")
	// } else{
	// 	fmt.Println("Maaf Siswamu Belum Lulus")
	// }

	var x int
	fmt.Print("Masukkan Nilai Siswamu : ")
	fmt.Scanln(&x)


	if 100 >= x && x >= 90{
		fmt.Println("Nilai Siswamu adalah A")
	}else if 90 >= x && x >= 80{
		fmt.Println("Nilai Siswamu adalah B")
	}else if 80 >= x && x >= 70 {
		fmt.Println("Nilai Siswamu adalah C")
	}else if 70 >= x && x >= 60{
		fmt.Println("Nilai Siswamu adalah D")
	}else if 60 >= x && x >= 0{
		fmt.Println("Nilai Siswamu adalah E")
	}else{
		fmt.Println("Nilai Tidak Valid")
	}
}