package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Lesson on slices")

	//declearing the size of slice is not important like array
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// inorder to add items to the list we append
	fruitList = append(fruitList, "Orange") //this will add item at end
	fmt.Print(fruitList)

	//start from position 1 and end at position 3
	fruitList = append(fruitList[1:3])
	fmt.Println("\nAfter removing first two elements : ", fruitList)

	highScore := make([]int, 4)
	highScore[0] = 256
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	//highScore[4] = 867 will not work

	// will print only the values avaible
	fmt.Println(highScore)

	highScore = append(highScore, 453, 343, 333)

	fmt.Println(highScore)

	sort.Ints(highScore)

	fmt.Println("Sorted high score : ", highScore)
    
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from slice based on index

	var courses = []string{"react", "javascript", "Java", "python"}

	fmt.Println(courses)
   
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	
	fmt.Println(courses)

	//.cap() checks capacity while .len()checks the capacity 

}
