# A Webcrawler written in Golang

Crawls a single domain, printing out a list of assets and links for each new page that it finds. Will include external links in the print out, but will not crawl them.

## Crawl something

Clone the repo into `$GOPATH/src` and do the following:

get deps:
```bash
  # inside the repo
  $ go get 
```
build:
```bash
  # inside the repo
  $ go build
```

crawl:
```bash
  # inside the repo
  $ ./web-crawler -u <url>
```
Make sure to include the protocol in your url, e.g. `http://`

example:
```bash
  # inside the repo
  $ ./web-crawler -u http://tomblomfield.com
```

## Test

Clone the repo and do the following:

get deps:
```bash
  # inside the repo
  $ go get 
```

run the tests
```bash
  # inside the repo
  $ go test 
```
A test server will be automatically spun up and torn down for the tests.
