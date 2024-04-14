package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gocolly/colly"
)

// initializing a data structure to keep the scraped data
type proxy struct {
	Ip   string `json:"proxy"`
	Port string `json:"port"`
}

func rotator() {

	c := colly.NewCollector()

	//User Agent change. Colly agents remain identifiable by anti-scrapping technologies by default.
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	var proxies []proxy

	c.OnHTML("table.table.table-striped.table-bordered", func(h *colly.HTMLElement) {

		proxy := proxy{
			Ip:   h.ChildText("td:nth-child(1)"),
			Port: h.ChildText("td:nth-child(2)"),
		}

		proxies = append(proxies, proxy)

		//  fmt.Println(proxy)

	})

	c.Visit("https://sslproxies.org/")
	c.Wait()

	r := rand.Intn(len(proxies))
	proxy := proxies[r]

	fmt.Println(proxy)

}

func ticker() {

	ticker := time.NewTicker(2 * time.Hour)

	for range ticker.C {
		rotator()
	}
}
