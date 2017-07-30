package main

import "fmt"

type Page struct {
	url           string
	assets, links []Link
}

func (self Page) Report() {
	fmt.Printf("\n\n[%s] \n\tASSETS:\n\t%v\n\tLINKS:\n\t%v",
		self.url,
		self.assets,
		self.links,
	)
}
