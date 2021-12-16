package main

import (
	"fmt"
	"net/http"
	"time"
)

type urlList []string

func main() {
	// links := urlList{
	// 	"http://google.com",
	// 	"http://facebook.com",
	// 	"http://amazon.ca",
	// 	"http://stackoverflow.com",
	// 	"http://golang.org",
	// }

	// c := make(chan string)

	// links.checkStatus(c)

	c := make(chan string)
	c <- "Hi there!"
}

func (l urlList) checkStatus(c chan string) {
	for _, link := range l {
		go checkLink(link, c)
	}

	for link := range c {
		// time.Sleep(5 * time.Second)
		// go checkLink(link, c)

		go func(nl string) {
			time.Sleep(5 * time.Second)
			checkLink(nl, c)
		}(link)
	}

}

func checkLink(link string, c chan string) {
	// time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
