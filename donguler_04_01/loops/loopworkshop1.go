package loops

import "fmt"

func Demo3() {
	aklimdakisayi := 80
	tahmin := 0
	kac := 1
	fmt.Println("Bir sayi tahmin ediniz")
	fmt.Scanln(&tahmin)
	for tahmin != aklimdakisayi {

		if tahmin < aklimdakisayi {
			fmt.Println("Daha buyuk bır sayı gırınız")
			fmt.Scanln(&tahmin)
		}
		if tahmin > aklimdakisayi {
			fmt.Println("Daha kucuk bır sayı gırınız")
			fmt.Scanln(&tahmin)
		}
		if tahmin == aklimdakisayi {
			fmt.Println("dogru bildiniz")

		}
		kac = kac + 1
	}

	if kac >= 1 && kac <= 3 {
		fmt.Printf("Bildiniz.tahmin sayısı %d :SÜPER", kac)
	}
	if kac > 3 && kac <= 6 {
		fmt.Printf("Bildiniz.tahmin sayısı %d :Güzel", kac)
	}
	if kac > 6 {
		fmt.Printf("Bildiniz.tahmin sayısı %d :Fena değil", kac)
	}

}
