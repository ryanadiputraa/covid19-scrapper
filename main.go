package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type World struct {
	TotalCoutry		int		`json:"totalCountry"`
	ConfirmedCase	int		`json:"confirmedCased"`
	Deaths			int		`json:"deaths"`
}

type Indonesia struct {
	Positive	int		`json:"positive"`
	Recover		int		`json:"recover"`
	Deaths		int		`json:"deaths"`	
}

type ResponseData struct {
	WorldCase 		World		`json:"worldCase"`
	IndonesiaCase 	Indonesia	`json:"indonesiaCase"`
}

type Response struct {
	Code	int				`json:"code"`
	Data	interface{}		`json:"data"`
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

	worldCase := World {
		TotalCoutry: cleanData[0],
		ConfirmedCase: cleanData[1],
		Deaths: cleanData[2],
	}

	indonesiaCase := Indonesia {
		Positive: cleanData[3],
		Recover: cleanData[4],
		Deaths: cleanData[5],
	}

	http.HandleFunc("/api/covidcases", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")

		response := Response {
			Code: http.StatusOK,
			Data: ResponseData {worldCase, indonesiaCase},
		}

		enc, err := json.Marshal(response)
		if err != nil {
			w.Header().Set("Status", strconv.Itoa(http.StatusInternalServerError))
			w.Write([]byte("cant get covid cases"))
		}

		w.Header().Set("Status", strconv.Itoa(http.StatusOK))
		w.Write(enc)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}