package rangee

import "fmt"

func Demo2() {
	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //for range ile 1 den 10 a kadar tek sayıları yazdırma
	i := 1
	for _, sayi := range sayilar {

		if i == 1 {
			fmt.Println(sayi)
			i = 0
		} else {
			i = 1

		}

	}
}
