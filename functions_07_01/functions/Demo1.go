package functions

import "fmt"

func Topla(sayi1 int, sayi2 int) int { //    { dan önceki int,int döndüreceğimiz anlamına gelir
	var toplam = sayi1 + sayi2
	return toplam
}

func Selamver(isim1 string, isim2 string) {
	fmt.Println("Merhaba", isim1)
	fmt.Println("Merhaba", isim2)
}
