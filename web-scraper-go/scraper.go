package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

// initializing a data structure to keep the scraped data
type PokemonProduct struct {
	url, image, name, price string
}

func main() {
	// initializing the slice of structs to store the data to scrape
	var pokemonProducts []PokemonProduct

	// creating a new Colly instance
	c := colly.NewCollector()

	// visiting the target page
	c.Visit("https://scrapeme.live/shop/")

	// Wait until Colly finishes executing
	c.Wait()

	// Now, proceed with opening the CSV file and writing data
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush() // Ensure flush is called to write any buffered data

	// writing the CSV headers
	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}
	if err := writer.Write(headers); err != nil {
		log.Fatalln("Error writing headers to CSV:", err)
	}

	// writing each Pokemon product as a CSV row
	for _, pokemonProduct := range pokemonProducts {
		record := []string{
			pokemonProduct.url,
			pokemonProduct.image,
			pokemonProduct.name,
			pokemonProduct.price,
		}

		if err := writer.Write(record); err != nil {
			log.Fatalln("Error writing record to CSV:", err)
		}
	}
}
