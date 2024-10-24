package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	// "net/http"
	// "log"
)

func main() {
	// initialises a new collector for visited url

	// visitedURLs := map[string]struct{}{}

	var scrapedDomain string

	_, err := fmt.Scanln(&scrapedDomain)
	if err != nil {
		log.Fatal(err)
	}

	// initialises a new collector
	c := colly.NewCollector(colly.AllowedDomains(scrapedDomain))

	// sets the user agent
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	// when it reaches a hyperlink, it visits it
	go c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Found link:", link)
		if "https://"+e.Request.URL.Hostname() != scrapedDomain {
			log.Fatal("Outside of domain")
		}

		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnError(func(e *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	go c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit(scrapedDomain)
}
