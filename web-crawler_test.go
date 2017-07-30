package main

import (
	"os"
	"testing"
	. "web-crawler/testServer"
)

func TestWebCrawler(t *testing.T) {
	os.Args = []string{"cmd", "-u=http://localhost:3000"}
	defer TestServer(t)()
	main()
}
