package main

import (
	"fmt"
	"log"

	"github.com/scgolang/launchpad"
)

func main() {
	lp, err := launchpad.Open()
	if err != nil {
		log.Fatal(err)
	}
	hits, err := lp.Hits()
	if err != nil {
		log.Fatal(err)
	}
	for hit := range hits {
		if hit.Err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d %d\n", hit.X, hit.Y)
	}
}
