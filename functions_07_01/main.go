package main

import (
	"fmt"
	"golesson/functions"
)

func main() {

	functions.Selamver("yusuf", "mert")

	sonuc := functions.Topla(1, 2)
	fmt.Println(sonuc)

}
