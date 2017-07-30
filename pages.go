package main

import "fmt"

var Reporter = func(p Page) {
	fmt.Printf("\n\n[%s] \n\tASSETS:  \n\t%v  \n\tLINKS:  \n\t%v",
		p.url,
		p.assets,
		p.links,
	)
}

type Page struct {
	url           string
	assets, links []Link
}

func (self Page) Report() {
	Reporter(self)
}
