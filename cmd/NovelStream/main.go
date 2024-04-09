package main

import (
	"log"
	"github.com/Odery/NovelStream/internal/scraper"
)

func main() {
	log.Println("START")

	scraper.ScrapAsuraImages("https://asuratoon.com/8612194254-i-am-the-fated-villain-chapter-1/", "images")
}