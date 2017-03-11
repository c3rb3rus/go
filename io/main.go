package main

import "fmt"

import "io/ioutil"
import "os"
import "log"

func main() {
	f, err := os.Open("readme.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", string(b))
}
