package sart

import "fmt"

func Demo2() {
	var para int = 1000
	var cek int = 0

	if cek > para {
		fmt.Println("Cekilmek istenen para bakiyeden fazla")

	} else if cek == 0 {
		fmt.Println("0 tl cekilemez")
	} else {
		fmt.Println("Paranız hazırlanıyor")
	}

}
