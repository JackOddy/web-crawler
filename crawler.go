package main

import (
	"fmt"
	"net/http"
)

func fetch(url string) (resp *http.Response, err error) {
	resp, err = http.Get(url)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	if resp.StatusCode > 299 {
		fmt.Println("ERROR: response code is ", resp.StatusCode)
		return
	}
	return
}

func Crawl(url string) {
	page, err := fetch(url)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	links := ScrapePage(url, page)

	for _, link := range links {
		if link.Crawlable() && VisitedLinks[link] != true {
			VisitedLinks[link] = true
			go Crawl(link.url)
		}
	}
}
