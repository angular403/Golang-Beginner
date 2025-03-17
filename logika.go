package main

import "fmt"

func logika() {
	var n int = 6

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}