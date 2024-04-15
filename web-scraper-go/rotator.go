package main

import (
	"fmt"
	"math/rand"

	// "time"

	"github.com/gocolly/colly"
)

// initializing a data structure to keep the scraped data
type Proxy struct {
	Ip   string `json:"proxy"`
	Port string `json:"port"`
}

func rotator() (Proxy, error) {

	c := colly.NewCollector()

	//User Agent change. Colly agents remain identifiable by anti-scrapping technologies by default.
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	var proxies []Proxy

	c.OnHTML("table.table.table-striped.table-bordered tr ", func(h *colly.HTMLElement) {

		ip := h.ChildText("td:nth-child(1)")
		port := h.ChildText("td:nth-child(2)")

		// fmt.Printf("IP='%s', 'PORT:'%s'	\n", ip, port) //debugging purposes

		//first line gets empty because it scrapes first element of  thetable which his a title
		if ip != "" && port != "" {

			p := Proxy{
				Ip:   ip,
				Port: port,
			}
			proxies = append(proxies, p)
		}

		//debug double check
		// fmt.Printf("Current Proxy List: %+v\n", proxies)

	})

	c.Visit("https://sslproxies.org/")
	c.Wait()

	if len(proxies) > 0 {
		//random proxy
		r := rand.Intn(len(proxies))
		randProxy := proxies[r]

		fmt.Println(randProxy)

		return randProxy, nil
	}

	return Proxy{}, fmt.Errorf("no proxies found")

}

// {52.35.240.119 1080}
// func ticker() {

// 	ticker := time.NewTicker(2 * time.Hour)
// 	defer ticker.Stop()

// 	for range ticker.C {
// 		rotator()
// 	}
// }
