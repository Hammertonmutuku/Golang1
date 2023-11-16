package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}


var wg sync.WaitGroup // pointers

var mut sync.Mutex // pointer 

func main() {
	//    go greeter("Hello")
	//    greeter("world")

	websitelist := []string{
		"https://google.com",
		"https://instagram.com",
		"https://facebook.com",
		"https://twitter.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)

	}
	wg.Wait()
	fmt.Println(signals)

}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		
		// log.Fatal(err)
		fmt.Println("Oops in endpoint")
	}

    mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}
