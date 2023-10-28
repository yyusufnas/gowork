package main

import (
	"golesson/goroutines"
	"time"
)

func main() {

	go goroutines.CiftSayilar()
	go goroutines.TekSayilar() //go ile aynı anda calısmalarını saglarız
	time.Sleep(10 * time.Second)
}
