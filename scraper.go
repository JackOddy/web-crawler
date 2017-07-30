package main

import (
	"bytes"
	"golang.org/x/net/html"
	tags "golang.org/x/net/html/atom"
	"net/http"
	"strings"
)

func createLink(tag html.Token) Link {
	var url string
	for _, attr := range tag.Attr {
		if attr.Key == "href" || attr.Key == "src" {
			url = formatUrl(attr.Val)
			break
		}
	}

	if tag.DataAtom == tags.A {
		return Link{url, true}
	}
	return Link{url, false}
}

func formatUrl(url string) string {
	var buffer bytes.Buffer

	url = strings.TrimSpace(url)

	if url != "" && string(url[0]) == "/" {
		buffer.WriteString(*Domain)
		buffer.WriteString(url)
		return buffer.String()
	}

	return url
}

func verifyTag(token html.Token) (link Link, ok bool) {
	link = createLink(token)
	if link.Valid() {
		ok = true
	}
	return
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
