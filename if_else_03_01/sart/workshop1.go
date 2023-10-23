package sart

import "fmt"

func Demo3() {
	var sayi1, sayi2, sayi3 int = 5, 10, 15
	if sayi1 > sayi2 && sayi1 > sayi3 {
		fmt.Printf("En Büyük sayi %d", sayi1)

	} else if sayi2 > sayi1 && sayi2 > sayi3 {
		fmt.Printf("En büyük sayi %d", sayi2)
	} else {
		fmt.Printf("En büyük sayi %d", sayi3)
	}

}
