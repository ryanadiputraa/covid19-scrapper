package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type WorldData struct {
	TotalCoutry		int		`json:"totalCountry"`
	ConfirmedCase	int		`json:"confirmedCased"`
	Deaths			int		`json:"deaths"`
}

func main() {
	collectedData := make([]string, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("covid19.go.id"),
	)

	collector.OnHTML(".col-md-3 div strong", func(h *colly.HTMLElement) {
	 	data := h.Text
		collectedData = append(collectedData, data)
	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	collector.Visit("https://covid19.go.id/")
	
	// trim and convert data
	for _, d := range collectedData {
		fmt.Println(d)
	}
}