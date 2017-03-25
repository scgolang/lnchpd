package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/scgolang/launchpad"
)

func main() {
	var (
		reset bool
	)
	flag.BoolVar(&reset, "reset", false, "reset all the buttons on the launchpad")
	flag.Parse()

	lp, err := launchpad.Open()
	if err != nil {
		log.Fatal(err)
	}
	if reset {
		_ = lp.Reset()
		return
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
