package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)

	sozluk["table"] = "masa"
	sozluk["book"] = "kitap"
	sozluk["car"] = "araba"

	fmt.Println(sozluk)

	delete(sozluk, "book") //sozlukten book degerı silinir

	fmt.Println(sozluk)

	_, kontrol := sozluk["table"] //table değeri map sozluk içinde var mı yok mu diye kontrol edilir

	fmt.Println(kontrol)
}
