package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")

	hitesh := User{"Hammerton", "hammertonmutuku@gmail.com", true, 24, wheel{23, "fhfh"}, wheel{34, "fhfhf"},car{"jfjf", "dhhdhdhd"}}

	fmt.Println(hitesh)

	fmt.Printf("hitesh details are: %+v\n", hitesh)
	
	fmt.Printf("Name is:  %v, and emil is: %v ", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	FrontWheel wheel
	Backwheel wheel
	car
}

type wheel struct {
	Radius int
	Material string
}

type car struct {
	make string
	model string
}