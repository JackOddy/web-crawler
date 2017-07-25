package main

import (
	"fmt"
	tags "golang.org/x/net/html/atom"
	"os"
)

var MaxDepth = 5
var VisitedLinks = make(map[Link]bool)

var Scrapable = map[tags.Atom]bool{
	tags.A:      true,
	tags.Img:    true,
	tags.Script: true,
	tags.Link:   true,
}

type HttpError struct {
	original string
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("missing URL argument")
	}

	go Crawl(os.Args[1])

	var input string
	fmt.Scanln(&input)
}
