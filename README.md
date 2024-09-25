# Concurrent-web-scraper-using-Golang

## Project Overview
This project implements a concurrent web scraper built in Go. It scrapes multiple websites simultaneously to extract specific information such as page titles and meta descriptions using Goroutines and channels, making it efficient for handling multiple requests concurrently.

## Objective
To develop a high-performance web scraper that can retrieve information from various websites concurrently, improving both speed and scalability.

## Features
- Concurrent scraping using Goroutines and channels for high efficiency.
- Scrapes HTML content and extracts specific data (titles, meta descriptions) using the `goquery` package.
- Stores scraped data in a CSV file for easy access and analysis.
- Synchronization using `sync.WaitGroup` to manage Goroutines.

## Tech Stack
- **Language**: Go (Golang)
- **Libraries**: `net/http`, `goquery`, `encoding/csv`
- **Concurrency**: Goroutines and channels
- **File Output**: CSV for storing scraped data

## How to Run

1. **Clone the repository**:
    ```bash
    git clone https://github.com/shivanshi1/Concurrent-web-scraper-using-Golang.git
    cd Concurrent-web-scraper-using-Golang
    ```

2. **Install dependencies**:
    ```bash
    go mod tidy
    ```

3. **Run the scraper**:
    ```bash
    go run main.go
    ```

4. **View the output**:
   The output will be saved to `scraped_data.csv` in the project directory.

## Example of Usage
The scraper can be customized to target different websites. Modify the `urls` array in the `main.go` file to scrape the URLs of your choice.
