package scrape

import (
	"fmt"
	"log"
	"testing"
)

func TestScrape(t *testing.T) {
	res, err := Scrape("今日はいい天気ですね")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
