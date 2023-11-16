package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Wecome to time study of golang")
	presentTIme := time.Now()
	fmt.Println(presentTIme)
	fmt.Println(presentTIme.Format("01-02-2006 Monday 15:04:05"))
	createdDate := time.Date(2023, time.April, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
