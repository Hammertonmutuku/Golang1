package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")

	hitesh := User{"Hammerton", "hammertonmutuku@gmail.com", true, 24 }

	// fmt.Println(hitesh)

	// fmt.Printf("hitesh details are: %+v\n", hitesh)
	
	// fmt.Printf("Name is:  %v, and emil is: %v ", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
   fmt.Println("Get user status: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}