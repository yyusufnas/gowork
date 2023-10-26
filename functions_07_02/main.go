package main

import (
	"fmt"
	"golesson/functions"
)

func main() {
	sonuc1, sonuc2, sonuc3, sonuc4 := functions.Islem(10, 6)
	fmt.Println("Toplam:", sonuc1)
	fmt.Println("Cıkarım:", sonuc2)
	fmt.Println("Carpım:", sonuc3)
	fmt.Println("Bolum:", sonuc4)

}
