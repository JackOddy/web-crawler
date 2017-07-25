package main

import (
	"fmt"
	"os"
	"strings"
)

type Link struct {
	url string
}

func (self Link) String() string {
	return fmt.Sprintf("\t[%.50s]\n\t", self.url)
}

func (self Link) Crawlable() (crawlable bool) {
	if strings.Contains(strings.ToLower(self.url), os.Args[1]) {
		crawlable = true
	}
	return
}

func (self Link) Valid() (result bool) {
	if len(self.url) == 0 {
		return false
	}
	return true
}
