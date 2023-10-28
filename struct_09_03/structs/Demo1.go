package structs

import "fmt"

type customers struct {
	Firstname string
	Lastname  string
	age       int
}

func (c customers) save() {
	fmt.Println("Eklendi", c.Firstname)
}
func (c customers) update() {
	fmt.Println("Güncellendi", c.Firstname)
}
func Demo1() {
	c := customers{Firstname: "Yusuf", Lastname: "Nas", age: 20}
	c.save()
	c.update()

}
