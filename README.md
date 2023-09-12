# Basic web scraper using Golang(Go)

We build a very basic web scraper in Go to scrap data from different website based on CSS selector.

## Package used

- github.com/gocolly/colly
- github.com/go-co-op/gocron

We build a web scraper that scraps books title and price from `books.toscrape.com`. It navigate to all pages and extract the relevent information using css selectors. We also used a scheduler which is set to run the scraper at a specified time interval.
