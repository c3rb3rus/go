package main

import (
	"fmt"

	"github.com/c3rb3rus/go/pointers/blaat"
)

func main() {
	myString := "Some text"

	// Pass by value
	fmt.Println(changeString(myString))
	fmt.Println(myString)

	// Pass by reference
	changeStringByRef(&myString)
	fmt.Println(myString)
	blaat.Blaat(&myString)
	fmt.Println(myString)
}

func changeString(myString string) string {
	myString += " changed"
	return myString
}

func changeStringByRef(myString *string) {
	*myString += " changed"
}
