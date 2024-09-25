package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func scrape(url string, wg *sync.WaitGroup, ch chan []string) {
	defer wg.Done() // Mark this Goroutine as done

	// Fetch the HTML content from the URL
	res, err := http.Get(url)
	if err != nil {
		ch <- []string{url, "Error fetching URL", ""}
		return
	}
	defer res.Body.Close()

	// Check for a successful status code
	if res.StatusCode != 200 {
		ch <- []string{url, fmt.Sprintf("Failed with status: %d", res.StatusCode), ""}
		return
	}

	// Parse the HTML document using goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		ch <- []string{url, "Error parsing HTML", ""}
		return
	}

	// Extract the title and meta description
	title := doc.Find("title").Text()
	metaDesc, _ := doc.Find("meta[name='description']").Attr("content")

	// Send the scraped data back to the channel
	ch <- []string{url, title, metaDesc}
}

func main() {
	// List of URLs to scrape
	urls := []string{
		"https://www.bbc.com",
		"https://www.nytimes.com",
		"https://developer.mozilla.org/en-US/",
		"https://stackoverflow.com",
		"https://techcrunch.com",
		"https://www.producthunt.com",
	}

	// Create or open the CSV file
	file, err := os.Create("scraped_data.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	// Initialize the CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the CSV header
	writer.Write([]string{"URL", "Title", "Meta Description"})

	// Initialize a WaitGroup and a channel
	var wg sync.WaitGroup
	ch := make(chan []string)

	// Scrape each URL concurrently
	for _, url := range urls {
		wg.Add(1)
		go scrape(url, &wg, ch)
	}

	// Goroutine to listen to the channel and write data to the CSV file
	go func() {
		for data := range ch {
			writer.Write(data)
		}
	}()

	// Wait for all Goroutines to finish
	wg.Wait()
	close(ch) // Close the channel when all Goroutines are done
}
