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
