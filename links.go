package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	tags "golang.org/x/net/html/atom"
	"strings"
)

type Link struct {
	url    string
	isLink bool
}

func (self Link) String() string {
	return fmt.Sprintf("\t[%.50s]\n\t", self.url)
}

func (self Link) ShouldCrawl(domain *string) bool {
	if strings.Contains(strings.ToLower(self.url), *domain) {
		if _, visited := VisitedLinks.LoadOrStore(self, true); !visited {
			return true
		}
	}
	return false
}

func (self Link) Valid() bool {
	if len(self.url) == 0 {
		return false
	}
	if self.url == "#" {
		return false
	}
	return true
}

func (self *Link) Extract(tag html.Token) *Link {
	for _, attr := range tag.Attr {
		if attr.Key == "href" || attr.Key == "src" {
			self.url = formatUrl(attr.Val)
			break
		}
	}

	if tag.DataAtom == tags.A {
		self.isLink = true
	}
	return self
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
