package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go testLink(link, c)
	}

	for link := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			testLink(link, c)
		}(link)
	}
}

func testLink(link string, c chan string) bool {
	_, err := http.Get(link)
	c <- link

	if err != nil {
		fmt.Println(link, "might be down")
		return false
	}
	fmt.Println(link, "is up!")
	return true
}
