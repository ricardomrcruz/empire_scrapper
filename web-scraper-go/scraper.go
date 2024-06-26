package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

// initializing a data structure to keep the scraped data
type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {

	const (
		seleniumPathh = "path"
	)

	var items []item

	c.OnHTML("div.sg-col-4-of-24.sg-col-4-of-12.s-result-item.s-asin.sg-col-4-of-16.sg-col.s-widget-spacing-small.sg-col-4-of-20", func(h *colly.HTMLElement) {

		item := item{
			Name:   h.ChildText("span.a-size-base-plus.a-color-base.a-text-normal"),
			Price:  h.ChildText("span.a-price-whole"),
			ImgUrl: h.ChildAttr("img.s-image", "src"),
		}

		items = append(items, item)

		fmt.Println(item)

	})

	err := c.Visit("https://www.amazon.fr/s?k=playstation+5+ps5+console")
	if err != nil {
		log.Fatal("error:", err)
	}

	// fmt.Println(items)

	content, err := json.Marshal(items)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("text_ps5.json", content, 0644)
	if err != nil {
		log.Fatal("Failed to write file:", err)
	}

}
