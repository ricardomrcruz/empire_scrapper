package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// initializing a data structure to keep the scraped data
type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {

	c := colly.NewCollector()

	
	//User Agent change. Colly agents remain identifiable by anti-scrapping technologies by default.
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"


	c.OnHTML("a.woocommerce-LoopProduct-link.woocommerce-loop-product__link", func(h *colly.HTMLElement) {

		item := item{
			Name:   h.ChildText("h2.woocommerce-loop-product__title"),
			Price:  h.ChildText("span.woocommerce-Price-amount.amount"),
			ImgUrl: h.ChildAttr("img", "src"),
		}

		fmt.Println(item)

	})

	c.Visit("https://scrapeme.live/shop/")

}
