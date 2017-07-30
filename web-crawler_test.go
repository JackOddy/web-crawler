package main

import (
	"os"
	"testing"
	. "time"
	. "web-crawler/testServer"
)

func TestWebCrawler(t *testing.T) {
	os.Args = []string{"cmd", "-u=http://localhost:3000"}
	defer TestServer(t)()

	var testPage Page
	Reporter = func(p Page) {
		testPage = p
	}

	var expectedLinks = []string{
		"http://localhost:3000/page-1.html",
		"http://localhost:3000/page-2.html",
		"http://localhost:3000/profile.html",
		"http://twitter.com/test",
		"http://linkedin.com/test",
	}
	var expectedAssets = []string{
		"http://localhost:3000/styles.css",
		"http://localhost:3000/scripts.js",
		"http://localhost:3000/image.jpg",
	}

	main()

	shouldFind(t, expectedLinks, testPage.links, "links")
	shouldFind(t, expectedAssets, testPage.assets, "assets")
}

func shouldFind(t *testing.T, expectedSubjects []string, actualSubjects []Link, subject string) {
	for _, url := range expectedSubjects {
		found := false
		for _, foundLink := range actualSubjects {
			if url == foundLink.url {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Could not find %s in testPage scraped %s", url, subject)
		}
	}
}
