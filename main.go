package main

import (
	"flag"
	tags "golang.org/x/net/html/atom"
	"golang.org/x/sync/syncmap"
)

var (
	VisitedLinks = syncmap.Map{}
	Domain       = flag.String("u", "", "Domain")
	Scrapable    = map[tags.Atom]bool{
		tags.A:      true,
		tags.Img:    true,
		tags.Script: true,
		tags.Link:   true,
	}
)

func main() {
	flag.Parse()
	firstLink := Link{*Domain, true}

	links := make(chan Link)
	pages := make(chan Page)

	go Crawl(firstLink, links, pages)

	n := 1
	for n > 0 {
		select {
		case link := <-links:
			if link.shouldCrawl() {
				go Crawl(link, links, pages)
				n++
			}
		case page := <-pages:
			go page.Report()
			n--
		}
	}
}
