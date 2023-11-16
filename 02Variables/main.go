package main

import "fmt"

//the capital L indicates that its  a public constant and can be accessed anywhere
const LoginToken string = "shdddhdhdhd"
func main() {
	var username string = "Hammerton"
	fmt.Println(username)
	fmt.Printf("Varivale is of type: %T \n", username)

	var isloggedin bool = true
	fmt.Println(isloggedin)
	fmt.Printf("Varivale is of type: %T \n", isloggedin)
    
	//just use variable type int if no specific number in mind
	//int8 accepts negative numbers 
	//unit8 does not accept negative numbers
	// the Higher the more numbers it can hold 
	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Varivale is of type: %T \n",smallval)
    
	//float32 only picks the first 5 digits after decimal point 
	//float64 is more accurate
	var smallfloat float32 = 255.122334344245
	fmt.Println(smallfloat)
	fmt.Printf("Varivale is of type: %T \n",smallfloat)

	var anotherVariable bool
	fmt.Println(anotherVariable)
	fmt.Printf("Varivale is of type: %T \n",anotherVariable)

	// implicit type 
	var website = "Hammerton"
	fmt.Println(website)

	//no vae style
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Varivale is of type: %T \n",LoginToken)

}
