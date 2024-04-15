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

	randProxy, err := rotator()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Using Proxy: %s:%s\n", randProxy.Ip, randProxy.Port)


	


	//Colly Scraping

	c := colly.NewCollector()

	//User Agent change. Colly agents remain identifiable by anti-scrapping technologies by default.
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	proxyURL := fmt.Sprintf("http://%s:%s", randProxy.Ip, randProxy.Port)
	if err := c.SetProxy(proxyURL); err != nil {
		log.Fatal("Failed proxy in the collector : ", err)
	}

	var items []item

	c.OnHTML("a.woocommerce-LoopProduct-link.woocommerce-loop-product__link", func(h *colly.HTMLElement) {

		item := item{
			Name:   h.ChildText("h2.woocommerce-loop-product__title"),
			Price:  h.ChildText("span.woocommerce-Price-amount.amount"),
			ImgUrl: h.ChildAttr("img", "src"),
		}

		items = append(items, item)

		// fmt.Println(item)

	})

	c.OnHTML("a.next.page-numbers", func(h *colly.HTMLElement) {
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(next_page)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.Visit("https://scrapeme.live/shop/")

	// fmt.Println(items)

	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("pokedex.json", content, 0644)
	// ticker()
}
