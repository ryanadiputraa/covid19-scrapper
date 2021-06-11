package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
	cleanData := make([]int, 0)
	for _, d := range collectedData {
		tr := strings.ReplaceAll(d, ".", "")
		d, err := strconv.Atoi(tr)
		if err != nil {
			log.Println("cant convert data")
		}
		cleanData = append(cleanData, d)
	}

	fmt.Println(cleanData)
}