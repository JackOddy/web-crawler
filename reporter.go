package main

import (
	"fmt"
)

func Error(err error) {
	fmt.Println("ERROR: ", err)
}

func Warning(info string) {
	fmt.Println("WARNING: ", info)
}

func Page(url string, assets, links []Link) {
	fmt.Printf("\n\n [%s] \n\tASSETS:\n\t%v\n\tLINKS:\n\t%v", url, assets, links)
}
