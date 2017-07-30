package main

import (
	"os"
	"testing"
	. "web-crawler/testServer"
)

func TestWebCrawler(t *testing.T) {
	os.Args = []string{"cmd", "-u=http://localhost:3000"}
	defer TestServer(t)()

	var testPage Page
	Reporter = func(p Page) {
		testPage = p
	}

	var presentLinks = []string{
		"http://localhost:3000/page-1.html",
		"http://localhost:3000/page-2.html",
		"http://localhost:3000/profile.html",
	}
	var presentAssets = []string{
		"http://localhost:3000/styles.css",
	}

	main()

	defer func() {

		for _, url := range presentLinks {
			found := false
			for _, foundLink := range testPage.links {
				if url == foundLink.url {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Could not find %s in testPage scraped links", url)
			}
		}

		for _, url := range presentAssets {
			found := false
			for _, foundLink := range testPage.assets {
				if url == foundLink.url {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Could not find %s in testPage scraped assets", url)
			}
		}
	}()
}
