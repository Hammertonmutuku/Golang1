package main

import "fmt"

func main() {
	fmt.Println("Lesson on Arrays")
    var fruitList [4] string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Tomato"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println(vegList)
}
