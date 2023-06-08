package main

import "fmt"

func main() {
	// LUAS PERMUKAAN TABUNG
	T := 20.0
	r := 4.0
	pi := 3.14
	L := 2 * pi * r * (r + T)
	fmt.Println(L)
}