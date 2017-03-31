package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

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
	go printHits(lp)

	scn := bufio.NewScanner(os.Stdin)
	for scn.Scan() {
		txt := scn.Text()

		var x, y uint8
		if _, err := fmt.Sscanf(txt, "%d %d", &x, &y); err != nil {
			log.Fatal(err)
		}
		if err := lp.Light(x, y, launchpad.Color{Green: launchpad.Full}); err != nil {
			log.Fatal(err)
		}
	}
	if err := scn.Err(); err != nil {
		log.Fatal(err)
	}
}

func printHits(lp *launchpad.Launchpad) {
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
