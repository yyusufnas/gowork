package loops

import "fmt"

func Demo5() {
	sayi1 := 0
	sayi2 := 0

	var toplam1, toplam2 int = 0, 0
	fmt.Println("İki sayi giriniz")
	fmt.Scanln(&sayi1)
	fmt.Scanln(&sayi2)
	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			toplam1 = toplam1 + i

		}
	}
	for j := 1; j < sayi2; j++ {

		if sayi2%j == 0 {
			toplam2 = toplam2 + j
		}
	}
	if toplam1 == sayi2 && toplam2 == sayi1 {
		fmt.Printf("%v ve %v arkadaş sayılardır", sayi1, sayi2)
	} else {
		fmt.Printf("%v ve %v arkadaş sayı değillerdir", sayi1, sayi2)
	}
}
