package main

import "fmt"

func main() {
	var durum bool = false
	var metin1 string = "YUSUF"
	var metin2 string = "YUSUF"

	durum = metin1 == metin2
	fmt.Print(durum)
	//olusturduugmuz metın1 ve metın2 strınglerını eşit mi diye kontrol edip boolean turundekı duruma atıyoruz
	//eğer eşitse true değilse false değerini döndürüyor
	// metin1 !=metin2 ==metin1 metin2 den farklı mı anlamına gelir
}
