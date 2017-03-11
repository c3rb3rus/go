package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := make(chan string, 5)

	go func() {
		for i := 1; i <= 2; i++ {
			fmt.Println("send", i)
			c <- "ping " + strconv.Itoa(i)
			//time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}
	}()

	go func() {
		for {
			fmt.Println("receive", <-c)
			//time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}
	}()

	var input string
	fmt.Scan(&input)
}
