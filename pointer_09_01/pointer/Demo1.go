package pointer

import "fmt"

func Demo1(sayi *int) {

	*sayi = *sayi + 1

	fmt.Println("Demodaki sayi:", *sayi)

}
