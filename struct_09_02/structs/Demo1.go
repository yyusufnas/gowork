package structs

import "fmt"

type product struct {
	name      string
	unitPrice float64
	brand     string
}

func Demo1() {
	fmt.Println(product{"Bilgisayar", 5.2, "excalibur"})

}
