package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	flgStat := flag.String("s", "total", "Total participants in the year.")
	flgYear := flag.Int("y", 2023, "Total participants in the year.")
	flag.Parse()

	switch *flgStat {
	case "total":
		handleTotal(*flgYear)
	}
}

func handleTotal(year int) {
	c := colly.NewCollector()
	c.OnHTML(
		"pre.stats span.stats-both, pre.stats span.stats-firstonly",
		func(e *colly.HTMLElement) {
			fmt.Println(e.Text)
		},
	)

	url := fmt.Sprintf("https://adventofcode.com/%v/stats", year)
	fmt.Println("Visiting URL:", url)

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}
