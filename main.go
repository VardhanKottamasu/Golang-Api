package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	baseUrl = "https://api.coinbase.com/v2/prices/spot?currency="
)

type Data struct {
	Base     string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

func main() {
	fmt.Println("Choose a currency type")
	fmt.Println("1.INR")
	fmt.Println("2.USD")
	fmt.Println("3.EUR")
	fmt.Println("4.JPY")
	var inp string
	fmt.Scan(&inp)
	
	currencyMap := make(map[string]string)
	currencyMap["1"] = "INR"
	currencyMap["2"] = "USD"
	currencyMap["3"] = "EUR"
	currencyMap["4"] = "JPY"
	
	var url string
	url=baseUrl+currencyMap[inp]
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error response from API:", response.Status)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	op := struct {
		Data Data `json:"data"`
	}{}
	// op := &Data{}
	err = json.Unmarshal(data, &op)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("Base currency is ",op.Data.Base)
	fmt.Println("Selected Currency is",op.Data.Currency)
	fmt.Println("Value of a BTC in the selected currency is",strings.Split(op.Data.Amount, ".")[0])
}
