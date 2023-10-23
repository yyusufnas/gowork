package arrays

import "fmt"

func Demo3() {
	sayilar := [5]int{10, 20, 30, 40, 50} //5 eleman tanımladım ve değerlerini verdim

	fmt.Println(sayilar)
	for i := 0; i < len(sayilar); i++ { //len ile sayilar dizisinin eleman sayısını alırız
		fmt.Println(sayilar[i])

	}
}
