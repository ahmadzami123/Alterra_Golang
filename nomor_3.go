package main

import "fmt"

func main() {
	var bilangan int

    // Masukkan bilangan yang ingin dicari faktornya
    fmt.Print("Masukkan bilangan: ")
    fmt.Scanln(&bilangan)


    // Loop untuk mencari faktor
    for i := 1; i <= bilangan; i++ {
        if bilangan%i == 0 {
            fmt.Println(i)
        }
    }
}