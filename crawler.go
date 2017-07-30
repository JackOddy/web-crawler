package main

import (
	"net/http"
)

func Crawl(link Link, links chan Link, pages chan Page, errors chan bool) {
	page, ok := fetch(link.url)
	if !ok {
		errors <- true
		return
	}

	go Scrape(link.url, page, links, pages)
}

func fetch(url string) (resp *http.Response, ok bool) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode > 299 {
		return resp, false
	}
	return resp, true
}
