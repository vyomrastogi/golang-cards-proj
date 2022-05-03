package main

import (
	"fmt"
	"net/http"
	"time"
)

// Build a website status checker

func main() {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}
	startTime := time.Now().UnixMilli()

	c := make(chan string) //channel which can communicate with string

	for _, website := range websites {
		go checkWebsite(website, c)
		//with child go routines, main routine does not wait for child to finish execution
		//channel helps in communication between routines
	}
	fmt.Println("Total execution time [", time.Now().UnixMilli()-startTime, " ms]")
}

func checkWebsite(website string, c chan string) {
	responseStart := time.Now().UnixMilli()
	_, err := http.Get(website)

	if err != nil {
		fmt.Println(website, "might be down!")
		fmt.Println("Response time [", time.Now().UnixMilli()-responseStart, " ms]")
		return //force stop of function execution
	}
	fmt.Println(website, "is up!")
	fmt.Println("Response time [", time.Now().UnixMilli()-responseStart, " ms]")
}

//TIL
//Each go process is executed as a go routine
//by default go attempts to use only one core
//go scheduler runs only one routine until makes a blocking call

//Java corollary
//Go Routine is similar to Thread
//Child Routine is similar to Child Thread
//Channel is kind of similar Futures
