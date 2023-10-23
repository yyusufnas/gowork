package loops

import "fmt"

func Demo4() {
	sayi := 0
	fmt.Println("Bir sayi giriniz")
	fmt.Scanln(&sayi)
	asalmi := true
	for i := 2; i < sayi; i++ {

		if sayi%i == 0 {
			asalmi = false
		}
	}

	if asalmi == true {
		fmt.Printf("%d asaldir", sayi)

	} else {
		fmt.Printf("%d asal değildir", sayi)
	}

}
