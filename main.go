package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML(
		"pre.stats span.stats-both, pre.stats span.stats-firstonly",
		func(e *colly.HTMLElement) {
			fmt.Println(e.Text)
		},
	)

    err := c.Visit("https://adventofcode.com/2023/stats")
    if err != nil {
        log.Fatal(err)
    }
}
