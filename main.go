package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	flgStat := flag.String("s", "graph", "Total participants in the year.")
	flgYear := flag.Int("y", 2023, "Total participants in the year.")
	flag.Parse()

	switch *flgStat {
	case "graph":
		handleGraph(*flgYear)
	case "total":
		handleTotal(*flgYear)
	}
}

func handleGraph(year int) {
	c := colly.NewCollector()

	c.OnHTML(
		"a",
		func(e *colly.HTMLElement) {
			href := e.Attr("href")
			pattern := fmt.Sprintf("/%v/day/", year)

			if strings.HasPrefix(href, pattern) {
				firstOnlyStars := e.ChildText(".stats-firstonly:last-child")
				firstOnlyStarsPattern := strings.Replace(firstOnlyStars, "*", "\\*", -1)
				re := regexp.MustCompile(firstOnlyStarsPattern + "$")
				fmt.Println(re.ReplaceAllString(e.Text, "|"+firstOnlyStars))
			}
		},
	)

	url := fmt.Sprintf("https://adventofcode.com/%v/stats", year)
	fmt.Println("Visiting URL:", url)
	fmt.Println("")

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}

func handleTotal(year int) {
	total := 0
	c := colly.NewCollector()

	c.OnHTML(
		".stats-both:nth-child(1), .stats-firstonly:nth-child(2)",
		func(e *colly.HTMLElement) {
			totalInt, err := strconv.Atoi(strings.TrimSpace(e.Text))
			if err != nil {
				return
			}
			total += totalInt
		},
	)

	url := fmt.Sprintf("https://adventofcode.com/%v/stats", year)
	fmt.Println("Visiting URL:", url)
	fmt.Println("")

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total Submissions for %v:\n", year)
	fmt.Println(total)
}
