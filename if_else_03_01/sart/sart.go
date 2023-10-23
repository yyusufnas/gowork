package sart

import "fmt"

func Demo1() {
	var para int = 1000
	var cek int = 400
	if cek > para {
		fmt.Println("Cekilmek istenen para bakiyeden fazla")

	}
	if cek <= para {
		fmt.Println("Paranız hazırlanıyor..")
		para = para - cek
		fmt.Printf("Güncel paranız=%d", para)
	}

}
