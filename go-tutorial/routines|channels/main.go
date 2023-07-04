package main

import (
	"log"
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
	log.Printf("Starting")
	c := make(chan string)
	count := 0
	for _, link := range links {
		count++
		go checkLink(link, c)
	}
	for i := 0; i < count; i++ {
		log.Println("i is", i)
		doSomething(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		//log.Println(link, "is down")
		//ac.Log("Bad Things")
		c <- link + " is down"
		return
	}
	c <- link + " is up"
	//log.Printf("%v is up", link)
}

func doSomething(s string) {
	log.Println("from do something", s)
}
