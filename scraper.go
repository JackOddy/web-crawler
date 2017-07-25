package main

import (
	"fmt"
	"golang.org/x/net/html"
	tags "golang.org/x/net/html/atom"
	"net/http"
	"strings"
)

func createLink(tag html.Token) Link {
	var url string
	for _, attr := range tag.Attr {
		if attr.Key == "href" || attr.Key == "src" {
			url = strings.TrimSpace(attr.Val)
			break
		}
	}
	return Link{url}
}

func appendLink(links []Link, token html.Token) []Link {
	link := createLink(token)
	if link.Valid() {
		links = append(links, link)
	}
	return links
}

// reads links
func ScrapePage(url string, resp *http.Response) []Link {
	var assets, links []Link
	page := html.NewTokenizer(resp.Body)

	for {
		_ = page.Next()
		token := page.Token()

		if token.Type == html.ErrorToken {
			break
		}

		if token.Type == html.StartTagToken && Scrapable[token.DataAtom] {
			switch token.DataAtom {
			case tags.A:
				links = appendLink(links, token)
			default:
				assets = appendLink(assets, token)
			}
		}

	}

	go fmt.Printf("\n\n [%s] \n\tASSETS:\n\t%v\n\tLINKS:\n\t%v", url, assets, links)
	return links
}
