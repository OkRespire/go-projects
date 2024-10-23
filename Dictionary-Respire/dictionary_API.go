package main

import (
	"github.com/gocolly/colly"
)

var (
	c             colly.Collector
	parsingResult wordMeaning
)

func init() {
	c := colly.NewCollector(colly.AllowedDomains("dictionary.cambridge.org"))
	c.AllowURLRevisit = true

	settingSearch()
}

func settingSearch() {
	c.OnHTML()
}
