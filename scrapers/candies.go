package scrapers

import (
	"io"
	"net/http"
	"strings"
	"williamzelesny/release-namer/releasenamer"

	"github.com/PuerkitoBio/goquery"
)

func GetCandies() ([]*releasenamer.Candy, error) {
	var candies []*releasenamer.Candy

	// Send a GET request to the Wikipedia page
	resp, err := http.Get("https://en.wikipedia.org/wiki/List_of_candies")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Parse the HTML using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// Find all the candy names in the table and add them to the candies array
	doc.Find(".wikitable.sortable td:first-child").Each(func(i int, s *goquery.Selection) {
		// Remove any links and parentheses from the text
		candyName := strings.TrimSpace(strings.TrimSuffix(s.Find("a").Text(), " ("))

		// Only add non-blank candy names to the array
		if candyName != "" {
			candies = append(candies, &releasenamer.Candy{Name: candyName})
		}
	})

	return candies, nil
}
