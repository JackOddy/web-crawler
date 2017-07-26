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

func (self Link) Crawlable() bool {
	if strings.Contains(strings.ToLower(self.url), os.Args[1]) {
		return true
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
