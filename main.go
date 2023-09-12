package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gocolly/colly"
)

type Book struct {
	Title string
	Price string
}

func bookScraper() {

	fmt.Println("Start scraping...")

	file, createErr := os.Create("export.csv")

	if createErr != nil {
		log.Fatal(createErr)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	headers := []string{"Title", "Price"}

	writer.Write(headers)

	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
	)

	c.OnRequest(func(r *colly.Request){
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response", r.StatusCode)
	})

	c.OnHTML(".product_pod", func(e *colly.HTMLElement){
		book := Book{}
		book.Title = e.ChildAttr(".image_container img", "alt")
		book.Price = e.ChildText(".price_color")
		row := []string{book.Title, book.Price}
		writer.Write(row)		
	})

	c.OnHTML(".next > a", func(h *colly.HTMLElement) {
		nextPage := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(nextPage)
	})

	c.Visit("https://books.toscrape.com/")

}

func main() {	
	bookScheduler := gocron.NewScheduler(time.UTC)
	bookScheduler.Every(2).Minutes().Do(bookScraper)
	bookScheduler.StartBlocking()
}