package crawler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const HUXIU = "http://www.huxiu.com"

func ExampleScrape() map[string]string {
	urlToTitle := make(map[string]string)
	// Request the HTML page.
	res, err := http.Get(HUXIU)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.mob-ctt").Find("a.transition").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		shortPath, err := s.Attr("href")
		if err == false {
			log.Fatal(err)
		}
		href := fmt.Sprintf("%s%s", HUXIU, shortPath)
		title := s.Text()
		urlToTitle[href] = title
	})
	return urlToTitle
}
