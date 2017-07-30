package main

import (
	"fmt"
	"strings"
)

type Link struct {
	url    string
	isLink bool
}

func (self Link) String() string {
	return fmt.Sprintf("\t[%.50s]\n\t", self.url)
}

func (self Link) shouldCrawl() bool {
	if strings.Contains(strings.ToLower(self.url), *Domain) {
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
