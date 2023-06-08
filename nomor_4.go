package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	
	// Memeriksa apakah bilangan memiliki faktor selain 1 dan dirinya sendiri
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	
	return true
}

func main() {
	var num int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&num)
	
	if isPrime(num) {
		fmt.Println("Angka", num, "adalah bilangan prima.")
	} else {
		fmt.Println("Angka", num, "bukan bilangan prima.")
	}
}