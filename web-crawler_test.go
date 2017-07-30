package main

import (
	"os"
	"testing"
	. "web-crawler/testServer"
)

func TestWebCrawler(t *testing.T) {
	var testPage Page
	Reporter = func(p Page) {
		testPage = p
	}

	os.Args = []string{"cmd", "-u=http://localhost:3000"}
	defer TestServer(t)()
	main()

	for _, link := range testPage.links {
		if _, ok := VisitedLinks.Load(link); !ok {
			t.Errorf("Could not find %s in VisitedLinks", link.url)
		}
		t.Logf("Successfully crawled and stored %s", link.url)
	}
	// VisitedLinks.Range(func(key interface{}, value interface{}) bool {
	// 	link := key.(Link)
	// })
}
