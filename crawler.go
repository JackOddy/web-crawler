package main

import (
	"errors"
	"fmt"
	"net/http"
)

func Crawl(link Link, links chan Link, pages chan Page, errors chan bool) {
	page, err := fetch(link.url)

	if err != nil {
		errors <- true
		return
	}

	go Scrape(link.url, page, links, pages)

}

func fetch(url string) (resp *http.Response, err error) {
	resp, err = http.Get(url)
	if resp.StatusCode > 299 {
		msg := fmt.Sprintf("response code is %d", resp.StatusCode)
		err = errors.New(msg)
	}
	return
}
