package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcome to file in golang")

	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)


	fmt.Println("Length is : ", length)

	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filname string) {
	databyte, err := os.ReadFile(filname)

	checkNilError(err)

	fmt.Println("Text data insile the file is \n", string(databyte))
}

func checkNilError (err error ) {
	if err != nil {
		panic(err)
	}
}