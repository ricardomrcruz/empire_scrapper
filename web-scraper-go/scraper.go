package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// initializing a data structure to keep the scraped data
type PokemonProduct struct {
	Name   string `json:name`
	Price  string `json:price`
	ImgUrl string `json:imgurl`
}

func main() {

	c := colly.NewCollector()

	c.OnHTML("a.woocommerce-LoopProduct-link.woocommerce-loop-product__link", func(h *colly.HTMLElement) {

		fmt.Println(h.Text)

	})

	c.Visit("https://scrapeme.live/shop/")

}
