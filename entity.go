package main

import (
	"strings"

	"github.com/antchfx/jsonquery"
)

type Coin struct {
	Name   string
	Symbol string
	Price  string
}

func GetCoin(j string) (Coin, error) {
	var coin Coin
	doc, err := jsonquery.Parse(strings.NewReader(j))
	if err != nil {

	}
	coin.Name = jsonquery.FindOne(doc, "data/*/name").InnerText()
	coin.Symbol = jsonquery.FindOne(doc, "data/*/symbol").InnerText()
	coin.Price = jsonquery.FindOne(doc, "data/*/quote/USD/price").InnerText()

	return coin, nil
}
