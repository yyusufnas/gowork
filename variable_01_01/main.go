package main

import "fmt"

func main() {
	var metin string = "Selam Yusuf"
	fmt.Println(metin) //ln alt satıra geçirir

	var sayı float32 = 10.2
	fmt.Print(sayı)

	sayi2 := 50 // :=kullanımıyla da degisken olusturabiliriz.
	//Türünü(int,string vs)belirtmeye gerek yok

	fmt.Printf("VERİ TİPİ=%T", sayi2)
	//f ifadesi type belirlerken işimize yarar örn burada sayı2 yi içeri atar ve %T si nedir diye ekrana bastırır.
	//%d deseydik sayı 2 yi bastırırdı
}
