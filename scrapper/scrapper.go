package scrapper

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

const SOURCE_URL = "https://monitordolarvenezuela.com/"

var NAMES_MAP = map[string]string{
	"BCV (Oficial)":      "BCV",
	"@EnParaleloVzla3":   "EnParaleloVzla",
	"@DolarToday":        "DolarToday",
	"@MonitorDolarWeb":   "MonitorDolar",
	"@EnParaleloVzlaVip": "EnParaleloVzlaVip",
	"Binance P2P":        "Binance",
}

type DolarPrices struct {
	BCV               float64
	EnParaleloVzla    float64
	DolarToday        float64
	MonitorDolar      float64
	EnParaleloVzlaVip float64
	Binance           float64
}

func GetDolarPrice() {
	c := colly.NewCollector()
	var dolarPrices DolarPrices
	mapPrices := make(map[string]float64)

	c.OnRequest(func(r *colly.Request) {
		println("Visiting", r.URL.String())
	})

	c.OnHTML("div > .web", func(e *colly.HTMLElement) {

		nodes := e.DOM.SiblingsFiltered("p")

		nodes.Each(func(i int, s *goquery.Selection) {

			rawName := s.Parent().ChildrenFiltered(".title-prome").Text()
			rawValue := strings.Split(s.Text(), " = ")[1]

			parsedValue, err := strconv.ParseFloat(strings.Replace(rawValue, ",", ".", 1), 32)

			if err != nil {
				return
			}

			mapPrices[NAMES_MAP[rawName]] = parsedValue
		})

		data, _ := json.Marshal(mapPrices)
		json.Unmarshal(data, &dolarPrices)
	})

	c.OnScraped(func(r *colly.Response) {
		// After scraping, add to database
		fmt.Printf("%+v", dolarPrices)
	})

	c.Visit(SOURCE_URL)
}
