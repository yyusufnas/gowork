package rangee

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"}
	for i, sehir := range sehirler { //sehirleri tek tek dolasır ve sehir degiskenine gönderir
		fmt.Println(i)
		fmt.Println(sehir)
	}
}
