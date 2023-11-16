package main

import "fmt"

func main() {
  greatings()

  results := adder(3, 4)

  fmt.Println("results is: ", results)

  r := proAdder(3,34,3,3)

  fmt.Println(r)
}

func adder(vOne int, vTwo int) int {
    return vOne + vTwo
}
func greatings() {
	fmt.Println("Namstay from the other side  ")
}

func proAdder(values ...int) int {
	total :=0

	for _, val := range values {
		total +=val
	}

	return total
}
