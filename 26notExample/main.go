package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Race condition - learnCodeonline")
   
    
	wg := &sync.WaitGroup{}

	var score =[]int{0}

	wg.Add(3)

	go func(wr *sync.WaitGroup){
		fmt.Println("One R")
		score = append(score, 1)
		wg.Done()
	}(wg)
	go func(wr *sync.WaitGroup){
		fmt.Println("Two R")
		score = append(score, 2)
		wg.Done()
	}(wg)
	
	go func(wr *sync.WaitGroup){
		fmt.Println("Two R")
		score = append(score, 3)
		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println(score)

}
