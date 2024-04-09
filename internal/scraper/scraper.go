package scraper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/Odery/NovelStream/internal/types"
)

//Scraps page of the popular Asura scans website
//Only for image scrapping. For info scraping about
//a Webnovel use ScrapAsuraDetails()
func ScrapAsuraImages(url string, path string){
	collector := newCollector()

	//Used for naming downloaded images (from 0 upwards)
	counter := 0
	collector.OnHTML("#readerarea img", func(e *colly.HTMLElement) {
		//Create the full path for the new file
		fullPath := filepath.Join(path, fmt.Sprintf("%d%s", counter, filepath.Ext(e.Attr("src"))))
		getImages(e.Attr("src"), fullPath)
		counter ++
	})

	collector.Visit(url)

	collector.Wait()
}

//Scraps page for the Novel description and cover image
//path is for cover image location
func ScrapAsuraDetails(url string, path string) types.Novel{

	//Initialize an empty Novel type 
	novel := types.Novel{
		Genre: make([]string, 0),
	}

	collector := newCollector()

	//Get cover image
	collector.OnHTML("img.wp-post-image", func(e *colly.HTMLElement) {
		//Create the full path for the new file
		fullPath := filepath.Join(path, fmt.Sprintf("cover%s", filepath.Ext(e.Attr("src"))))
		getImages(e.Attr("src"), fullPath)
		novel.CoverImage = fullPath
	})

	//Get Title
	collector.OnHTML("h1.entry-title", func(e *colly.HTMLElement) {
		novel.Title = e.Text
	})

	//Get Author
	collector.OnHTML("<b>Artist</b>", func(e *colly.HTMLElement) {
		if e.Text == "Artist" {
			nextSibling := e.DOM.Next()

			if nextSibling.Is("span") {
				novel.Author = nextSibling.Text()
			}
		}
	})

	//Get Genre
	collector.OnHTML("img.wp-post-image", func(e *colly.HTMLElement) {
		novel.Genre = nil
	})

	//Get Summary
	collector.OnHTML("", func(e *colly.HTMLElement) {
		novel.Summary = e.Attr("?")
	})

	//Get publish date
	collector.OnHTML("", func(e *colly.HTMLElement) {
		novel.PublishDate = time.Time{}
	})

	//Get update date
	collector.OnHTML("", func(e *colly.HTMLElement) {
		novel.UpdatedDate = time.Time{}
	})

	//Get status
	collector.OnHTML("", func(e *colly.HTMLElement) {
		novel.Status = e.Attr("?")
	})


	collector.Visit(url)

	collector.Wait()

	return novel
}


//Downloads images from a web page
func getImages(url string, fullPath string) {
	//Create the file to store the image
	newFile, err := os.Create(fullPath)
	if err != nil {
		log.Printf("[ERROR]: Couldn't create a new file. %s", err)
	}
	defer newFile.Close()

	//Getting the image
	response, err := http.Get(url)
	if err != nil {
		log.Printf("[ERROR]: Couldn't get the image. %s\n", err)
	}
	defer response.Body.Close()

	//Writing recieved image to the file
	_, err = io.Copy(newFile, response.Body)
	if err != nil {
		log.Printf("[ERROR]: Couldn't write to file. %s", err)
	}
}

func newCollector() colly.Collector {
	collector := colly.NewCollector(
		//Turn on async mode
		colly.Async(),
		colly.IgnoreRobotsTxt(),
	)

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	collector.Limit(&colly.LimitRule{
		Parallelism: 2,
		Delay: 5 * time.Second,
	})

	return *collector
}
