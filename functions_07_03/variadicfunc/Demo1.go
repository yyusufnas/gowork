package variadicfunc

func Islem(sayilar ...int) int { //variadic fonksiyonlar ile parametre sayısı belirtmeden dizi yapısı gibi kullanabilirsiniz
	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam = toplam + sayilar[i]
	}
	return toplam
}
