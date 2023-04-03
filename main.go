package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// c := colly.NewCollector(
	// 	colly.AllowedDomains("en.wikipedia.org"),
	// )
	// c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
	// 	links := e.ChildAttrs("a", "href")
	// 	fmt.Println(links)
	// })
	// c.Visit("https://en.wikipedia.org/wiki/Web_scraping")

	c := colly.NewCollector()
	c.OnHTML("table#customers", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			writer.Write([]string{
				el.ChildText("td:nth-child(1)"),
				el.ChildText("td:nth-child(2)"),
				el.ChildText("td:nth-child(3)"),
			})
		})
		fmt.Println("Scrapping Complete")
	})
	c.Visit("https://www.w3schools.com/html/html_tables.asp")
}
