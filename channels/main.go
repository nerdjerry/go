package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!!")
		c <- "It seems not working"
		return
	}
	fmt.Println(link, "is working correct")
	c <- "Yeah it is working"
}
