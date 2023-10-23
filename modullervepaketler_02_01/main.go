package main

import (
	"fmt"
	"goders/yusuf" //goders modülünün içindeki  tanımlamak istediğimiz paketi burada tanımlıyoruz
) //birden fazla paketi bu şekilde tanımlayabiliriz

func main() {
	yusuf.Demo2() //paket ismi.fonksiyon ismi ile istediğimiz fonksiyonu çalıştırabiliriz
	fmt.Println("nas")
}
