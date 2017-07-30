package main

import (
	"golang.org/x/net/html"
	"net/http"
)

func Scrape(url string, resp *http.Response, found chan Link, pages chan Page) {
	var assets, links []Link

	page := html.NewTokenizer(resp.Body)

	for {
		_ = page.Next()
		token := page.Token()
		var link Link
		var ok bool

		if token.Type == html.ErrorToken {
			go func() { pages <- Page{url, assets, links} }()
			break
		}

		if isOpeningOrSelfClosing(token) && Scrapable[token.DataAtom] {
			link, ok = verifyTag(token)
		}

		if ok {
			if link.isLink {
				go func() { found <- link }()
				links = append(links, link)
			} else {
				assets = append(assets, link)
			}
		}
	}
}

func isOpeningOrSelfClosing(t html.Token) bool {
	switch t.Type {
	case html.StartTagToken:
		return true
	case html.SelfClosingTagToken:
		return true
	default:
	}
	return false
}

func verifyTag(token html.Token) (link Link, ok bool) {
	link.Extract(token)
	if link.Valid() {
		ok = true
	}
	return
}
