package arrays

import "fmt"

func Demo4() {

	var sayilar [2][3]int //2 satır 3 sutunlu bir dizi
	sayilar[0][0] = 1
	sayilar[0][1] = 3
	sayilar[0][2] = 5
	sayilar[1][0] = 2
	sayilar[1][1] = 4
	sayilar[1][2] = 6
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(sayilar[i][j])
		}
	}
}
