package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3) //3 elemanlı bir dizi oluşturduk
	isimler[0] = "jose"
	isimler[1] = "jose1"
	isimler[2] = "jose2"

	isimler = append(isimler, "nas") //diziye sonradan ekleme yapmak için append yapısını kullanırız
	fmt.Println(isimler[1:4])        //1. elemandan baslar 4.elemana kadar yazdırır,4 dahil değil
	fmt.Println(isimler[1:])         //1. elemandan baslar son elemana kadar yazdırır
	fmt.Println(isimler[:4])         //en baştan başlar 4 e kadar yazdırır

}
