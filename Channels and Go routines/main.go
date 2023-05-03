package main

import (
	"fmt"
	"net/http"
	"time"
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
		go checkLink(link, channel)
	}

	// for i := 0; i < len(links); i++ {
	// 	go checkLink(<-channel)
	// }

	// for {
	// 	go checkLink(<-channel, channel)
	// }

	for link := range channel {
		go func(linkArg string) {
			time.Sleep(5 * time.Second)
			checkLink(linkArg, channel)
		}(link)
	}

}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		channel <- link
		return
	} else {
		fmt.Println(link, "is up!")
		channel <- link
	}
}
