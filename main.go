package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	channel := make(chan string)

	for _, link := range links {
		// By Placing go keyword we create a new Go Routine
		go checkLink(link, channel)
	}

	for {
		//fmt.Println(<-channel)
		checkLink(<-channel, channel)
	}

}

func checkLink(link string, channel chan string) {
	ip, err := http.Get(link) // This is Blocking Call!! So Main Go Routine is suspended
	if err != nil {
		fmt.Printf("Error occurred on link: %v", link)
		//channel <- "Error occurred on link: " + link.
		channel <- link
		return
	}
	fmt.Printf("%v is up with statusCode %d\n", link, ip.StatusCode)
	//channel <- "It's up"
	channel <- link
}
