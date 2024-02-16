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

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	ip, err := http.Get(link)
	if err != nil {
		fmt.Printf("Error occurred on link: %v", link)
		return
	}
	fmt.Printf("%v is up with statusCode %d\n", link, ip.StatusCode)
}
