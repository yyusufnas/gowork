package main

import (
	"fmt"
	"golesson/pointer"
)

func main() {
	sayi := 20
	pointer.Demo1(&sayi)
	fmt.Println("Maindeki sayi:", sayi)
}
